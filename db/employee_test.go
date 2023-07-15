package db

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDB_CreateEmployee(t *testing.T) {
	db := establishDBConnection(t)

	t.Run("add new employee", func(t *testing.T) {
		depid, err := db.CreateDepartment("sample-department")
		require.Nil(t, err)

		sampleEmployee := &Employee{
			FirstName:    "Stephen",
			LastName:     "Brown",
			Username:     "sbrown",
			Password:     "qwerty",
			Email:        "sbrown1992@gmail.com",
			DOB:          time.Now(),
			Position:     "addnewemployee",
			DepartmentID: depid,
		}

		empid, err := db.CreateEmployee(sampleEmployee)
		require.Nil(t, err)
		assert.Greater(t, empid, 0)
	})

	t.Run("add new employee to non-existing department", func(t *testing.T) {
		db.ResetDB()

		employee := &Employee{
			FirstName:    "Stephen",
			DepartmentID: 1,
		}

		_, err := db.CreateEmployee(employee)
		assert.True(t, errors.Is(err, ErrDepartmentNotFound))
	})
}

func TestDB_GetEmployee(t *testing.T) {
	db := establishDBConnection(t)

	t.Run("get existing employee", func(t *testing.T) {
		depid, err := db.CreateDepartment("sample-department")
		require.Nil(t, err)

		dob := time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)
		employee := &Employee{
			FirstName:    "Stephen",
			LastName:     "Brown",
			Username:     "sbrown",
			Password:     "qwerty",
			Email:        "sbrown1992@gmail.com",
			DOB:          dob,
			Position:     "addnewemployee",
			DepartmentID: depid,
		}
		employeeid, err := db.CreateEmployee(employee)
		require.Nil(t, err)

		storedEmployee, err := db.GetEmployee(employeeid)
		require.Nil(t, err)

		assert.Equal(t, employeeid, storedEmployee.ID)
		assert.Equal(t, employee.LastName, storedEmployee.LastName)
		assert.Equal(t, employee.Username, storedEmployee.Username)
		assert.Equal(t, employee.Password, storedEmployee.Password)
		assert.Equal(t, employee.Email, storedEmployee.Email)
		assert.Equal(t, employee.DOB, storedEmployee.DOB)
		assert.Equal(t, employee.Position, storedEmployee.Position)
		assert.Equal(t, employee.DepartmentID, storedEmployee.DepartmentID)
	})

	t.Run("non-existing employee", func(t *testing.T) {
		db := establishDBConnection(t)
		err := db.ResetDB()
		require.Nil(t, err)

		_, err = db.GetEmployee(1)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})
}
