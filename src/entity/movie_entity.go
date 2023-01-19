package entity

type Movie struct {
	Id          int     `json:"id" gorm:"column:id" query:"id"`
	Title       string  `json:"title" query:"title" gorm:"column:title"`
	Description string  `json:"description" query:"description" gorm:"column:description"`
	Rating      float32 `json:"rating" query:"rating" gorm:"column:rating"`
	Image       string  `json:"image" query:"image" gorm:"column:image"`
	CreatedAt   string  `json:"created_at" query:"created_at" gorm:"column:created_at"`
	UpdatedAt   string  `json:"updated_at" query:"updated_at" gorm:"column:updated_at"`
}

func (Movie) TableName() string {
	return "tbl_movie"
}
