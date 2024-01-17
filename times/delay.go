package times

import (
	"math/rand"
	"time"
)

func RandomDelay(delay time.Duration) {
	time.Sleep(time.Duration(rand.Int63n(int64(delay))))
}
