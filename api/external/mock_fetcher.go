package external

import "errors"

// MockFetcher is a controllable test double for RatingsFetcher.
//
// Usage — always succeeds with a default rating:
//
//	f := &MockFetcher{ProviderName: "imdb", DefaultRating: 75}
//
// Usage — per-title ratings:
//
//	f := &MockFetcher{ProviderName: "imdb", Ratings: map[string]int{"Inception": 82}}
//
// Usage — fails for the first N calls then succeeds (retry testing):
//
//	f := &MockFetcher{
//	    ProviderName:  "imdb",
//	    DefaultRating: 75,
//	    Err:           errors.New("network timeout"),
//	    FailTimes:     2,
//	}
//
// Note: MockFetcher is not safe for concurrent use. Use it in single-goroutine tests only.
type MockFetcher struct {
	// ProviderName is returned by Name().
	ProviderName string

	// Ratings maps movie title to a specific rating; takes priority over DefaultRating.
	Ratings map[string]int

	// DefaultRating is returned when the title is not in Ratings (if non-zero).
	DefaultRating int

	// Err is the error returned on failing calls.
	Err error

	// FailTimes is the number of leading calls that return Err before succeeding.
	// If FailTimes is 0 and Err is non-nil, every call returns Err.
	FailTimes int

	// CallCount is incremented on every call to FetchRating.
	CallCount int
}

func (m *MockFetcher) Name() string { return m.ProviderName }

// FetchRating returns the configured rating for movieTitle, subject to the failure behaviour.
func (m *MockFetcher) FetchRating(movieTitle string) (int, error) {
	m.CallCount++
	if m.Err != nil && (m.FailTimes == 0 || m.CallCount <= m.FailTimes) {
		return 0, m.Err
	}
	if rating, ok := m.Ratings[movieTitle]; ok {
		return rating, nil
	}
	if m.DefaultRating != 0 {
		return m.DefaultRating, nil
	}
	return 0, errors.New("mock fetcher: no rating configured for movie: " + movieTitle)
}

// ErrFetcher always returns an error regardless of input.
// Use this to test the exhausted-retries code path.
//
// Usage:
//
//	f := &ErrFetcher{ProviderName: "imdb", Err: errors.New("service unavailable")}
type ErrFetcher struct {
	// ProviderName is returned by Name().
	ProviderName string

	// Err is the error returned on every call. If nil, a default error is used.
	Err error

	// CallCount is incremented on every call to FetchRating.
	CallCount int
}

func (e *ErrFetcher) Name() string { return e.ProviderName }

// FetchRating always returns an error.
func (e *ErrFetcher) FetchRating(_ string) (int, error) {
	e.CallCount++
	if e.Err != nil {
		return 0, e.Err
	}
	return 0, errors.New("ErrFetcher: always fails")
}
