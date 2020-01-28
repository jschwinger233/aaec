package http

import (
	"time"

	"github.com/google/uuid"
)

type Meta struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Package   string `json:"package"`
}

func NewMeta(pkg string) *Meta {
	return &Meta{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Unix(),
		Package:   pkg,
	}
}

type Event struct {
	Meta
	Type string `json:"type"`
}

func NewEvent(pkg string, typ string) *Event {
	return &Event{
		Meta: *NewMeta(pkg),
		Type: typ,
	}
}

type Instruction struct {
	Meta
	Extra map[string]string `json:"extra"`
}

func NewInstruction(pkg string, extra map[string]string) *Instruction {
	return &Instruction{
		Meta:  *NewMeta(pkg),
		Extra: extra,
	}
}
