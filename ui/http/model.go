package http

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Type      string `json:"type"`
	Content   []byte `json:"content"`
}

func NewEvent(typ string, content []byte) *Event {
	return &Event{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Unix(),
		Type:      typ,
		Content:   []byte(content),
	}
}

type Events struct {
	Events []*Event `json:"events"`
}

func (es Events) AddEvent(e *Event) {
	es.Events = append(es.Events, e)
}
