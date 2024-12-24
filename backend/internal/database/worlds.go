package database

import "database/sql"

type WorldDAO interface {
	GetWorldById(id string) (*World, error)
	CreateWorld(world *WorldInsert) (*World, error)
	UpdateWorld(world *WorldUpdate) (*World, error)
	DeleteWorld(id string) error
}

type worldDAO struct {
	db *sql.DB
}

type WorldInsert struct {
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	Background string `json:"background"`
}

type World struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	Background string `json:"background"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsDeleted  bool   `json:"is_deleted"`
}

type WorldUpdate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Background string `json:"background"`
}

func (s *service) World() WorldDAO {
	return &worldDAO{
		db: s.db,
	}
}

func (w *worldDAO) GetWorldById(id string) (*World, error) {
	const query = `SELECT id, name, created_by, created_at, updated_at, background, is_deleted FROM worlds WHERE id = $1`
	row := w.db.QueryRow(query, id)
	world := &World{}
	err := row.Scan(&world.ID, &world.Name, &world.CreatedBy, &world.CreatedAt, &world.UpdatedAt, &world.Background, &world.IsDeleted)

	if err != nil {
		return nil, err
	}
	return world, nil
}

func (w *worldDAO) CreateWorld(world *WorldInsert) (*World, error) {
	const query = `INSERT INTO worlds (name, created_by, background) VALUES ($1, $2, $3) RETURNING id, name, created_by, created_at, updated_at, background, is_deleted`

	res := &World{}
	err := w.db.QueryRow(query, world.Name, world.CreatedBy, world.Background).Scan(&res.ID, &res.Name, &res.CreatedBy, &res.CreatedAt, &res.UpdatedAt, &res.Background, &res.IsDeleted)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (w *worldDAO) UpdateWorld(world *WorldUpdate) (*World, error) {
	const query = `UPDATE worlds SET name = $1, background = $2, updated_at = NOW() WHERE id = $3 RETURNING id, name, created_by, created_at, updated_at, background, is_deleted`

	res := &World{}
	err := w.db.QueryRow(query, world.Name, world.Background, world.ID).Scan(&res.ID, &res.Name, &res.CreatedBy, &res.CreatedAt, &res.UpdatedAt, &res.Background, &res.IsDeleted)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (w *worldDAO) DeleteWorld(id string) error {
	const query = `UPDATE worlds SET is_deleted = TRUE WHERE id = $1`
	_, err := w.db.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
