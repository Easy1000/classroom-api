package main

import "time"

// Define Student model
type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" gorm:"unique"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}

// Define Class model
type Class struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ClassName   string `json:"class_name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	TeacherName string `json:"teacher_name"`
}

// Define Enrollment model
type Enrollment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	StudentID      uint      `json:"student_id"`
	ClassID        uint      `json:"class_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`

	// Relationships
	Student Student `json:"student" gorm:"foreignKey:StudentID"`
	Class   Class   `json:"class" gorm:"foreignKey:ClassID"`
}

// Define Comment model
type Comment struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
