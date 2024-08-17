package adapter

import (
	"FaisalBudiono/see-this/internal/app/port"
)

func NewDB() port.DB {
	db := new(fakeDB)
	db.makeDummyResource()

	return db
}
