package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
// TODO: VALIDATE CLAIMS
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	aud := []string{fmt.Sprintf("https://%s/api/v2/", os.Getenv("AUTH0_DOMAIN")),
		fmt.Sprintf("https://%s/userinfo", os.Getenv("AUTH0_DOMAIN"))}

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		aud,
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return middleware.CheckJWT(next)
}

func TokenExtractor(r *http.Request) (*string, error) {
	auth := strings.Split(r.Header.Get("Authorization"), " ")
	if len(auth) != 2 || auth[0] != "Bearer" {
		return nil, errors.New("invalid token")
	}
	return &auth[1], nil
}

func (s *Server) injectUserHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := TokenExtractor(r)
		if err != nil {
			log.Fatalf("Failed to extract token")
		}

		// Split the JWT into its parts
		parts := strings.Split(*token, ".")
		if len(parts) != 3 {
			log.Fatalf("Invalid JWT format")
		}

		// Decode the payload (second part of the token)
		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			log.Fatalf("Failed to decode JWT payload: %v", err)
		}

		// Parse the JSON payload
		var claims map[string]interface{}
		err = json.Unmarshal(payload, &claims)
		if err != nil {
			log.Fatalf("Failed to parse JWT payload: %v", err)
		}

		sub := claims["sub"].(string)

		user, err := s.db.User().GetUserByAuth0(sub)
		if err != nil {
			log.Fatalf("Failed to fetch user %v", err)
		}

		var userStr string
		if user == nil {
			userStr = ""
		} else {
			userStr = user.ID
		}

		r.Header.Add("User", userStr)

		next.ServeHTTP(w, r)
	})
}

func (s *Server) CombinedAuthMiddleware(next http.Handler) http.Handler {
	return s.authMiddleware(s.injectUserHeader(next))
}
