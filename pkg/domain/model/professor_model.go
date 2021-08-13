package model

type Professor struct {
	ID        string `db:"id"`
	LabName   string `db:"lab_name"`
	ProfName  string `db:"prof_name"`
	ProfRoman string `db:"prof_roman"`
}
