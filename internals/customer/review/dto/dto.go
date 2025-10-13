package custreviewdto

import "time"

type CreateReviewRequest struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Rating    uint   `json:"rating" binding:"required,min=1,max=5"`
	Comment   string `json:"comment"`
	UserID    uint   `json:"-"`
}

type ReviewResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	UserID    uint      `json:"user_id"`
	Rating    uint      `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
