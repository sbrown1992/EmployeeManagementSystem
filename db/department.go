package db

import "errors"

type Department struct {
	ID   int
	Name string
}

func (Department) TableName() string {
	return "EMS.department"
}

func (db *DB) CreateDepartment(name string) (int, error) {
	dep := &Department{
		Name: name,
	}

	result := db.conn.Create(dep)
	if result.Error != nil {
		return -1, result.Error
	}

	return dep.ID, nil
}

func (db *DB) GetDepartment(id int) (*Department, error) {
	department := &Department{}
	result := db.conn.First(&department, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return department, nil
}

// **********
// * Errors *
// **********

// TODO replace get department error with this
var ErrDepartmentNotFound = errors.New("department not found")
