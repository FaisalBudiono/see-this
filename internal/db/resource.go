package db

import "time"

var (
	tempResources []Resource
	id            int
)

func Init() {
	tempResources = makeDummyResource()
}

func GetAllResources() []Resource {
	return tempResources
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
