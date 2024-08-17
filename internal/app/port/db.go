package port

type DB interface {
	FindBySlug(slug string) (Resource, error)
	GetAllResources() ([]Resource, error)
	SaveResource(name, slug string) (Resource, error)
}
