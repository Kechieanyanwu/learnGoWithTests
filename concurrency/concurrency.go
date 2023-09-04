package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool { //notice we call the function with a function type, so we have more flexibility
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// we are no longer iterating over the urls and writing the results to the map directly. We are now passing through channels
	for _, url := range urls { //this section passes the results from the separate go routines into a channel
		go func(u string) { // an anonymous function
			resultChannel <- result{u, wc(u)} // by sending the results into a channel, we avoid data races and ensure one write is made per time
		}(url)
	}

	for i := 0; i < len(urls); i++ { //this section populates the results map with the individual results as gotten from the channel
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}

// with the above, We have parallelized the part of the code that we wanted to make faster,
// while making sure that the part that cannot happen in parallel still happens linearly.
// We have also communicated across the multiple processes involved by using channels
