package linkperf

import (
	"testing"
)

func TestNewLinkPerf(t *testing.T) {
	lp := NewLinkProfiler([]string{"https://www.google.com"}, 2)
	googlePerf := lp.perf["https://www.google.com"]
	if len(googlePerf) != 2 {
		t.Fatalf("wanted googlePerf length to be 2, got %d", len(googlePerf))
	}

	if googlePerf[0].GotFirstResponseByte != 0 {
		t.Fatalf("wanted first perf struct's GotFirstResponseByte to be default value of 0, got: %d", googlePerf[0].GotFirstResponseByte)
	}
}
func TestRun(t *testing.T) {
	// Run link profiler on two links. Try each link twice.
	linkProfiler := NewLinkProfiler([]string{"https://www.google.com"}, 1)
	linkProfiler.Run()
	// googleProfiles := linkProfiler.Get("https://www.google.com")

	// if ! (googleProfiles[0].GotFirstResponseByte > 0) {
	// 	t.Fatalf("wanted first google's element's GotFirstResponseByte to be greater than 0")
	// }
	// if ! (googleProfiles[1].GotFirstResponseByte > 0) {
	// 	t.Fatalf("wanted second google's element's GotFirstResponseByte to be greater than 0")
	// }
	// yahooProfiles := linkProfiler.Get("https://www.yahoo.com")

	// if ! (yahooProfiles[0].GotFirstResponseByte > 0) {
	// 	t.Fatalf("wanted first yahoo's element's GotFirstResponseByte to be greater than 0")
	// }
	// if ! (yahooProfiles[1].GotFirstResponseByte > 0) {
	// 	t.Fatalf("wanted second yahoo's element's GotFirstResponseByte to be greater than 0")
	// }
}
