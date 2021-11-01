package entity

import (
	//"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Assess{},
		&User{},
		&Symptom{},
		&State{},
		&Case{},
		&AssessmentSheet{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "pokpak",
		Email:    "pokpak@gmail.com",
		Password: string(password),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(password),
	})

	var pokpak User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "pokpak@gmail.com").Scan(&pokpak)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// --- Assess Data
	Assess1 := Assess{
		AssessData:  "normal",
	}
	db.Model(&Assess{}).Create(&Assess1)

	Assess2 := Assess{
		AssessData:  "emergency",
	}
	db.Model(&Assess{}).Create(&Assess2)

	// State Data
	State1 := State{
		StateData: "unconscious",
	}
	db.Model(&State{}).Create(&State1)

	State2 := State{
		StateData: "conscious",
	}
	db.Model(&State{}).Create(&State2)

	// Symptom Data
	Symptom1 := Symptom{
		SymptomData: "emergency",
	}
	db.Model(&Symptom{}).Create(&Symptom1)

	Symptom2 := Symptom{
		SymptomData: "covid",
	}
	db.Model(&Symptom{}).Create(&Symptom2)

	// --- Case Data
	Case1 := Case{
		Name:  "นายa",
		Owner: pokpak,
	}
	db.Model(&Case{}).Create(&Case1)

	Case2 := Case{
		Name:  "นายb",
		Owner: pokpak,
	}
	db.Model(&Case{}).Create(&Case2)

	// ams 1
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet{
		Case: 		Case1,
		Symptom:    Symptom1,
		State:      State1,
		Assess: 	Assess1,
		AssessTime: time.Now(),
	})
	// ams 2
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet{
		Case: 		Case2,
		Symptom:    Symptom2,
		State:      State2,
		Assess: 	Assess2,
		AssessTime: time.Now(),
	})


}
