package models

type LoginRequest struct {
	Username string `extensions:"x-order=0" json:"username" validate:"required"`
	Password string `extensions:"x-order=1" json:"password" validate:"required"`
}
