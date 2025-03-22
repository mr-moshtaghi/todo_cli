package contract

import "todo_cli/entity"

type UserWriteStore interface {
	Save(u entity.User)
}

type UserReadStore interface {
	Load() []entity.User
}
