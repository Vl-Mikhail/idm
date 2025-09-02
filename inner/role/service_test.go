package role

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateRole(role Entity) (int64, error) {
	args := m.Called(role)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRepo) FindById(id int64) (role Entity, err error) {
	args := m.Called(id)
	return args.Get(0).(Entity), args.Error(1)
}

func (m *MockRepo) FindAll() (role []Entity, err error) {
	args := m.Called()
	return args.Get(0).([]Entity), args.Error(1)
}

func (m *MockRepo) FindByIds(ids []int64) (role []Entity, err error) {
	args := m.Called(ids)
	return args.Get(0).([]Entity), args.Error(1)
}

func (m *MockRepo) DeleteById(id int64) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRepo) DeleteByIds(ids []int64) ([]int64, error) {
	args := m.Called(ids)
	return args.Get(0).([]int64), args.Error(1)
}

type MockRepoStub struct {
}

func (m *MockRepoStub) CreateRole(role Entity) (int64, error) {
	panic("implement me")
}

func (m *MockRepoStub) FindAll() ([]Entity, error) {
	panic("implement me")
}

func (m *MockRepoStub) FindByIds(ids []int64) ([]Entity, error) {
	panic("implement me")
}

func (m *MockRepoStub) DeleteById(id int64) (int64, error) {
	panic("implement me")
}

func (m *MockRepoStub) DeleteByIds(ids []int64) ([]int64, error) {
	panic("implement me")
}

func (m *MockRepoStub) FindById(id int64) (role Entity, err error) {

	var mockEntity = Entity{
		Id:        1,
		Name:      "John Doe",
		CreatedAt: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC),
	}

	return mockEntity, nil
}

func TestFindById(t *testing.T) {

	// создаём экземпляр объекта с ассерт-функциями
	var a = assert.New(t)

	t.Run("should return found role", func(t *testing.T) {

		var repo = new(MockRepo)
		var svc = NewService(repo)

		var entity = Entity{
			Id:        1,
			Name:      "John Doe",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		var want = entity.toResponse()

		repo.On("FindById", int64(1)).Return(entity, nil)

		var got, err = svc.FindById(1)

		a.Nil(err)
		a.Equal(want, got)
		a.True(repo.AssertNumberOfCalls(t, "FindById", 1))
	})

	t.Run("should return wrapped error", func(t *testing.T) {

		var repo = new(MockRepo)
		var svc = NewService(repo)

		var entity = Entity{}
		var err = errors.New("database error")
		var want = fmt.Errorf("error finding role with id 1: %w", err)

		repo.On("FindById", int64(1)).Return(entity, err)

		var response, got = svc.FindById(1)

		a.Empty(response)
		a.NotNil(got)
		a.Equal(want, got)
		a.True(repo.AssertNumberOfCalls(t, "FindById", 1))
	})

	t.Run("should return role stub test", func(t *testing.T) {
		var repo = new(MockRepoStub)
		var svc = NewService(repo)

		var entity = Entity{
			Id:        1,
			Name:      "John Doe",
			CreatedAt: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC),
		}

		var want = entity.toResponse()
		var got, err = svc.FindById(1)

		a.Nil(err)
		a.Equal(want, got)
	})
}
