package model

type Article struct {
	Id    int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title string `json:"title"`
}

func (Article) TableName() string {
	return "article"
}
