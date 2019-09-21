package entry

import (
	"time"
)

// Based off https://github.com/rupa/z/wiki/frecency
func Frecency(e *Entry) float64 {
	dx := int(time.Now().Unix()) - e.LastVisited
	if dx < 3600 {
		return float64(e.VisitedCount) * 4
	}
	if dx < 86400 {
		return float64(e.VisitedCount) * 2
	}
	if dx < 604800 {
		return float64(e.VisitedCount) / 2
	}
	return float64(e.VisitedCount) / 4
}
