package movie

import (
	"github.com/ksaepudin/movie/src/entity"
	"github.com/ksaepudin/movie/src/repository/db"
	"gorm.io/gorm"
)

type MovieRepository interface {
	AddMovie(in interface{}) (interface{}, error)
	UpdateMovie(in interface{}) (interface{}, error)
	DeleteMovie(id int) (interface{}, error)
	GetAllMovie() (interface{}, error)
	GetMovieByID(id int) (interface{}, error)
	SearchMovie(search string) (interface{}, error)
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(conn db.DbDriver) MovieRepository {
	return &movieRepository{
		db: conn.Db().(*gorm.DB),
	}
}

func (m *movieRepository) AddMovie(in interface{}) (interface{}, error) {
	req := in.(*entity.Movie)

	err := m.db.Debug().Model(&entity.Movie{}).Create(req).Error
	if err != nil {
		return nil, err
	}
	return "Sukses Insert", nil
}
func (m *movieRepository) UpdateMovie(in interface{}) (interface{}, error) {
	req := in.(*entity.Movie)

	err := m.db.Debug().Model(&entity.Movie{}).Where("id = ?", req.Id).Updates(&entity.Movie{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		UpdatedAt:   req.UpdatedAt,
	}).Error
	if err != nil {
		return nil, err
	}
	return "Sukses Update", nil
}
func (m *movieRepository) DeleteMovie(id int) (interface{}, error) {
	err := m.db.Debug().Model(&entity.Movie{}).Where("id = ?", id).Delete(id).Error
	if err != nil {
		return nil, err
	}

	return "Sukses Deleted", nil
}

func (m *movieRepository) GetAllMovie() (interface{}, error) {
	var res []*entity.Movie

	err := m.db.Model(&entity.Movie{}).Debug().Order("id ASC").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *movieRepository) SearchMovie(search string) (interface{}, error) {
	var res []*entity.Movie

	err := m.db.Model(&entity.Movie{}).Debug().Where("(title like ? or description like ? or rating like ?)", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *movieRepository) GetMovieByID(id int) (interface{}, error) {
	var res *entity.Movie

	err := m.db.Debug().Model(res).Where("id = ?", id).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
