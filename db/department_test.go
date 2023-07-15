package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDB_CreateDepartment(t *testing.T) {
	db := establishDBConnection(t)

	t.Run("add new department", func(t *testing.T) {
		id, err := db.CreateDepartment("testindustries-addnewdepartment")
		require.Nil(t, err)
		assert.Greater(t, id, 0)
	})
}

func TestDB_GetDepartment(t *testing.T) {
	db := establishDBConnection(t)

	t.Run("existing department", func(t *testing.T) {
		depname := "testindustries-getexistingdepartment"
		id, err := db.CreateDepartment(depname)
		require.Nil(t, err)

		dep, err := db.GetDepartment(id)
		require.Nil(t, err)

		assert.Equal(t, depname, dep.Name)
		assert.Equal(t, id, dep.ID)
	})

	t.Run("non-existing department", func(t *testing.T) {
		db := establishDBConnection(t)
		err := db.ResetDB()
		require.Nil(t, err)

		_, err = db.GetDepartment(1)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})
}
