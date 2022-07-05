package utils

import "Oenone/model"

func ToUserDto(user model.User) model.UserDto {
	return model.UserDto{
		ID:       string(user.ID[:]),
		Username: user.Username,
	}
}
