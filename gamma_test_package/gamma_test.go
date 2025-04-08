package gamma

import (
	"testing"
)

const packageName = "gamma"

func TestPass(t *testing.T) {
	t.Parallel()
	t.Logf("%s: This test always passes\n", packageName)
}

func TestFail(t *testing.T) {
	t.Parallel()
	t.Fatalf("%s: This test always fails\n", packageName)
}
