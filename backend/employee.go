package backend

import (
	"database/sql"
	"fmt"
)

// Employee struct (represents a row in the database)
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}

// EmployeeService struct (used for Wails binding)
type EmployeeService struct {
	DB *sql.DB
}

// NewEmployeeService - Create a new instance of EmployeeService
func NewEmployeeService(db *sql.DB) *EmployeeService {
	return &EmployeeService{DB: db}
}

// AddEmployee - Add a new employee
func (e *EmployeeService) AddEmployee(name string, age int, position string, salary int) error {
	query := "INSERT INTO employees (name, age, position, salary) VALUES (?, ?, ?, ?)"
	_, err := e.DB.Exec(query, name, age, position, salary)
	if err != nil {
		return err
	}
	fmt.Println("Employee added successfully")
	return nil
}

// GetEmployees - Fetch all employees
func (e *EmployeeService) GetEmployees() ([]Employee, error) {
	rows, err := e.DB.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Position, &emp.Salary)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

// GetEmployee - Fetch a single employee by ID
func (e *EmployeeService) GetEmployee(id int) (Employee, error) {
	var emp Employee
	query := "SELECT * FROM employees WHERE id = ?"
	err := e.DB.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Position, &emp.Salary)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

// UpdateEmployee - Update employee details
func (e *EmployeeService) UpdateEmployee(id int, name string, age int, position string, salary int) error {
	query := "UPDATE employees SET name = ?, age = ?, position = ?, salary = ? WHERE id = ?"
	_, err := e.DB.Exec(query, name, age, position, salary, id)
	if err != nil {
		return err
	}
	fmt.Println("Employee updated successfully")
	return nil
}

// DeleteEmployee - Remove an employee
func (e *EmployeeService) DeleteEmployee(id int) error {
	query := "DELETE FROM employees WHERE id = ?"
	_, err := e.DB.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println("Employee deleted successfully")
	return nil
}
