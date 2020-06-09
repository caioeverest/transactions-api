package repository

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func (m *Mock) Save(model Model) error             { return m.Called(model).Error(0) }
func (m *Mock) Find(query, model Model) error      { return m.Called(query, model).Error(0) }
func (m *Mock) FindAll(model Model) error          { return m.Called(model).Error(0) }
func (m *Mock) FindById(id int, model Model) error { return m.Called(id, model).Error(0) }
func (m *Mock) Delete(id int) error                { return m.Called(id).Error(0) }
