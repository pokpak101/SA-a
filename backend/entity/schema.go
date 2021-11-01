package entity

import (
	"time"

	//"golang.org/x/text/date"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Password string
	Case []Case `gorm:"foreignKey:OwnerID"`
}

type Case struct {
	gorm.Model
	Name	string `gorm:"uniqueIndex"`
	CaseTime    time.Time
	OwnerID           *uint8
	Owner             User
	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:CaseID"`
	//State   []State  `gorm:"foreignKey:CaseID"`
	//Symptom []Symptom `gorm:"foreignKey:CaseID"`
}

type Symptom struct {
	gorm.Model
	SymptomData string `gorm:"uniqueIndex"`
	//CaseID	*uint
	//Case	Case
	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:SymptomID"`
}

type State struct {
	gorm.Model
	StateData         string `gorm:"uniqueIndex"`
	//CaseID	*uint
	//Case	Case
	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:StateID"`
}

type Assess struct {
	gorm.Model
	AssessData        string `gorm:"uniqueIndex"`
	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:AssessID"`
}

type AssessmentSheet struct {
	gorm.Model
	AssessTime time.Time

	CaseID	*uint
	Case	Case `gorm:"references:id"`

	SymptomID   *uint
	Symptom     Symptom	`gorm:"references:id"`

	StateID *uint
	State   State	`gorm:"references:id"`

	AssessID *uint
	Assess   Assess	`gorm:"references:id"`
}
