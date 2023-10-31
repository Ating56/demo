package dto

import "my_blog_back/model"

type UserDto struct {
	ID   uint   `json:"ID"`
	Name string `json:"name"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		ID:   user.ID,
		Name: user.Name,
	}
}
