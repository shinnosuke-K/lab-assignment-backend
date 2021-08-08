package db

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
)

var Driver repository.DBConnector

func Open() {
	//dsn := "root:@tcp(lab-db:3306)/questionnaire?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(localhost:3306)/questionnaire?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	Driver.DB, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := Driver.Ping(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func InsertLabs(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	csvFile := csv.NewReader(file)
	records, err := csvFile.ReadAll()
	if err != nil {
		return err
	}

	type Prof struct {
		ID        string `db:"id"`
		LabName   string `db:"lab_name"`
		ProfName  string `db:"prof_name"`
		ProfRoman string `db:"prof_roman"`
	}

	profs := make([]*Prof, len(records))
	for n, r := range records {
		profs[n] = &Prof{
			ID:        uuid.New().String(),
			LabName:   r[0],
			ProfName:  r[1],
			ProfRoman: r[2],
		}
	}

	q := `
			insert into questionnaire.professors(id, lab_name, prof_name, prof_roman)
			values(:id,:lab_name,:prof_name,:prof_roman)
		
		`
	if _, err := Driver.NamedExec(q, profs); err != nil {
		return err
	}

	return nil
}

func InsertUsers(path string) error {

	file, err := os.Open(path)
	if err != nil {

		return err
	}
	defer file.Close()

	csvFile := csv.NewReader(file)
	records, err := csvFile.ReadAll()
	if err != nil {
		return err
	}

	type User struct {
		ID         string `db:"id"`
		StudentNum int64  `db:"student_num"`
		Password   string `db:"password"`
		Graduate   bool   `db:"graduate"`
		Entered    bool   `db:"entered"`
	}

	users := make([]User, len(records))

	for n, r := range records {

		num, err := strconv.ParseInt(r[0], 10, 64)
		if err != nil {
			return err
		}

		h := sha256.New()
		h.Write([]byte(r[1]))
		hashedPass := fmt.Sprintf("%x", h.Sum(nil))

		graduate, err := strconv.ParseBool(r[2])
		if err != nil {
			return err
		}

		entered, err := strconv.ParseBool(r[3])
		if err != nil {
			return err
		}

		users[n] = User{
			ID:         uuid.NewString(),
			StudentNum: num,
			Password:   hashedPass,
			Graduate:   graduate,
			Entered:    entered,
		}
	}

	q := `
			insert into questionnaire.users(id, student_num, password, graduate, entered)
			values(:id, :student_num, :password, :graduate, :entered)
		
		`
	if _, err := Driver.NamedExec(q, users); err != nil {
		return err
	}

	return nil
}
