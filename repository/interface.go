package repository

type Model interface{}

type Interface interface {
	Save(model Model) error
	Find(query, model Model) (err error)
	FindAll(model Model) (err error)
	FindById(id int, model Model) (err error)
	Delete(id int) error
}
