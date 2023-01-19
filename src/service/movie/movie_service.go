package movie

import (
	"net/http"
	"strconv"

	"github.com/ksaepudin/movie/src/entity"
	usecase "github.com/ksaepudin/movie/src/usecase/movie"
	"github.com/labstack/echo/v4"
)

type MovieService interface {
	AddMovie(c echo.Context) error
	UpdateMovie(c echo.Context) error
	DeleteMovie(c echo.Context) error
	GetAllMovie(c echo.Context) error
	GetMovieByID(c echo.Context) error
}

type movie struct {
	movieUc usecase.MovieUsecase
}

func NewMovieService(uc usecase.MovieUsecase) MovieService {
	return &movie{movieUc: uc}
}

func (m *movie) AddMovie(c echo.Context) error {
	req := new(entity.Movie)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	res, err := m.movieUc.AddMovie(req)
	if err != nil {
		return c.JSON(http.StatusPreconditionFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (m *movie) UpdateMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	req := new(entity.Movie)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	req.Id = id

	res, err := m.movieUc.UpdateMovie(req)
	if err != nil {
		return c.JSON(http.StatusPreconditionFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (m *movie) DeleteMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	res, err := m.movieUc.DeleteMovie(id)
	if err != nil {
		return c.JSON(http.StatusPreconditionFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (m *movie) GetAllMovie(c echo.Context) error {
	search := c.QueryParam("search")
	res, err := m.movieUc.GetAllMovie(search)
	if err != nil {
		return c.JSON(http.StatusPreconditionFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func (m *movie) GetMovieByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	res, err := m.movieUc.GetMovieByID(id)
	if err != nil {
		return c.JSON(http.StatusPreconditionFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
