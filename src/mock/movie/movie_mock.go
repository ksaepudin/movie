package movie

import (
	"github.com/stretchr/testify/mock"
)

type MockMovieRepo struct {
	mock.Mock
}

func (m *MockMovieRepo) AddMovie(in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
func (m *MockMovieRepo) UpdateMovie(in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
func (m *MockMovieRepo) DeleteMovie(id int) (interface{}, error) {
	call := m.Called(id)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
func (m *MockMovieRepo) GetAllMovie() (interface{}, error) {
	call := m.Called()
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
func (m *MockMovieRepo) GetMovieByID(id int) (interface{}, error) {
	call := m.Called(id)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
func (m *MockMovieRepo) SearchMovie(search string) (interface{}, error) {
	call := m.Called(search)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
