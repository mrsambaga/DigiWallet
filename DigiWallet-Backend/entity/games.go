package entity

type Boxes struct {
	BoxId uint64  `gorm:"PrimaryKey" json:"box_id" binding:"required"`
	Prize float64 `json:"prize"`
}

type Leaderboard struct {
	UserId       uint64 `gorm:"PrimaryKey"`
	Name         string
	TotalWinning float64
}

type Chance struct {
	UserId uint64 `gorm:"PrimaryKey"`
	Chance int
}
