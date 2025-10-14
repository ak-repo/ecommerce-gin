package bannerdto

import "time"

// CreateBannerRequest - DTO for creating a banner
type CreateBannerRequest struct {
	Title       string `form:"title" binding:"required"` // matches <input name="title">
	Description string `form:"description"`              // matches <textarea name="description">
	IsActive    bool   `form:"-"`                // checkbox for active/inactive
	ImageURL    string `json:"-" form:"-"`               // handled separately after upload
}

// UpdateBannerRequest - DTO for updating a banner
type UpdateBannerRequest struct {
	ID          uint   `form:"id"` // hidden field in edit form
	Title       string `form:"title" binding:"omitempty"`
	Description string `form:"description"`
	IsActive    bool   `form:"is_active"`
	ImageURL    string `json:"-" form:"-"` // updated only if a new image is uploaded
}

// BannerResponse - DTO for returning banner data
type BannerResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
