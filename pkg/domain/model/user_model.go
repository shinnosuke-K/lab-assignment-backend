package model

type User struct {
	ID         string `db:"id"`
	StudentNum string `db:"student_num"`
	Password   string `db:"password"`
	Graduate   bool   `db:"graduate"`
	Entered    bool   `db:"entered"`
}
