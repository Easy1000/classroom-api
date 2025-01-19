package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const apiURL = "https://jsonplaceholder.typicode.com/comments"

func seedDatabase() {
	// Create initial classes
	classes := []Class{
		{
			ClassName:   "Math 101",
			StartTime:   "09:00 AM",
			EndTime:     "10:30 AM",
			TeacherName: "Mr. Johnson",
		},
		{
			ClassName:   "Science 101",
			StartTime:   "11:00 AM",
			EndTime:     "12:30 PM",
			TeacherName: "Dr. Smith",
		},
	}

	// Insert classes and track their IDs
	for i, class := range classes {
		if err := db.Create(&class).Error; err != nil {
			log.Printf("Error seeding class %s: %v\n", class.ClassName, err)
		} else {
			classes[i].ID = class.ID
		}
	}

	// Create initial students
	students := []Student{
		{
			FirstName:      "Alice",
			LastName:       "Anderson",
			Email:          "alice.anderson@example.com",
			DateOfBirth:    time.Date(2005, 6, 15, 0, 0, 0, 0, time.UTC),
			EnrollmentDate: time.Now(),
		},
		{
			FirstName:      "Bob",
			LastName:       "Brown",
			Email:          "bob.brown@example.com",
			DateOfBirth:    time.Date(2004, 9, 25, 0, 0, 0, 0, time.UTC),
			EnrollmentDate: time.Now(),
		},
	}

	// Insert students and track their IDs
	for i, student := range students {
		if err := db.Create(&student).Error; err != nil {
			log.Printf("Error seeding student %s: %v\n", student.Email, err)
		} else {
			students[i].ID = student.ID
		}
	}

	// Create initial enrollments
	enrollments := []Enrollment{
		{
			StudentID:      students[0].ID, // Alice is enrolled in Math 101
			ClassID:        classes[0].ID,
			EnrollmentDate: time.Now(),
		},
		{
			StudentID:      students[1].ID, // Bob is enrolled in Science 101
			ClassID:        classes[1].ID,
			EnrollmentDate: time.Now(),
		},
	}

	// Insert enrollments
	for _, enrollment := range enrollments {
		if err := db.Create(&enrollment).Error; err != nil {
			log.Printf("Error seeding enrollment (StudentID: %d, ClassID: %d): %v\n", enrollment.StudentID, enrollment.ClassID, err)
		}
	}

	comments, err := FetchComments()
	if err != nil {
		log.Fatalf("Error fetching comments: %v", err)
	}

	for _, comment := range comments {
		if err := db.Create(&comment).Error; err != nil {
			log.Printf("Error inserting comment (ID: %d): %v", comment.ID, err)
		}
	}

	log.Println("Database seeding completed successfully!")
}

// FetchComments fetches comments data from the external API
func FetchComments() ([]Comment, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return comments, nil
}
