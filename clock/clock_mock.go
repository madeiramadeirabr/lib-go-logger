package clock

import "time"

type Mock struct{}

func (clockMock Mock) GetCurrentTimestamp() time.Time {
	return time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
}

func (clockMock Mock) GetZeroTime() time.Time {
	return time.Time{}
}
