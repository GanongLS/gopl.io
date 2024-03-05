// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {

	var dilbert = Employee{}
	var luki = Employee{Name: "Luki Subandi", ID: 13109043, Address: "Kebayoran lama selatan"}

	fmt.Printf("employee: %+v\n, %+v\n\n", luki, dilbert)

	luki.Salary += 5000 // promoted, for refactor few lines of code
	luki.Position = "Programmer"

	position := &luki.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	var employeeOfTheMonth *Employee = &luki
	employeeOfTheMonth.Position += " (proactive team player)"

	(*employeeOfTheMonth).Position += " (who guiding fellow employee)"

	fmt.Printf("employee: %+v\n, %+v\n\n", luki, dilbert)

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason

}

func EmployeeByID(id int) *Employee {
	var dilbertManager = Employee{Name: "Luki Subandi", ID: 13109043, Address: "Kebayoran lama selatan"}

	return &dilbertManager
}
