package todobusiness

import (
	"context"
	"errors"
	"testing"
	todomodel "todo-app/module/item/model"

	"github.com/stretchr/testify/mock"
)

type MockCreateTodoItemStorage struct {
	mock.Mock
}

func (m *MockCreateTodoItemStorage) CreateItem(ctx context.Context, data *todomodel.TodoItem) error {
	args := m.Called(ctx, data)

	return args.Error(0)
}

func TestCreateNewItemSuccess(t *testing.T) {

	mockCreateTodoItemStorage := new(MockCreateTodoItemStorage)
	var ctx context.Context
	data := todomodel.TodoItem{
		Title:  "1",
		Status: "Doing",
	}
	mockCall := mockCreateTodoItemStorage.On("CreateItem", ctx, &data).Return(nil)

	createBiz := NewCreateTodoItemBiz(mockCreateTodoItemStorage)

	result := createBiz.CreateNewItem(ctx, &data)

	if result != nil {
		t.Errorf("Expected result of CreateNewItem(ctx, data) to be nil, but got %d", result)
	}
	mockCall.Unset()
	mockCreateTodoItemStorage.AssertExpectations(t)
}

func TestCreateNewItemFail(t *testing.T) {

	mockCreateTodoItemStorage := new(MockCreateTodoItemStorage)
	var ctx context.Context
	data := todomodel.TodoItem{
		Title:  "1",
		Status: "Doing",
	}
	mockCall := mockCreateTodoItemStorage.On("CreateItem", ctx, &data).Return(errors.New("error"))

	createBiz := NewCreateTodoItemBiz(mockCreateTodoItemStorage)

	result := createBiz.CreateNewItem(ctx, &data)
	if result == nil {
		t.Errorf("Expected result of CreateNewItem(ctx, data) to be error, but got %d", result)
	}
	mockCall.Unset()
	mockCreateTodoItemStorage.AssertExpectations(t)
}
