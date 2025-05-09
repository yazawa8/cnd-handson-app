package test

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockProjectRepository はProjectRepositoryインターフェースのモック実装
type MockProjectRepository struct {
	mock.Mock
}

// GetProjectByID のモック
func (m *MockProjectRepository) GetProjectByID(projectID uuid.UUID) (*model.Project, error) {
	args := m.Called(projectID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), args.Error(1)
}

// GetProjectsByOwnerID のモック
func (m *MockProjectRepository) GetProjectsByOwnerID(ownerID uuid.UUID) ([]*model.Project, error) {
	args := m.Called(ownerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Project), args.Error(1)
}

// ListProjects のモック
func (m *MockProjectRepository) ListProjects() ([]*model.Project, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Project), args.Error(1)
}

// CreateProject のモック
func (m *MockProjectRepository) CreateProject(project *model.Project) (*model.Project, error) {
	args := m.Called(project)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), args.Error(1)
}

// UpdateProject のモック
func (m *MockProjectRepository) UpdateProject(project *model.Project) (*model.Project, error) {
	args := m.Called(project)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), args.Error(1)
}

// DeleteProject のモック
func (m *MockProjectRepository) DeleteProject(projectID uuid.UUID) error {
	args := m.Called(projectID)
	return args.Error(0)
}
