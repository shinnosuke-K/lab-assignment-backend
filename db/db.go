package db

import (
	"encoding/csv"
	"os"

	"github.com/google/uuid"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "root:@tcp(lab-db:3306)/questionnaire?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InsertLabs(db *gorm.DB) error {
	file, err := os.Open("list.csv")
	if err != nil {
		return err
	}

	csvFile := csv.NewReader(file)
	records, err := csvFile.ReadAll()
	if err != nil {
		return err
	}

	//id varchar(255) not null ,
	//	lab_name varchar(255),
	//	prof_name varchar(255),
	//	prof_roman varchar(255),

	type Lab struct {
		ID        string
		LabName   string
		ProfName  string
		ProfRoman string
	}

	for _, r := range records {

		prof := Lab{
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

	type Student struct {
		UserID     string
		StudentNum int64
		Password   string
		Entered    bool
	}

	for _, r := range records {

		num, err := strconv.ParseInt(r[0], 10, 64)
		if err != nil {
			return err
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(r[1]), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		entered, err := strconv.ParseBool(r[2])
		if err != nil {
			return err
		}

		student := Student{
			UserID:     uuid.NewString(),
			StudentNum: num,
			Password:   string(hash),
			Entered:    entered,
		}

		if err := db.Create(&student).Error; err != nil {
			return err
		}
	}
	return nil
}
