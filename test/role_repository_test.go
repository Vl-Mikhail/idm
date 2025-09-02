package tests

import (
	"idm/inner/database"
	"idm/inner/role"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRepository(t *testing.T) {
	a := assert.New(t)
	var db = database.ConnectDb(".env_test")
	CreateRoleTestTables(db)

	var clearDatabase = func() {
		db.MustExec("delete from role")
	}
	defer func() {
		if r := recover(); r != nil {
			clearDatabase()
		}
	}()
	var roleRepository = role.NewRoleRepository(db)
	var fixture = NewRoleFixture(roleRepository)

	t.Run("find an role by id", func(t *testing.T) {
		var newRoleId = fixture.Role("Test Name")

		got, err := roleRepository.FindById(newRoleId)

		a.Nil(err)
		a.NotEmpty(got)
		a.NotEmpty(got.ID)
		a.NotEmpty(got.CreatedAt)
		a.NotEmpty(got.UpdatedAt)
		a.Equal("Test Name", got.Name)
		clearDatabase()
	})

	t.Run("find all roles", func(t *testing.T) {
		fixture.Role("Test Name")
		got, err := roleRepository.FindAll()

		a.Nil(err)
		a.NotEmpty(got)
		a.Len(got, 1)
		clearDatabase()
	})

	t.Run("find an roles by ids", func(t *testing.T) {
		var newRoleId = fixture.Role("Test Name")

		var roleIds = []int64{newRoleId}

		got, err := roleRepository.FindByIds(roleIds)

		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(newRoleId, got[0].ID)
		clearDatabase()
	})

	t.Run("find an roles by ids", func(t *testing.T) {
		var newRoleId = fixture.Role("Test Name")

		var roleIds = []int64{newRoleId}

		got, err := roleRepository.FindByIds(roleIds)

		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(newRoleId, got[0].ID)
		clearDatabase()
	})

	t.Run("Delete an role by id", func(t *testing.T) {
		var newRoleId = fixture.Role("Test Name")

		got, err := roleRepository.DeleteById(newRoleId)

		a.Nil(err)
		a.Equal(newRoleId, got)

		got2, err := roleRepository.FindAll()

		a.Nil(err)
		a.Empty(got2)

		clearDatabase()
	})

	t.Run("Delete an role by ids", func(t *testing.T) {
		var newRoleId = fixture.Role("Test Name")

		got, err := roleRepository.DeleteByIds([]int64{newRoleId})

		a.Nil(err)
		a.Equal(newRoleId, got[0])

		got2, err := roleRepository.FindAll()

		a.Nil(err)
		a.Empty(got2)

		clearDatabase()
	})

}
