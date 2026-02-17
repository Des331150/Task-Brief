package main

import "time"

type CurrentStage string

const (
	Queued = "queued"
	Done   = "done"
)

type Jobs struct {
	ID            int64
	AudioID       int64
	CurrentStatus CurrentStage
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
