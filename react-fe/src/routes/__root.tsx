import { createRootRoute, Link, Outlet } from '@tanstack/react-router'

const RootLayout = () => (
  <>
    <div className="p-2 flex gap-2">
      <Link to="/" className="[&.active]:font-bold">
        Home
      </Link>{' '}
      <Link to="/dashboard" className="[&.active]:font-bold">
        Dashboard
      </Link>
      <Link to="/search" className="[&.active]:font-bold">
        Search
      </Link>
    </div>
    <hr />
    <Outlet />
  </>
)

export const Route = createRootRoute({ component: RootLayout })