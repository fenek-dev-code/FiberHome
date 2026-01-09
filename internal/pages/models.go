package pages

import (
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type RegisterFormModel struct {
	Username string
	Email    string
	Password string
}

type IRegisterFormModel interface {
	FormValue(key string, defaultValue ...string) string
}

func NewRegisterFormModel(form IRegisterFormModel) *RegisterFormModel {
	return &RegisterFormModel{
		Username: form.FormValue("username"),
		Email:    form.FormValue("email"),
		Password: form.FormValue("password"),
	}
}

func (f *RegisterFormModel) Validate() *validate.Errors {
	// Add validation logic here
	return validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: f.Email, Message: "Email не может быть пустым"},
		&validators.StringIsPresent{Name: "Username", Field: f.Username, Message: "Имя пользователя не может быть пустым"},
		&validators.StringIsPresent{Name: "Password", Field: f.Password, Message: "Пароль не может быть пустым"},
	)
}
