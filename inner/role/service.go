package role

import (
	"fmt"
)

type Service struct {
	repo Repo
}

type Repo interface {
	CreateRole(role Entity) (int64, error)
	FindById(id int64) (Entity, error)
	FindAll() ([]Entity, error)
	FindByIds(ids []int64) ([]Entity, error)
	DeleteById(id int64) (int64, error)
	DeleteByIds(ids []int64) ([]int64, error)
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) CreateRole(name string) (int64, error) {
	var roleId, err = svc.repo.CreateRole(Entity{Name: name})
	if err != nil {
		return 0, fmt.Errorf("error create role with name %s: %w", name, err)
	}

	return roleId, nil
}

func (svc *Service) FindById(id int64) (Response, error) {
	var entity, err = svc.repo.FindById(id)
	if err != nil {
		return Response{}, fmt.Errorf("error finding role with id %d: %w", id, err)
	}

	return entity.toResponse(), nil
}

func (svc *Service) FindAll() ([]Response, error) {
	var roles, err = svc.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error finding all roles: %w", err)
	}

	var responses []Response

	for index, role := range roles {
		responses[index] = role.toResponse()
	}

	return responses, nil
}

func (svc *Service) FindByIds(ids []int64) ([]Response, error) {
	var roles, err = svc.repo.FindByIds(ids)
	if err != nil {
		return nil, fmt.Errorf("error finding roles by ids %v: %w", ids, err)
	}

	var responses []Response

	for index, role := range roles {
		responses[index] = role.toResponse()
	}

	return responses, nil
}

func (svc *Service) DeleteById(id int64) (int64, error) {
	var roleId, err = svc.repo.DeleteById(id)
	if err != nil {
		return 0, fmt.Errorf("error delete role with id %d: %w", id, err)
	}

	return roleId, nil
}

func (svc *Service) DeleteByIds(ids []int64) ([]int64, error) {
	var roleIds, err = svc.repo.DeleteByIds(ids)
	if err != nil {
		return nil, fmt.Errorf("error delete roles by ids %v: %w", ids, err)
	}

	return roleIds, nil
}
