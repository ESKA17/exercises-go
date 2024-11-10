package problem7

import (
	"fmt"
	"math/rand"
	"time"
)

//TODO: identify the data race
// fix the issue.

func task() {
	start := time.Now()
	time.AfterFunc(
		randomDuration(), func() {
			fmt.Println(time.Now().Sub(start))
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
