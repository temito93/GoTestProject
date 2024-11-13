package dto

import "goproject/internal/modules/user/entity"

type CreateUserDTO struct {
	Name  string `json:"name" binding:"required,max=50"`
	Email string `json:"email" binding:"required,email"`
}

func (dto *CreateUserDTO) ToUserEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Name:  dto.Name,
		Email: dto.Email,
	}
}
