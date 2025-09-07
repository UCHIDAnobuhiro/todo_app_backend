package usecase_test

import (
	"testing"
	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"
	"todo_backend/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepo struct{ mock.Mock }

func (m *MockTodoRepo) FindByUser(userID uint) ([]domain.Todo, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Todo), args.Error(1)
}

func (m *MockTodoRepo) Create(todo domain.Todo) error {
	return m.Called(todo).Error(0)
}

func (m *MockTodoRepo) Update(todo domain.Todo) error {
	return m.Called(todo).Error(0)
}

func (m *MockTodoRepo) Delete(userID uint, id int) error {
	return m.Called(userID, id).Error(0)
}

var _ repository.TodoRepository = (*MockTodoRepo)(nil)

func TestGetTodos_CallsRepoWithUserID_AndReturnsList(t *testing.T) {
	mockRepo := new(MockTodoRepo)
	uc := usecase.NewTodoUsecase(mockRepo)

	expected := []domain.Todo{
		{ID: 1, UserID: 42, Title: "Test Todo", Completed: false},
	}
	mockRepo.On("FindByUser", uint(42)).Return(expected, nil)

	todos, err := uc.GetTodos(42)
	assert.NoError(t, err)
	assert.Equal(t, expected, todos)

	mockRepo.AssertExpectations(t)
}

func TestGetTodo_PropagatesErrorFromRepo(t *testing.T) {
	repo := new(MockTodoRepo)
	uc := usecase.NewTodoUsecase(repo)

	userID := uint(7)
	expected := []domain.Todo{
		{ID: 1, Title: "A", Completed: false, UserID: userID},
		{ID: 2, Title: "B", Completed: true, UserID: userID},
	}
	repo.On("FindByUser", userID).Return(expected, nil).Once()

	got, err := uc.GetTodos(userID)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	repo.AssertExpectations(t)
}

func TestAddTodo_CallsCreate(t *testing.T) {
	// given
	repo := new(MockTodoRepo)
	uc := usecase.NewTodoUsecase(repo)

	in := domain.Todo{ID: 0, Title: "new", Completed: false, UserID: 1}
	repo.On("Create", in).Return(nil).Once()

	// when
	err := uc.AddTodo(in)

	// then
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateTodo_CallsUpdate(t *testing.T) {
	// given
	repo := new(MockTodoRepo)
	uc := usecase.NewTodoUsecase(repo)

	in := domain.Todo{ID: 10, Title: "edited", Completed: true, UserID: 1}
	repo.On("Update", in).Return(nil).Once()

	// when
	err := uc.UpdateTodo(in)

	// then
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteTodo_CallsDeleteWithUserAndID(t *testing.T) {
	// given
	repo := new(MockTodoRepo)
	uc := usecase.NewTodoUsecase(repo)

	userID := uint(1)
	id := 10
	repo.On("Delete", userID, id).Return(nil).Once()

	// when
	err := uc.DeleteTodo(userID, id)

	// then
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
