package dto

type BoxRequestDTO struct {
	BoxId uint64 `json:"box_id" binding:"required"`
}
