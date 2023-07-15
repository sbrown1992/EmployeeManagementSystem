package db

import (
	"fmt"
	"time"
)

type Employee struct {
	ID           int
	FirstName    string
	LastName     string
	Username     string
	Password     string
	Email        string
	DOB          time.Time
	Position     string
	DepartmentID int
	Department   Department
}

func (Employee) TableName() string {
	return "EMS.employee"
}

func (db *DB) CreateEmployee(employee *Employee) (int, error) {
	_, err := db.GetDepartment(employee.DepartmentID)
	if err != nil {
		return -1, fmt.Errorf("%w: %d", ErrDepartmentNotFound, employee.DepartmentID)
	}

	result := db.conn.Create(employee)
	if result.Error != nil {
		return -1, result.Error
	}

	return employee.ID, nil
}

func (db *DB) GetEmployee(id int) (*Employee, error) {
	employee := &Employee{}
	result := db.conn.First(&employee, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return employee, nil
}
