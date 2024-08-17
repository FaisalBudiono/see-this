package db

import (
	"errors"
	"time"
)

var (
	tempResources []Resource
	idR           int
	idF           int
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
	idR++
	return Resource{
		Id:   idR,
		Name: name,
		Slug: slug,
		Fields: []ResourceField{
			newField("wan", "string", "", false, true),
			newField("two", "string", "null", true, false),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func newField(
	name string,
	fieldType string,
	defaultValue string,
	isNullable bool,
	isArray bool,
) ResourceField {
	idF++
	return ResourceField{
		Id:            idF,
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
