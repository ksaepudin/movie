package movie

import (
	"errors"
	"time"

	"github.com/ksaepudin/movie/src/entity"
	repo "github.com/ksaepudin/movie/src/repository/db/movie"
)

type MovieUsecase interface {
	AddMovie(in interface{}) (interface{}, error)
	UpdateMovie(in interface{}) (interface{}, error)
	DeleteMovie(id int) (interface{}, error)
	GetAllMovie(search string) (interface{}, error)
	GetMovieByID(id int) (interface{}, error)
}

type movieUsecase struct {
	movieRepo repo.MovieRepository
}

func NewMovieUsecase(movieRepo repo.MovieRepository) MovieUsecase {
	return &movieUsecase{
		movieRepo: movieRepo,
	}
}

func (m *movieUsecase) AddMovie(in interface{}) (interface{}, error) {
	if in == nil {
		return nil, errors.New("Request Cannot be nil")
	}
	req := in.(*entity.Movie)
	if req.CreatedAt == "" {
		req.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	res, err := m.movieRepo.AddMovie(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *movieUsecase) UpdateMovie(in interface{}) (interface{}, error) {
	if in == nil {
		return nil, errors.New("Request Cannot be nil")
	}
	req := in.(*entity.Movie)
	if req.UpdatedAt == "" {
		req.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	res, err := m.movieRepo.UpdateMovie(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *movieUsecase) DeleteMovie(id int) (interface{}, error) {
	if id == 0 {
		return nil, errors.New("Request Cannot be 0")
	}
	res, err := m.movieRepo.DeleteMovie(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *movieUsecase) GetAllMovie(search string) (interface{}, error) {
	var (
		resRepo interface{}
		res     []*entity.Movie
		err     error
	)

	if search != "" {
		resRepo, err = m.movieRepo.SearchMovie(search)
	} else {
		resRepo, err = m.movieRepo.GetAllMovie()
	}

	if err != nil {
		return nil, err
	}

	res = resRepo.([]*entity.Movie)
	return res, nil
}

func (m *movieUsecase) GetMovieByID(id int) (interface{}, error) {
	res, err := m.movieRepo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
