package linkperf

import (
	"fmt"
	"testing"
)

func TestNewLinkPerf(t *testing.T) {
	lp := NewLinkPerf([]string{"https://www.google.com"}, 2)
	googlePerf := lp.perf["https://www.google.com"]
	if len(googlePerf) != 2 {
		t.Fatalf("wanted googlePerf length to be 2, got %d", len(googlePerf))
	}
	
	if googlePerf[0].GotFirstResponseByte != 0 {
		t.Fatalf("wanted first perf struct's GotFirstResponseByte to be default value of 0, got: %d", googlePerf[0].GotFirstResponseByte)
	}
}
func TestRun(t *testing.T) {
	lp := NewLinkPerf([]string{"https://www.google.com"}, 2)
	fmt.Println(lp.Run())
}
