package linkperf

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"

	"golang.org/x/exp/maps"
)

// A link performance package that collects data on HTTP performance on a list of links.

type Profile struct {
	DNSStart             int64 // Time to DNS start in
	GotFirstResponseByte int64 // Time from request to first byte received in microseconds

}
type LinkProfiler struct {
	mu    *sync.Mutex
	perf  map[string][]Profile
	tries int
}

func NewLinkProfiler(links []string, tries int) LinkProfiler {
	lp := LinkProfiler{mu: &sync.Mutex{}, tries: tries, perf: make(map[string][]Profile)}
	// Transform simple slice of links to internal map of link to
	// slice of Profile structs. The length of the slice should be initialized
	// to the amountof tries.
	for _, l := range links {
		perfs := make([]Profile, tries)
		lp.perf[l] = perfs
	}
	return lp
}

// "Private" method to profile one link. Internally sets the map of links
// to slices of profile structs.
func (lp *LinkProfiler) profileLink(l string, try int, ch chan string) {
	req, err := http.NewRequest("GET", l, nil)

	if err != nil {
		lp.perf[l][try] = Profile{GotFirstResponseByte: -1}
		ch <- err.Error()
		return
	}

	var start time.Time
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			since := time.Since(start)
			fmt.Printf("Time to GotFirstResponseByte for %s in %v\n", l, since)
			lp.mu.Lock()
			defer lp.mu.Unlock()
			lp.perf[l][try].GotFirstResponseByte = int64(since)
		},
		DNSStart: func(di httptrace.DNSStartInfo) {
			since := time.Since(start)
			fmt.Printf("Time to DNSStart for %s in %v\n", l, since)
			lp.mu.Lock()
			defer lp.mu.Unlock()
			lp.perf[l][try].DNSStart = int64(since)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	start = time.Now()

	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		lp.perf[l][try] = Profile{GotFirstResponseByte: -1}
		ch <- err.Error()
		return
	}
	ch <- l
}

// Get a slice of Profile structs for the given link.
func (lp *LinkProfiler) Get(link string) []Profile {
	return lp.perf[link]
}

// Runs a performance profiler on
func (lp *LinkProfiler) Run() {
	for _, l := range maps.Keys(lp.perf) {
		ch := make(chan string)
		for try := 0; try < lp.tries; try++ {
			go lp.profileLink(l, try, ch)
		}
		for range maps.Keys(lp.perf) {
			<-ch
		}
	}
}
