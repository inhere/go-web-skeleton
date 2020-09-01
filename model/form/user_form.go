package form

// CreateUserForm create new user
type CreateUserForm struct {
	// Username user name
	Username string `json:"username" form:"username" validate:"required|min:1|max:2" example:"demo"`
	// Accept user licence
	Accept uint `json:"accept" form:"accept" validate:"min=1|max=2" example:"1"`
	// Email user email
	Email string `json:"email" form:"email" validate:"required|email"`
	// Age user age
	Age uint `json:"age" form:"age" validate:"required|min:1|max:99"`
	// Password user password
	Password int64 `json:"password" form:"password" validate:"required|max:32"`
}
