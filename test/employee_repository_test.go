package tests

import (
	"idm/inner/database"
	"idm/inner/employee"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository(t *testing.T) {
	a := assert.New(t)
	var db = database.ConnectDb(".env_test")
	CreateEmployeeTestTables(db)

	var clearDatabase = func() {
		db.MustExec("delete from employee")
	}
	defer func() {
		if r := recover(); r != nil {
			clearDatabase()
		}
	}()
	var employeeRepository = employee.NewRepository(db)
	var fixture = NewEmployeeFixture(employeeRepository)

	t.Run("find an employee by id", func(t *testing.T) {
		var newEmployeeId = fixture.Employee("Test Name")

		got, err := employeeRepository.FindById(newEmployeeId)

		a.Nil(err)
		a.NotEmpty(got)
		a.NotEmpty(got.Id)
		a.NotEmpty(got.CreatedAt)
		a.NotEmpty(got.UpdatedAt)
		a.Equal("Test Name", got.Name)
		clearDatabase()
	})

	t.Run("find all employees", func(t *testing.T) {
		fixture.Employee("Test Name")
		got, err := employeeRepository.FindAll()

		a.Nil(err)
		a.NotEmpty(got)
		a.Len(got, 1)
		clearDatabase()
	})

	t.Run("find an employees by ids", func(t *testing.T) {
		var newEmployeeId = fixture.Employee("Test Name")

		var emplIds = []int64{newEmployeeId}

		got, err := employeeRepository.FindByIds(emplIds)

		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(newEmployeeId, got[0].Id)
		clearDatabase()
	})

	t.Run("find an employees by ids", func(t *testing.T) {
		var newEmployeeId = fixture.Employee("Test Name")

		var emplIds = []int64{newEmployeeId}

		got, err := employeeRepository.FindByIds(emplIds)

		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(newEmployeeId, got[0].Id)
		clearDatabase()
	})

	t.Run("Delete an employee by id", func(t *testing.T) {
		var newEmployeeId = fixture.Employee("Test Name")

		got, err := employeeRepository.DeleteById(newEmployeeId)

		a.Nil(err)
		a.Equal(newEmployeeId, got)

		got2, err := employeeRepository.FindAll()

		a.Nil(err)
		a.Empty(got2)

		clearDatabase()
	})

	t.Run("Delete an employee by ids", func(t *testing.T) {
		var newEmployeeId = fixture.Employee("Test Name")

		got, err := employeeRepository.DeleteByIds([]int64{newEmployeeId})

		a.Nil(err)
		a.Equal(newEmployeeId, got[0])

		got2, err := employeeRepository.FindAll()

		a.Nil(err)
		a.Empty(got2)

		clearDatabase()
	})

	t.Run("Create an employee with tx", func(t *testing.T) {
		tx, err := employeeRepository.BeginTransaction()
		emplId, err := employeeRepository.CreateEmployeeTx(tx, employee.Entity{Name: "Test Name"})

		a.Nil(err)

		err = tx.Commit()

		a.Nil(err)
		a.NotEmpty(emplId)

		got, err := employeeRepository.FindById(emplId)

		a.Nil(err)
		a.Equal("Test Name", got.Name)

		clearDatabase()
	})

}
