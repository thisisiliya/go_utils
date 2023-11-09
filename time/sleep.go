package time

import (
	"math/rand"
	"time"
)

func RandomSleep(min, max int) {

	randomDelay := rand.Intn(max-min) + min

	time.Sleep(time.Duration(randomDelay) * time.Second)
}
