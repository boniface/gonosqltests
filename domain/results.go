package domain

import "time"

type Results struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Duration int64     `json:duration`
	Objects  int64     `json:"objects"`
}
