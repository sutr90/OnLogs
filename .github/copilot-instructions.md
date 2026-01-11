# OnLogs — Copilot Instructions for AI Coding Agents

Purpose: Give AI agents the minimal, repository-specific knowledge to be productive quickly.

1) Big picture
- Backend: Go HTTP server in `application/backend` (entry: `main.go`) — exposes many REST endpoints in `app/routes` and stores logs and metadata in LevelDB under `/leveldb`.
- Frontend: Svelte + Vite app in `application/frontend` (entry: `src/App.svelte`, `src/main.js`). Bundled into Docker image for deployment.
- Deployment: Small Docker image built from `Dockerfile` and helper scripts `build.sh` / `release.sh`. Typical usage is via Docker or Docker Compose (see top-level `README.md`).

2) Key flows & integration points
- Log ingestion and streaming: `app/streamer` collects logs; `AGENT=true` toggles agent mode (no UI, sends logs to remote `HOST`).
- Auth / user: JWT secret is read/written at `leveldb/JWT_secret`. Initial user created by `app/util.CreateInitUser()` on startup.
- Routes map: `main.go` registers all API routes using `ONLOGS_PATH_PREFIX` — when changing endpoints, update `main.go` and `app/routes`.

3) Developer workflows (how to run/build/test locally)
- Backend dev: from repo root run `cd application/backend && go run .` (ensure `go.mod` dependencies installed). Default port is `2874`.
- Frontend dev: `cd application/frontend && npm install && npm run dev` (Vite). The frontend expects `ONLOGS_PATH_PREFIX` substitution — `app/util.ReplacePrefixVariableForFrontend()` prepares static files during backend start.
- Full stack via Docker: use the Docker Compose snippet in `README.md` or build image with `build.sh`.
- Tests: backend tests live alongside packages (files ending `_test.go`) — run `cd application/backend && go test ./...`.

4) Project-specific conventions and patterns
- Environment-driven behavior: many behaviors (port, JWT, agent mode, path prefix, docker socket path) are controlled by env vars — prefer adding env var handling to startup logic rather than hardcoding.
- LevelDB is the single source of truth for storage. Files under `app/*` read/write LevelDB directly — keep migrations/backwards compatibility in mind when changing schemas.
- Route handlers are thin; business logic belongs in `app/*` packages (e.g., `app/db`, `app/streamer`, `app/util`). Follow existing package separation.

5) Files to inspect for changes (fast lookup)
- Backend entry: `application/backend/main.go`
- HTTP handlers: `application/backend/app/routes/routes.go`
- Streaming logic: `application/backend/app/streamer/streamer.go`
- DB helpers: `application/backend/app/db/containerdb.go` (and other `app/db` files)
- Frontend entry: `application/frontend/src/main.js` and `application/frontend/src/App.svelte`
- Docker: top-level `Dockerfile`, `application/Dockerfile`, `build.sh`, `release.sh`

6) Common edit examples
- Add API endpoint: add handler in `app/routes`, register in `main.go` under the correct `ONLOGS_PATH_PREFIX`, add tests in the corresponding package.
- Change JWT behavior: update `application/backend/main.go` initialization and `leveldb/JWT_secret` handling; ensure tests cover token generation/validation.

7) Safety & test expectations
- Keep the LevelDB API usage consistent; unit tests assume DB operations are side-effect free when using temporary paths.
- When touching streaming, watch for goroutines started in `main.go` (`streamer.StreamLogs()` and `db.DeleteUnusedTokens()`); ensure shutdown behavior is preserved.

8) When you are unsure
- Search for usages under `application/backend/app` for examples. Prefer existing helper functions (`app/util`, `app/db`) rather than introducing new global state.

If any area is unclear or you want more detail (CI, release process, or frontend routing specifics), say which part and I'll expand or update this file.
