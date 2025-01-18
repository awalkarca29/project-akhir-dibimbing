package helper

import "project-akhir-awal/entity"

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}

type UsersFormatter struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	RoleID  int         `json:"role_id"`
	Roles   entity.Role `json:"role"`
	Photo   string      `json:"photo"`
	Address string      `json:"address"`
	Phone   string      `json:"phone"`
}

func FormatUserDetail(user entity.User) UsersFormatter {
	userFormatter := UsersFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.RoleID = user.RoleID
	userFormatter.Roles = user.Roles
	userFormatter.Photo = user.Photo
	userFormatter.Address = user.Address
	userFormatter.Phone = user.Phone

	return userFormatter
}

func FormatUsers(users []entity.User) []UsersFormatter {
	usersFormatter := []UsersFormatter{}

	for _, user := range users {
		userFormatter := FormatUserDetail(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}

type UserUpdateFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
	// Roles   entity.Role `json:"role"`
	Photo   string `json:"photo"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func FormatUserUpdate(user entity.User) UserUpdateFormatter {
	userUpdateFormatter := UserUpdateFormatter{}
	userUpdateFormatter.ID = user.ID
	userUpdateFormatter.Name = user.Name
	userUpdateFormatter.Email = user.Email
	userUpdateFormatter.RoleID = user.RoleID
	// userUpdateFormatter.Roles = user.Roles
	userUpdateFormatter.Photo = user.Photo
	userUpdateFormatter.Address = user.Address
	userUpdateFormatter.Phone = user.Phone

	return userUpdateFormatter
}
