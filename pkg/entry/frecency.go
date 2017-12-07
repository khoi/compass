package entry

import (
	"time"
)

// Based off https://github.com/rupa/z/wiki/frecency
func Frecency(e *Entry) int {
	dx := int(time.Now().Unix()) - e.LastVisited
	if dx < 3600 {
		return e.VisitedCount * 4
	}
	if dx < 86400 {
		return e.VisitedCount * 2
	}
	if dx < 604800 {
		return e.VisitedCount / 2
	}
	return e.VisitedCount / 4
}
