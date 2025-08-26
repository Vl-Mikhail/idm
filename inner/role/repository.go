package role

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// добавить новый элемент в коллекцию
func (r *Repository) CreateRole(role Entity) (roleId int64, err error) {
	err = r.db.Get(&roleId, "INSERT INTO role (name) VALUES ($1) returning id", role.Name)
	return roleId, err
}

// найти элемент коллекции по его id
func (r *Repository) FindById(id int64) (role Entity, err error) {
	err = r.db.Get(&role, "SELECT * FROM role WHERE id = $1", id)
	return role, err
}

// найти все элементы коллекции
func (r *Repository) FindAll() (roles []Entity, err error) {
	err = r.db.Select(&roles, "SELECT * FROM role")
	return roles, err
}

// найти слайс элементов коллекции по слайсу их id
func (r *Repository) FindByIds(ids []int64) (roles []Entity, err error) {
	query, args, err := sqlx.In("SELECT * FROM role WHERE id IN (?);", ids)

	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.Select(&roles, query, args...)

	return roles, err
}

// удалить элемент коллекции по его id
func (r *Repository) DeleteById(id int64) (roleId int64, err error) {
	err = r.db.Get(&roleId, "DELETE FROM role WHERE id = $1 returning id", id)
	return roleId, err
}

// удалить элементы по слайсу их id
func (r *Repository) DeleteByIds(ids []int64) (roleIds []int64, err error) {
	query, args, err := sqlx.In("DELETE FROM role WHERE id IN (?) RETURNING id", ids)

	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&roleIds, query, args...)

	return roleIds, err
}
