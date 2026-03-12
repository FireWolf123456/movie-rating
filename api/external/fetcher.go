package external

// RatingsFetcher fetches a rating (1–100) for a movie from an external source.
type RatingsFetcher interface {
	// Name identifies the rating source (e.g. "imdb").
	// Use it as the Source field when persisting a Ratings entry.
	Name() string
	// FetchRating returns a 1–100 integer rating for the given movie title.
	FetchRating(movieTitle string) (int, error)
}

// NewProviders returns the three rating providers used by the startup sync job.
// In tests, pass MockFetcher or ErrFetcher instances directly instead.
func NewProviders() []RatingsFetcher {
	return []RatingsFetcher{
		&MockFetcher{ProviderName: "imdb", DefaultRating: 72},
		&MockFetcher{ProviderName: "rottentomatoes", DefaultRating: 68},
		&MockFetcher{ProviderName: "metacritic", DefaultRating: 65},
	}
}
