package reviewdto

import "time"

// For admin listing
type AdminReviewResponse struct {
	ID          uint      `json:"id"`
	ProductID   uint      `json:"product_id"`
	ProductName string    `json:"product_name"`
	UserID      uint      `json:"user_id"`
	UserName    string    `json:"user_name"`
	Rating      uint      `json:"rating"`
	Comment     string    `json:"comment"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// For admin approval request
type ReviewStatusUpdateRequest struct {
	Status string `json:"status" binding:"required,oneof=APPROVED REJECTED"`
}
