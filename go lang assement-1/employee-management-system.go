package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

// Add Employee
func addEmployee(id int, name string, age int, department string) error {
	// Check for unique ID
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}
	
	// Validate age
	if age <= 18 {
		return errors.New("age must be greater than 18")
	}
	
	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}

// Search Employee by ID or Name
func searchEmployee(query string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprint(emp.ID) == query || strings.EqualFold(emp.Name, query) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// List Employees by Department
func listEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			result = append(result, emp)
		}
	}
	return result
}

// Count Employees by Department
func countEmployeesByDepartment(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	// Constants for Departments
	const (
		HR = "HR"
		IT = "IT"
	)

	// Add employees
	_ = addEmployee(1, "Alice", 25, HR)
	_ = addEmployee(2, "Bob", 30, IT)
	_ = addEmployee(3, "Charlie", 28, HR)

	// Search Employee
	if emp, err := searchEmployee("2"); err == nil {
		fmt.Println("Employee found:", *emp)
	} else {
		fmt.Println(err)
	}

	// List Employees by Department
	fmt.Println("Employees in HR:", listEmployeesByDepartment(HR))

	// Count Employees by Department
	fmt.Printf("Number of employees in IT: %d\n", countEmployeesByDepartment(IT))
}
