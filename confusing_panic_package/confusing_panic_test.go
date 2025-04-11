package confusingpanicpackage

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInnocent(t *testing.T) {
	t.Parallel()

	for i := range 1000 {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			sleepTime := rand.Intn(100)
			t.Logf("%s: sleeping for %d ms", t.Name(), sleepTime)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		})
	}

	t.Log("This test is innocent")
}

func TestPanic(t *testing.T) {
	t.Parallel()

	sleepTime := rand.Intn(100)
	t.Logf("%s: sleeping for %d ms", t.Name(), sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	panic("This test is designed to panic")
}

