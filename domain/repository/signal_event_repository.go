package repository

import (
	"IkezawaYuki/craft/domain/entity"
	"time"
)

type SignalEventRepository interface {
	Save(events *entity.SignalEvent) bool
	GetByCount(loadEvents int) *entity.SignalEvents
	GetSignalEventsAfterTime(timeTime time.Time) *entity.SignalEvents
}
