package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Enter the birthYear")
		return
	}
	birthYear, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid Year")
		return
	}
	birthMonth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid Month")
		return
	}
	birthDay, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid Day")
		return
	}
	birthDate := time.Date(
		birthYear,
		time.Month(birthMonth),
		birthDay,
		0, 0, 0, 0,
		time.Local,
	)

	now := time.Now()

	duration := now.Sub(birthDate)

	days := int(duration.Hours() / 24)

	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())

	age := currentYear - birthYear
	totalMonths := age*12 + (currentMonth - birthMonth)

	fmt.Println("Age in year:", age)
	fmt.Println("Age in month:", totalMonths)
	fmt.Println("Age in days:", days)

}
