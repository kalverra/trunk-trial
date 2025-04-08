package alphatestpackage

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"
)

const packageName = "alpha"

func TestPass(t *testing.T) {
	t.Parallel()
	t.Logf("%s: This test always passes\n", packageName)
}

func TestFail(t *testing.T) {
	t.Parallel()
	t.Fatalf("%s: This test always fails\n", packageName)
}

func TestFailLargeOutput(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping large output test for now")
	for range 1000 {
		t.Logf("%s: This is a log line\n", packageName)
	}
	t.Fatalf("%s: This test always fails\n", packageName)
}

func TestSubTestsAllPass(t *testing.T) {
	t.Parallel()

	t.Run("Pass1", func(t *testing.T) {
		t.Logf("%s: This subtest always passes\n", packageName)
	})

	t.Run("Pass2", func(t *testing.T) {
		t.Logf("%s: This subtest always passes\n", packageName)
	})
}

func TestSubTestsAllFail(t *testing.T) {
	t.Parallel()

	t.Run("Fail1", func(t *testing.T) {
		t.Fatalf("%s: This subtest always fails\n", packageName)
	})

	t.Run("Fail2", func(t *testing.T) {
		t.Fatalf("%s: This subtest always fails\n", packageName)
	})
}

func TestSubTestsSomeFail(t *testing.T) {
	t.Parallel()

	t.Run("Pass", func(t *testing.T) {
		t.Logf("%s: This subtest always passes\n", packageName)
	})

	t.Run("Fail", func(t *testing.T) {
		t.Fatalf("%s: This subtest always fails\n", packageName)
	})
}

func TestSubTestsSomePanic(t *testing.T) {
	t.Parallel()

	t.Run("Pass", func(t *testing.T) {
		t.Logf("%s: This subtest always passes\n", packageName)
	})

	t.Run("Panic", func(t *testing.T) {
		t.Skip("Skipping panics for now")
		panic("This subtest always panics")
	})
}

func TestFailInParentAfterSubTests(t *testing.T) {
	t.Parallel()

	t.Run("Pass1", func(t *testing.T) {
		t.Logf("%s: This subtest always passes\n", packageName)
	})

	t.Run("Pass2", func(t *testing.T) {
		t.Logf("%s: This subtest always passes\n", packageName)
	})

	t.Fatalf("%s: This test always fails\n", packageName)
}

func TestFailInParentBeforeSubTests(t *testing.T) {
	t.Parallel()

	t.Fatalf("%s: This test always fails\n", packageName) //nolint:revive

	t.Run("Pass1", func(t *testing.T) {
		t.Parallel()
		t.Logf("%s: This subtest always passes\n", packageName)
	})

	t.Run("Pass2", func(t *testing.T) {
		t.Parallel()
		t.Logf("%s: This subtest always passes\n", packageName)
	})
}

// This test should have a 50% pass ratio
func TestFlaky(t *testing.T) {
	t.Parallel()

	// Track if the test has run before
	stateFile := "tmp_test_flaky_state"

	// If the state file does not exist, create it and fail the test
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		if err := os.WriteFile(stateFile, []byte("run once"), 0644); err != nil { //nolint:gosec
			t.Fatalf("THIS IS UNEXPECTED: failed to create state file: %v", err)
		}
		t.Fatalf("This is a designed flaky test working as intended")
	} else {
		t.Cleanup(func() {
			err := os.Remove(stateFile)
			if err != nil {
				t.Fatalf("THIS IS UNEXPECTED: failed to remove state file: %v", err)
			}
		})
	}

	t.Logf("%s: This test passes after the first run\n", packageName)
}

func TestRandomFlaky(t *testing.T) {
	t.Parallel()
	random := rand.Intn(4)

	if random == 0 {
		t.Fatal("This test is designed to flake 1/4 of the time")
	}

	t.Log("This test is designed to pass 3/4 of the time")
}

func TestSkipped(t *testing.T) {
	t.Parallel()
	t.Skip("This test is intentionally skipped")
}

func TestPanic(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping panics for now")
	panic("This test intentionally panics")
}

func TestFlakyPanic(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping panics for now")

	// Track if the test has run before
	stateFile := "tmp_test_flaky_panic_state"

	// If the state file does not exist, create it and fail the test
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		if err := os.WriteFile(stateFile, []byte("run once"), 0644); err != nil { //nolint:gosec
			t.Fatalf("THIS IS UNEXPECTED: failed to create state file: %v", err)
		}
		panic("This is a designed flaky test panicking as intended")
	}
	t.Cleanup(func() {
		err := os.Remove(stateFile)
		if err != nil {
			t.Fatalf("THIS IS UNEXPECTED: failed to remove state file: %v", err)
		}
	})
}

func TestRace(t *testing.T) {
	t.Parallel()
	t.Logf("This test should trigger a failure if run with the -race flag, but otherwise pass")

	var (
		numGoroutines = 100
		sharedCounter int
		wg            sync.WaitGroup
	)

	worker := func(id int) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sharedCounter++
			_ = sharedCounter * id
		}
	}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(i)
	}

	wg.Wait()

	// Log the result
	t.Logf("Final value of sharedCounter: %d", sharedCounter)
}

func TestTimeout(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping timeout test for now")

	deadline, ok := t.Deadline()
	if !ok {
		log.Fatal("This test should have a deadline")
	}

	t.Logf("This test will sleep %s in order to timeout", time.Until(deadline).String())
	// Sleep until the deadline
	time.Sleep(time.Until(deadline))
	t.Logf("This test should have timed out")
}
