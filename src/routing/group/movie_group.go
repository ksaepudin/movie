package group

import (
	db "github.com/ksaepudin/movie/src/repository/db"
	repo "github.com/ksaepudin/movie/src/repository/db/movie"
	service "github.com/ksaepudin/movie/src/service/movie"
	usecase "github.com/ksaepudin/movie/src/usecase/movie"
	"github.com/labstack/echo/v4"
)

// This function is to register a service into the service group
func MovieRoute(e *echo.Echo) {
	// This function is to register a service
	conn := db.ConnMovieDb
	movieRepo := repo.NewMovieRepository(conn)
	movieUc := usecase.NewMovieUsecase(movieRepo)
	movieService := service.NewMovieService(movieUc)

	// set group service
	grpMovie := e.Group("/movie")
	grpMovie.POST("", movieService.AddMovie)
	grpMovie.PATCH("", movieService.UpdateMovie)
	grpMovie.DELETE("", movieService.DeleteMovie)
	grpMovie.GET("", movieService.GetAllMovie)
	grpMovie.GET(":id", movieService.GetMovieByID)
}
