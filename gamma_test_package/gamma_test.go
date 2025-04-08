package gamma

import (
	"math/rand"
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

func TestPass2(t *testing.T) {
	t.Parallel()
	t.Logf("%s: This test always passes\n", packageName)
}

func TestPass3(t *testing.T) {
	t.Parallel()
	t.Logf("%s: This test always passes\n", packageName)
}

func TestRandomFlaky(t *testing.T) {
	t.Parallel()
	random := rand.Intn(4)

	if random == 0 {
		t.Fatal("This test is designed to flake 1/4 of the time")
	}

	t.Log("This test is designed to pass 3/4 of the time")
}
