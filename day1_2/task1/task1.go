package day1_2

import (
	"fmt"
)

func averageGrade(grade map[string]float32, totalSubject int) float32 {
	var totalScore float32
	for _, score := range grade {
		totalScore += score
	}
	average := totalScore / float32(totalSubject)
	return average
}

func Task1() {
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Println("Enter the total number of subjects:")
	var total int
	fmt.Scan(&total)

	gradeMap := make(map[string]float32)

	for i := 0; i < total; i++ {
		fmt.Println("Enter the subject name:")
		var subject string
		fmt.Scan(&subject)

		fmt.Println("Enter the grade:")
		var grade float32
		fmt.Scan(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
			i-- // Decrement i to repeat the current iteration
			continue
		}

		gradeMap[subject] = grade
	}

	fmt.Printf("Name: %s\n", name)
	for key, value := range gradeMap {
		fmt.Printf("Subject: %s, Grade: %.2f\n", key, value)
	}
	fmt.Printf("Average grade is: %.2f\n", averageGrade(gradeMap, total))
}
