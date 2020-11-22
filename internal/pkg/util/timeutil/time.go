package timeutil

import "time"

func TimePointer() *time.Time {
	t := time.Now()
	return &t
}
