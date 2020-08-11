package repository

import "IkezawaYuki/craft/domain/entity"

type SignalEventRepository interface {
	Save(events *entity.SignalEvent) bool
	GetByCount(loadEvents int) *entity.SignalEvents
}
