package concurrency

import ()

type result struct {
	string
	bool
}

// WebsiteChecker defines a function type
type WebsiteChecker func(string) bool

// CheckWebsite checks if the url is correct
func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	retCh := make(chan result)

	for _, url := range urls {
		go func(u string) {
			retCh <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-retCh
		results[result.string] = result.bool
	}

	return results
}
