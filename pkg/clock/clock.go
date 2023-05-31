package clock

import "time"

type Clock struct{}

func (clock Clock) GetCurrentTimestamp() time.Time {
	return time.Now()
}

func (clock Clock) GetZeroTime() time.Time {
	return time.Time{}
}
