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

//"fmt"
//"net/http"
//"net/http/httptrace"

//"time"
// "golang.org/x/exp/maps"

type Perf struct {
	GotFirstResponseByte int // Time from request to first byte received in ms

}
type LinkPerf struct {
	mu    *sync.Mutex
	perf  map[string][]Perf
	tries int
}

// Initialize a new LinkProfiler.
func NewLinkPerf(links []string, tries int) LinkPerf {
	lp := LinkPerf{mu: &sync.Mutex{}, tries: tries, perf: make(map[string][]Perf)}
	// Transform simple slice of links to internal map of link to
	// slice of Perf structs. The length of the slice should be initialized
	// to the amountof tries.
	for _, l := range links {
		perfs := make([]Perf, tries)
		lp.perf[l] = perfs
	}
	return lp
}

// Runs a performance profile on time to first byte received for
// links.
func (lp *LinkPerf) Run() map[string][]Perf {
	ch := make(chan string)
	for i, l := range maps.Keys(lp.perf) {
		for j := 0; j <= lp.tries - 1; j++ {
			go func(l string, i int) {
				req, err := http.NewRequest("GET", l, nil)
				if err != nil {
					lp.perf[l][i] = Perf{GotFirstResponseByte: -1}
					ch <- err.Error()
					return
				}
				var start time.Time
				trace := &httptrace.ClientTrace{
					GotFirstResponseByte: func() {
						since := time.Since(start)
						fmt.Printf("Got %s in %v\n", l, since)
						lp.mu.Lock()
						defer lp.mu.Unlock()
						lp.perf[l][i] = Perf{GotFirstResponseByte: int(time.Since(start) / time.Millisecond)}
					},
				}
				req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
				start = time.Now()
				if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
					lp.perf[l][i] = Perf{GotFirstResponseByte: -1}
					ch <- err.Error()
					return
				}
				ch <- l
			}(l, i)
		}
	}
	for range maps.Keys(lp.perf) {
		<-ch
	}
	return lp.perf
}
