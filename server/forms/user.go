package forms

type UserForm struct{}

type RegisterForm struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type UpdateForm struct {
// 	Email string `json:"email" binding:"required"`
// 	Name  string `json:"name" binding:"required"`
// }

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
