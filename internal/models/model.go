package model

import (
	"time"

	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	Name      string    `gorm:"name"`
	Vacancies []Vacancy `gorm:"many2many:gradesvacancies;"`
}

type Customer struct {
	gorm.Model
	Name      string
	Vacancies []Vacancy `gorm:"foreignKey:CustomerID"`
}

type Employee struct {
	gorm.Model
	FirstName      string
	LastName       string
	IsHr           string
	Vacancies      []Vacancy `gorm:"many2many:employeesvacancies;"`
	OwnedVacancies []Vacancy `gorm:"foreignKey:OwnerID"`
}

type Vacancy struct {
	gorm.Model
	Title       string      `json:"title"`
	RoleName    string      `json:"role_name"`
	Description string      `json:"description"`
	Salary      int         `json:"salary"`
	Country     string      `json:"country"`
	Region      string      `json:"region"`
	City        string      `json:"city"`
	IsActive    bool        `json:"is_active" gorm:"is_active"`
	Candidates  []Candidate `gorm:"many2many:candidatesvacancies;"`
	Grades      []Grade     `gorm:"many2many:gradesvacancies;"`
	Employees   []Employee  `gorm:"many2many:employeesvacancies;"`
	OwnerID     uint
	CustomerID  uint
	//TODO связать с Customer
}

type Candidate struct {
	gorm.Model
	LastName      string
	FirstName     string
	MiddleName    string
	Gender        int8
	BirthDate     time.Time
	Age           int8
	Country       string
	Region        string
	City          string
	Citizenship   string
	HasWorkpermit string
	Vacancies     []Vacancy `gorm:"many2many:candidatesvacancies;"`
}
