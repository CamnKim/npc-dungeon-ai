import { createFileRoute } from '@tanstack/react-router';
import OnboardingPage from '../../pages/onboard';

export const Route = createFileRoute('/_authenticated/onboard')({
  component: OnboardingPage,
});
