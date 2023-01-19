package movie

import (
	"errors"
	"testing"

	"github.com/ksaepudin/movie/src/entity"
	mockRepo "github.com/ksaepudin/movie/src/mock/movie"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestAddMovie(t *testing.T) {
	movieRepo := &mockRepo.MockMovieRepo{}
	req := &entity.Movie{
		Id:          0,
		Title:       "test",
		Description: "test",
		Rating:      0,
		Image:       "test",
		CreatedAt:   "test",
		UpdatedAt:   "test",
	}
	res := &entity.Movie{
		Id:          9999,
		Title:       "succsess",
		Description: "succsess",
		Rating:      999,
		Image:       "succsess",
		CreatedAt:   "succsess",
		UpdatedAt:   "succsess",
	}
	Convey("Test Usecase Add Movie", t, func() {
		Convey("negative scenarios", func() {
			Convey("request nil scenarios", func() {
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.AddMovie(nil)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("AddMovie some error scenarios", func() {
				movieRepo.On("AddMovie", mock.Anything).Return(nil, errors.New("some error")).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.AddMovie(req)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positif scenarios", func() {
			Convey("CreateAt not empety scenarios", func() {
				movieRepo.On("AddMovie", mock.Anything).Return(res, nil).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.AddMovie(req)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
			Convey("CreateAt empety scenarios", func() {
				req.CreatedAt = ""
				movieRepo.On("AddMovie", mock.Anything).Return(res, nil).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.AddMovie(req)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})

	})
}

func TestUpdateMovie(t *testing.T) {
	movieRepo := &mockRepo.MockMovieRepo{}
	req := &entity.Movie{
		Id:          0,
		Title:       "test",
		Description: "test",
		Rating:      0,
		Image:       "test",
		CreatedAt:   "test",
		UpdatedAt:   "test",
	}
	res := &entity.Movie{
		Id:          9999,
		Title:       "succsess",
		Description: "succsess",
		Rating:      999,
		Image:       "succsess",
		CreatedAt:   "succsess",
		UpdatedAt:   "succsess",
	}
	Convey("Test Usecase Update Movie", t, func() {
		Convey("negative scenarios", func() {
			Convey("request nil scenarios", func() {
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.UpdateMovie(nil)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("UpdateMovie some error scenarios", func() {
				movieRepo.On("UpdateMovie", mock.Anything).Return(nil, errors.New("some error")).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.UpdateMovie(req)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positif scenarios", func() {
			Convey("CreateAt not empety scenarios", func() {
				movieRepo.On("UpdateMovie", mock.Anything).Return(res, nil).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.UpdateMovie(req)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
			Convey("CreateAt empety scenarios", func() {
				req.UpdatedAt = ""
				movieRepo.On("UpdateMovie", mock.Anything).Return(res, nil).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.UpdateMovie(req)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})

	})
}

func TestGetAllMovie(t *testing.T) {
	movieRepo := &mockRepo.MockMovieRepo{}
	req := "search"
	res := []*entity.Movie{
		&entity.Movie{
			Id:          9999,
			Title:       "succsess",
			Description: "succsess",
			Rating:      999,
			Image:       "succsess",
			CreatedAt:   "succsess",
			UpdatedAt:   "succsess",
		},
	}
	var resInterface interface{}
	resInterface = res
	Convey("Test Usecase Update Movie", t, func() {
		Convey("negative scenarios", func() {
			Convey("SearchMovie some error scenarios", func() {
				movieRepo.On("SearchMovie", mock.Anything).Return(nil, errors.New("some error")).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.GetAllMovie(req)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("GetAllMovie some error scenarios", func() {
				movieRepo.On("GetAllMovie").Return(nil, errors.New("some error")).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.GetAllMovie("")
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive scenarios", func() {
			Convey("SearchMovie  scenarios", func() {
				movieRepo.On("SearchMovie", mock.Anything).Return(resInterface, nil).Once()
				uc := NewMovieUsecase(movieRepo)
				resp, err := uc.GetAllMovie(req)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
			// Convey("GetAllMovie  scenarios", func() {
			// 	movieRepo.On("GetAllMovie").Return(res, nil).Once()
			// 	uc := NewMovieUsecase(movieRepo)
			// 	resp, err := uc.GetAllMovie("")
			// 	So(resp, ShouldNotBeNil)
			// 	So(err, ShouldBeNil)
			// })
		})
	})
}

// func TestDeleteMovie(t *testing.T) {
// 	movieRepo := &mockRepo.MockMovieRepo{}
// 	req := 0
// 	// res := "succsess"

// 	Convey("Test Usecase Update Movie", t, func() {
// 		Convey("negative scenarios", func() {
// 			Convey("request 0 scenarios", func() {
// 				uc := NewMovieUsecase(movieRepo)
// 				resp, err := uc.DeleteMovie(0)
// 				So(resp, ShouldBeNil)
// 				So(err, ShouldNotBeNil)
// 			})
// 			Convey("DeleteMovie some error scenarios", func() {
// 				movieRepo.On("DeleteMovie").Return(nil, errors.New("some error")).Once()
// 				uc := NewMovieUsecase(movieRepo)
// 				resp, err := uc.DeleteMovie(req)
// 				So(resp, ShouldBeNil)
// 				So(err, ShouldNotBeNil)
// 			})
// 		})
// 		Convey("Positif scenarios", func() {
// 			Convey("CreateAt empety scenarios", func() {
// 				movieRepo.On("DeleteMovie", mock.Anything).Return(res, nil).Once()
// 				uc := NewMovieUsecase(movieRepo)
// 				resp, err := uc.DeleteMovie(req)
// 				So(resp, ShouldNotBeNil)
// 				So(err, ShouldBeNil)
// 			})
// 		})

// 	})
// }
