package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phoneNumber"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	RoleID          uint   //`json:"role_id" validate:"required"`
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type UpdateRequest struct {
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Password        *string `json:"password"`
	ConfirmPassword *string `json:"confirm_password" `
	PhoneNumber     string  `json:"phone_number"`
	Username        string  `json:"username"`
	RoleID          uint    //`json:"role_id" validate:"required"`
}
