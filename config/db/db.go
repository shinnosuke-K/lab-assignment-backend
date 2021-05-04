package db

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() *gorm.DB {
	//dsn := "root:@tcp(lab-db:3306)/questionnaire?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(localhost:3306)/questionnaire?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := sqlDB.Ping(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return db
}

func InsertLabs(db *gorm.DB) error {
	file, err := os.Open("lab.csv")
	if err != nil {
		return err
	}

	csvFile := csv.NewReader(file)
	records, err := csvFile.ReadAll()
	if err != nil {
		return err
	}

	type Prof struct {
		ID        string
		LabName   string
		ProfName  string
		ProfRoman string
	}

	for _, r := range records {

		prof := Prof{
			ID:        uuid.New().String(),
			LabName:   r[0],
			ProfName:  r[1],
			ProfRoman: r[2],
		}
		if err := db.Create(&prof).Error; err != nil {
			return err
		}
	}

	return nil
}

func InsertStudent(db *gorm.DB) error {
	file, err := os.Open("student.csv")
	if err != nil {
		return err
	}

	csvFile := csv.NewReader(file)
	records, err := csvFile.ReadAll()
	if err != nil {
		return err
	}

	type User struct {
		ID         string
		StudentNum int64
		Password   string
		Graduate   bool
		Entered    bool
	}

	for _, r := range records {

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

		student := User{
			ID:         uuid.NewString(),
			StudentNum: num,
			Password:   hashedPass,
			Graduate:   graduate,
			Entered:    entered,
		}

		if err := db.Create(&student).Error; err != nil {
			return err
		}
	}
	return nil
}
