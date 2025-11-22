import { createFileRoute } from '@tanstack/react-router'
import { Dashboard as DashboardComponent } from '../components/Dashboard'

export const Route = createFileRoute('/dashboard')({
  component: DashboardComponent,
})
