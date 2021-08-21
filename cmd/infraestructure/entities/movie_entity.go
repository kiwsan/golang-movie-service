package entities

type MovieEntity struct {
	Id       int64 `gorm:"primaryKey"`
	Name     string
	Director string
	Writer   string
	Stars    string
}

func (MovieEntity) TableName() string {
	return "movies"
}
