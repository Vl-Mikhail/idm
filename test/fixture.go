package tests

import (
	"idm/inner/employee"
	"idm/inner/role"

	"github.com/jmoiron/sqlx"
)

func CreateEmployeeTestTables(db *sqlx.DB) {
	var schema = `
		create table if not exists employee
		(
			id         bigint primary key generated always as identity,
			name       text                                  not null,
			created_at timestamptz default current_timestamp not null,
			updated_at timestamptz default current_timestamp not null
		);
	`

	db.MustExec(schema)
}

func CreateRoleTestTables(db *sqlx.DB) {
	var schema = `
		create table if not exists role
		(
			id         bigint primary key generated always as identity,
			name       text                                  not null,
			created_at timestamptz default current_timestamp not null,
			updated_at timestamptz default current_timestamp not null
		);
	`

	db.MustExec(schema)
}

type EmployeeFixture struct {
	employees *employee.EmployeeRepository
}

func NewEmployeeFixture(employees *employee.EmployeeRepository) *EmployeeFixture {
	return &EmployeeFixture{employees}
}

func (f *EmployeeFixture) Employee(name string) int64 {
	var entity = employee.EmployeeEntity{
		Name: name,
	}
	var newId, err = f.employees.CreateEmployee(entity)
	if err != nil {
		panic(err)
	}
	return newId
}

type RoleFixture struct {
	role *role.RoleRepository
}

func NewRoleFixture(role *role.RoleRepository) *RoleFixture {
	return &RoleFixture{role}
}

func (f *RoleFixture) Role(name string) int64 {
	var entity = role.RoleEntity{
		Name: name,
	}
	var newId, err = f.role.CreateRole(entity)
	if err != nil {
		panic(err)
	}
	return newId
}
