package model

type User struct {
	ID          string `db:"id"`
	StudentName string `db:"student_name"`
	Password    string `db:"password"`
	Graduate    bool   `db:"graduate"`
	Entered     bool   `db:"entered"`
}
