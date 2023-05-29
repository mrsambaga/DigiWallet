package entity

type Sources struct {
	SourceId uint64 `gorm:"PrimaryKey"`
	Name     string
}
