package entity

type LoginEntity struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
