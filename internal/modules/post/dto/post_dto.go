package dto

import "goproject/internal/modules/post/entity"

type CreatePostDTO struct {
	Title  string `json:"title" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

func (dto *CreatePostDTO) ToPostEntity() *entity.PostEntity {
	return &entity.PostEntity{
		Title:  dto.Title,
		UserID: dto.UserID,
	}
}
