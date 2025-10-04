package types


type User struct {
	Id string `json:"id"`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}