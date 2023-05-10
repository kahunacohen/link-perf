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
	linkProfiler := NewLinkProfiler([]string{"https://www.google.com", "https://www.yahoo.com", "https://www.microsoft.com"}, 3)
	linkProfiler.Run()
	googleProfiles := linkProfiler.Get("https://www.google.com")
	if len(googleProfiles) != 3 {
		t.Fatalf("wanted length 3 for google profiles, got %d", len(googleProfiles))
	}
	if googleProfiles[0].GotFirstResponseByte < 1 {
		t.Fatalf("wanted GotFirstResponseByte for google profile, got %d", googleProfiles[0].GotFirstResponseByte)
	}
	if googleProfiles[0].DNSStart < 1 {
		t.Fatalf("wanted DNSStart for yahoo profile, got %d", googleProfiles[0].DNSStart)
	}
	yahooProfiles := linkProfiler.Get("https://www.yahoo.com")
	if len(yahooProfiles) != 3 {
		t.Fatalf("wanted length 3 for yahoo profiles, got %d", len(yahooProfiles))
	}
	if yahooProfiles[0].GotFirstResponseByte < 1 {
		t.Fatalf("wanted GotFirstResponseByte for yahoo profile, got %d", yahooProfiles[0].GotFirstResponseByte)
	}
	if yahooProfiles[0].DNSStart < 1 {
		t.Fatalf("wanted DNSStart for yahoo profile, got %d", yahooProfiles[0].DNSStart)
	}
}
