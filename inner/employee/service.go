package employee

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	repo Repo
}

type Repo interface {
	CreateEmployee(employee Entity) (int64, error)
	FindById(id int64) (Entity, error)
	FindAll() ([]Entity, error)
	FindByIds(ids []int64) ([]Entity, error)
	DeleteById(id int64) (int64, error)
	DeleteByIds(ids []int64) ([]int64, error)
	BeginTransaction() (*sqlx.Tx, error)
	FindByNameTx(tx *sqlx.Tx, name string) (bool, error)
	CreateEmployeeTx(tx *sqlx.Tx, employee Entity) (int64, error)
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) CreateEmployee(name string) (int64, error) {
	tx, err := svc.repo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("creating employee panic: %v", r)
			errTx := tx.Rollback()
			if errTx != nil {
				err = fmt.Errorf("creating employee: rolling back transaction errors: %w, %w", err, errTx)
			}
		} else if err != nil {
			errTx := tx.Rollback()
			if errTx != nil {
				err = fmt.Errorf("creating employee: rolling back transaction errors: %w, %w", err, errTx)
			}
		} else {
			errTx := tx.Commit()
			if errTx != nil {
				err = fmt.Errorf("creating employee: commiting transaction error: %w", errTx)
			}
		}
	}()

	if err != nil {
		return 0, fmt.Errorf("error create employee: error creating transaction: %w", err)
	}

	isExist, err := svc.repo.FindByNameTx(tx, name)

	if err != nil {
		return 0, fmt.Errorf("error finding employee by name: %s, %w", name, err)
	}

	if isExist {
		return 0, fmt.Errorf("employee with name %s already exists", name)
	}

	newEmployeeId, err := svc.repo.CreateEmployeeTx(tx, Entity{Name: name})

	if err != nil {
		err = fmt.Errorf("error creating employee with name: %s %v", name, err)
	}

	return newEmployeeId, err
}

func (svc *Service) FindById(id int64) (Response, error) {
	var entity, err = svc.repo.FindById(id)
	if err != nil {
		return Response{}, fmt.Errorf("error finding employee with id %d: %w", id, err)
	}

	return entity.toResponse(), nil
}

func (svc *Service) FindAll() ([]Response, error) {
	var employees, err = svc.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error finding all employees: %w", err)
	}

	var responses []Response

	for index, employee := range employees {
		responses[index] = employee.toResponse()
	}

	return responses, nil
}

func (svc *Service) FindByIds(ids []int64) ([]Response, error) {
	var employees, err = svc.repo.FindByIds(ids)
	if err != nil {
		return nil, fmt.Errorf("error finding employees by ids %v: %w", ids, err)
	}

	var responses []Response

	for index, employee := range employees {
		responses[index] = employee.toResponse()
	}

	return responses, nil
}

func (svc *Service) DeleteById(id int64) (int64, error) {
	var employeeId, err = svc.repo.DeleteById(id)
	if err != nil {
		return 0, fmt.Errorf("error delete employee with id %d: %w", id, err)
	}

	return employeeId, nil
}

func (svc *Service) DeleteByIds(ids []int64) ([]int64, error) {
	var employeeIds, err = svc.repo.DeleteByIds(ids)
	if err != nil {
		return nil, fmt.Errorf("error delete employees by ids %v: %w", ids, err)
	}

	return employeeIds, nil
}
