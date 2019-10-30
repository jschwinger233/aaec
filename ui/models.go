package ui

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type EventKind interface {
	Parse([]byte) (interface{}, error)
}

type Foreground struct{}
type Disable struct{}

type ForegroundData struct {
	Package string
}

type DisableData struct {
	Package string
	Once    bool
}

func (k Foreground) Parse(bytes []byte) (interface{}, error) {
	data := ForegroundData{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (k Disable) Parse(bytes []byte) (interface{}, error) {
	data := DisableData{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

type ResultStatus string

const (
	Created   ResultStatus = "created"
	Completed ResultStatus = "completed"
	Failed    ResultStatus = "failed"
)

type EventId uuid.UUID

type UnixTime int64

type Event struct {
	Id        EventId
	CreatedAt UnixTime
	Kind      EventKind
	Content   interface{}
}

func NewEventId(id string) (EventId, error) {
	uid, err := uuid.FromBytes([]byte(id))
	return EventId(uid), err
}

func NewUnixTime(second int) (UnixTime, error) {
	return UnixTime(second), nil
}

func NewEventKind(kind string) (EventKind, error) {
	kinds := map[string]EventKind{
		"foreground": Foreground{},
		"disable":    Disable{},
	}
	if kind, ok := kinds[kind]; ok {
		return kind, nil
	}
	return nil, errors.New("unknown event kind")
}

func NewEvent(id string, created int, kind string, content []byte) (event Event, err error) {
	eventId, err := NewEventId(id)
	if err != nil {
		logger.Errorf("invalid event id: %v", err)
		return
	}

	createdAt, err := NewUnixTime(created)
	if err != nil {
		logger.Errorf("invalid unix timestamp: %v", err)
		return
	}

	eventKind, err := NewEventKind(kind)
	if err != nil {
		logger.Errorf("invalid event kind: %v", err)
		return
	}

	eventContent, err := eventKind.Parse(content)
	if err != nil {
		logger.Errorf("invalid event content: %v", err)
		return
	}

	event = Event{
		Id:        eventId,
		CreatedAt: createdAt,
		Kind:      eventKind,
		Content:   eventContent,
	}
	return
}

type Result struct {
	EventId
	Status ResultStatus
	Data   []byte
}
