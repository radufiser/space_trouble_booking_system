package domain

import "time"

type Launch struct {
	LaunchpadId string
	Name        string
	Date        time.Time
}
