package db

// Model interface which all models should implement
type Model interface {
	Save(Model) (Model, error)
	Delete() (bool, error)
}
