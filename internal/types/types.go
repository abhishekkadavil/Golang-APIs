package types

type GoApiUser struct {
	Id    int
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
