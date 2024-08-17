package db

import (
	"errors"
	"time"
)

var (
	tempResources []Resource
	id            int
)

func init() {
	tempResources = makeDummyResource()
}

func GetAllResources() []Resource {
	return tempResources
}

func FindBySlug(slug string) (Resource, error) {
	for _, r := range tempResources {
		if r.Slug == slug {
			return r, nil
		}
	}

	return Resource{}, errors.New("resource not found")
}

func SaveResource(name, slug string) {
	tempResources = append(tempResources, newResource(name, slug))
}

func makeDummyResource() []Resource {
	return []Resource{
		newResource("asd", "asd"),
		newResource("kambing", "keramas-kambing"),
	}
}

func newResource(name, slug string) Resource {
	id++
	return Resource{
		Id:        id,
		Name:      name,
		Slug:      slug,
		Field:     make([]ResourceField, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
