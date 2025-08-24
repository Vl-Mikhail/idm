package role

import "github.com/jmoiron/sqlx"

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

type RoleEntity struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// добавить новый элемент в коллекцию
func (r *RoleRepository) CreateRole(role RoleEntity) (roleId int64, err error) {
	err = r.db.Get(&roleId, "INSERT INTO role (name) VALUES ($1) returning id", role.Name)
	return roleId, err
}

// найти элемент коллекции по его id
func (r *RoleRepository) FindById(id int64) (role RoleEntity, err error) {
	err = r.db.Get(&role, "SELECT * FROM role WHERE id = $1", id)
	return role, err
}

// найти все элементы коллекции
func (r *RoleRepository) FindAll() (roles []RoleEntity, err error) {
	err = r.db.Select(&roles, "SELECT * FROM role")
	return roles, err
}

// найти слайс элементов коллекции по слайсу их id
func (r *RoleRepository) FindByIds(ids []int64) (roles []RoleEntity, err error) {
	query, args, err := sqlx.In("SELECT * FROM role WHERE id IN ($1);", ids)

	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.Select(&roles, query, args...)

	return roles, err
}

// удалить элемент коллекции по его id
func (r *RoleRepository) DeleteById(id int64) (roleId int64, err error) {
	err = r.db.Get(&roleId, "DELETE FROM role WHERE id = $1 returning id", id)
	return roleId, err
}

// удалить элементы по слайсу их id
func (r *RoleRepository) DeleteByIds(ids []int64) (roleIds []int64, err error) {
	query, args, err := sqlx.In("DELETE FROM role WHERE id IN ($1) RETURNING id", ids)

	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&roleIds, query, args...)

	return roleIds, err
}
