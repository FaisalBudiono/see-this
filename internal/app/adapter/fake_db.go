package adapter

import (
	"FaisalBudiono/see-this/internal/app/port"
	"errors"
	"time"
)

type fakeDB struct {
	res []port.Resource
	idR int
	idF int
}

func (db *fakeDB) FindBySlug(slug string) (port.Resource, error) {
	for _, r := range db.res {
		if r.Slug == slug {
			return r, nil
		}
	}

	return port.Resource{}, errors.New("resource not found")
}

func (db *fakeDB) GetAllResources() ([]port.Resource, error) {
	return db.res, nil
}

func (db *fakeDB) SaveResource(name, slug string) (port.Resource, error) {
	r := db.newResource(name, slug)
	db.res = append(db.res, r)

	return r, nil
}

func (db *fakeDB) makeDummyResource() {
	db.res = []port.Resource{
		db.newResource("asd", "asd"),
		db.newResource("kambing", "keramas-kambing"),
	}
}

func (db *fakeDB) newResource(name, slug string) port.Resource {
	db.idR++
	return port.Resource{
		Id:   db.idR,
		Name: name,
		Slug: slug,
		Fields: []port.ResourceField{
			db.newField("wan", "string", "", false, true),
			db.newField("two", "string", "null", true, false),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (db *fakeDB) newField(
	name string,
	fieldType string,
	defaultValue string,
	isNullable bool,
	isArray bool,
) port.ResourceField {
	db.idF++
	return port.ResourceField{
		Id:            db.idF,
		FieldName:     name,
		ResourceRefId: 0,
		Type:          fieldType,
		IsNullable:    isNullable,
		IsArray:       isArray,
		Default:       defaultValue,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
