# LSCS Central Auth API

## TODOs

- [x] feat: testing db sqlite (for dev)
- [ ] feat: port database to MongoDB or PostgreSQL (for production)
- [ ] feat: add token generation (JWT and refresh tokens)
    - [ ] feat: add `id`, `email`, etc. from google `profile` in JWT as Claims
    - [ ] docs: frontend to verify JWT
- [ ] build: dockerize for dev and prod builds
- [ ] chore: update to es6

## Route Endpoints

Everything will be redirected to `/` after successful login

`/login/google`
- endpoint for logging in to google

`/logout`
- for logging out
