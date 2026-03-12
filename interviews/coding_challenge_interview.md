# Movie Rating API — Coding Challenge

## Position: Sr. Software Engineer - Backend Go
## Time Limit: 30 minutes

---

## Summary

This is a working Go API for rating movies. Your task is to add two features to it. Spend a couple of minutes reading the codebase before you start — the existing patterns are intentional and your implementation should follow them.

---

## Codebase Orientation

```
api/
├── main.go            # Entry point — wires DB, router, HTTP server
├── models/models.go   # Data structs: Movies, MovieRatings, Ratings
├── db/db.go           # DB interface — extend this for new DB operations
├── db/setup.go        # DB init, migrations, seeding
├── http/http.go       # HTTP handlers — add new routes here
├── app/app.go         # Business logic — add new functions here
└── external/          # RatingsFetcher interface + mock (provided, see below)
```

**Before you write any code, note the following:**

- The project uses **GORM v1** (`github.com/jinzhu/gorm`), not v2. See the [v1 docs](https://v1.gorm.io/docs/).
- `GetMovies()` and `GetMovieRatings()` in `db/db.go` each have a `time.Sleep(1 * time.Second)`. **Do not remove it** — it simulates a slow downstream service and is intentional.
- Extract gorilla/mux path variables like this: `vars := mux.Vars(r); id := vars["movie_id"]`
- `{movie_id}` in the path maps to `movies.id` (the `Movies` table PK).
- All API routes are prefixed with `/api` (see `ConfigureRouter` in `http/http.go`).

---

## Data Model

| Table | Key Fields | Notes |
|---|---|---|
| `movies` | `id`, `title`, `genre`, `plot` | Movie metadata |
| `movie_ratings` | `id`, `title` | Linked to `movies` by matching `title` string (no DB-level FK) |
| `ratings` | `movie_ratings_id`, `source`, `value` | Belongs to `movie_ratings`; `value` is an integer |

---

## Rating Scale

- All ratings (user-submitted and externally fetched) use the **1–100** integer scale.

---

## Task 1 — `POST /api/movies/{movie_id}/ratings`

Implement an endpoint that lets a user submit a rating for a movie.

**Requirements:**

1. Accept a JSON body: `{"rating": <int>}`
2. Return `400 Bad Request` if the rating is not between 1 and 100.
3. Return `404 Not Found` if no movie exists with the given `movie_id`.
4. Persist the new rating using the existing `Ratings` struct (source: `"user"`).
5. Extend the `DB` interface in `db/db.go` with any new methods you need and implement them on `dbClient`.
6. Register the route in `http/http.go`; add business logic in `app/app.go`.
7. Write unit tests covering: happy path, invalid rating, movie not found.

---

## Task 2 — Startup ratings sync

At application startup, fetch ratings for every movie in the DB from all three rating providers and persist each result.

**Provided for you:** `api/external/fetcher.go` defines the `RatingsFetcher` interface and `NewProviders()` which returns the three providers. `api/external/mock_fetcher.go` provides `MockFetcher` and `ErrFetcher` for use in tests. **Do not make real HTTP calls.**

**Requirements:**

1. Write a function that accepts `[]RatingsFetcher` and a DB client; call it from `main.go` as a goroutine after the router is configured.
2. Fetch ratings concurrently — use goroutines and channels to fan out across all providers and all movies. All fetches must run in parallel; do not call providers or movies sequentially.
3. Persist each result as a new `Ratings` entry; use the provider's `Name()` as the `Source` field.
4. Implement retry logic: retry up to **3 times** with a short delay on transient errors; log and skip if all retries are exhausted.
5. Access to shared data must be **thread-safe**.
6. Write unit tests covering: all providers succeed, one provider fails transiently and recovers, one provider exhausts all retries.

---

## Evaluation Criteria

| Area | What we look for |
|---|---|
| **Code quality** | Clean, readable, idiomatic Go; follows existing patterns |
| **Error handling** | All edge cases handled; correct HTTP status codes |
| **Testing** | Tests are meaningful, not just coverage padding |
| **Concurrency** | Correct use of goroutines; no data races |
