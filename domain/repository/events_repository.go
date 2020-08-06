package repository

import "IkezawaYuki/craft/domain/entity"

type EventsRepository interface {
	Save(events *entity.Events) bool
}
