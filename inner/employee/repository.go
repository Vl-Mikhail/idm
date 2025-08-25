package employee

import (
	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

type EmployeeEntity struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// добавить новый элемент в коллекцию
func (r *EmployeeRepository) CreateEmployee(employee EmployeeEntity) (employeeId int64, err error) {
	err = r.db.Get(&employeeId, "INSERT INTO employee (name) VALUES ($1) returning id", employee.Name)
	return employeeId, err
}

// найти элемент коллекции по его id
func (r *EmployeeRepository) FindById(id int64) (employee EmployeeEntity, err error) {
	err = r.db.Get(&employee, "SELECT * FROM employee WHERE id = $1", id)
	return employee, err
}

// найти все элементы коллекции
func (r *EmployeeRepository) FindAll() (employees []EmployeeEntity, err error) {
	err = r.db.Select(&employees, "SELECT * FROM employee")
	return employees, err
}

// найти слайс элементов коллекции по слайсу их id
func (r *EmployeeRepository) FindByIds(ids []int64) (employees []EmployeeEntity, err error) {
	query, args, err := sqlx.In("SELECT * FROM employee WHERE id IN (?);", ids)

	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.Select(&employees, query, args...)

	return employees, err
}

// удалить элемент коллекции по его id
func (r *EmployeeRepository) DeleteById(id int64) (employeeId int64, err error) {
	err = r.db.Get(&employeeId, "DELETE FROM employee WHERE id = $1 returning id", id)
	return employeeId, err
}

// удалить элементы по слайсу их id
func (r *EmployeeRepository) DeleteByIds(ids []int64) (employeeIds []int64, err error) {
	query, args, err := sqlx.In("DELETE FROM employee WHERE id IN (?) RETURNING id", ids)

	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&employeeIds, query, args...)

	return employeeIds, err
}
