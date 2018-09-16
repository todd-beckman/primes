package timing

import (
	"time"
)

func GetBoolWithTiming(operation func() bool) (bool, time.Duration) {
	start := time.Now()
	result := operation()
	end := time.Now()

	elapsed := end.Sub(start)
	return result, elapsed
}
