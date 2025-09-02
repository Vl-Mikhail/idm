package employee

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// добавить новый элемент в коллекцию
func (r *Repository) CreateEmployee(employee Entity) (employeeId int64, err error) {
	err = r.db.Get(&employeeId, "INSERT INTO employee (name) VALUES ($1) returning id", employee.Name)
	return employeeId, err
}

// найти элемент коллекции по его id
func (r *Repository) FindById(id int64) (employee Entity, err error) {
	err = r.db.Get(&employee, "SELECT * FROM employee WHERE id = $1", id)
	return employee, err
}

// найти все элементы коллекции
func (r *Repository) FindAll() (employees []Entity, err error) {
	err = r.db.Select(&employees, "SELECT * FROM employee")
	return employees, err
}

// найти слайс элементов коллекции по слайсу их id
func (r *Repository) FindByIds(ids []int64) (employees []Entity, err error) {
	query, args, err := sqlx.In("SELECT * FROM employee WHERE id IN (?);", ids)

	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.Select(&employees, query, args...)

	return employees, err
}

// удалить элемент коллекции по его id
func (r *Repository) DeleteById(id int64) (employeeId int64, err error) {
	err = r.db.Get(&employeeId, "DELETE FROM employee WHERE id = $1 returning id", id)
	return employeeId, err
}

// удалить элементы по слайсу их id
func (r *Repository) DeleteByIds(ids []int64) (employeeIds []int64, err error) {
	query, args, err := sqlx.In("DELETE FROM employee WHERE id IN (?) RETURNING id", ids)

	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&employeeIds, query, args...)

	return employeeIds, err
}
