package test

import (
	"os"
	"testing"
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/repository/test"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	// テスト実行中はDB接続をスキップする
	if err := os.Setenv("DB_SKIP_CONNECT", "true"); err != nil {
		// テストの初期化中なので、panicを使用するのが適切
		panic("Failed to set environment variable: " + err.Error())
	}
}

// テスト用プロジェクトデータの作成
func createTestProject() *model.Project {
	return &model.Project{
		ID:          uuid.New(),
		Name:        "Test Project",
		Description: "This is a test project",
		OwnerID:     uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// CreateProjectのシンプルなテスト
func TestCreateProject(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := new(test.MockProjectRepository)
	projectService := service.NewProjectService(mockRepo)

	// テストデータ
	name := "Test Project"
	description := "This is a test project"
	ownerID := uuid.New()

	// モックの設定
	mockRepo.On("CreateProject", mock.AnythingOfType("*model.Project")).Return(
		&model.Project{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
			OwnerID:     ownerID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, nil)

	// テスト実行
	result, err := projectService.CreateProject(name, description, ownerID)

	// 検証
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, name, result.Name)
	assert.Equal(t, description, result.Description)
	assert.Equal(t, ownerID, result.OwnerID)

	// モック呼び出しの検証
	mockRepo.AssertExpectations(t)
}

// GetProjectのシンプルなテスト
func TestGetProject(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := new(test.MockProjectRepository)
	projectService := service.NewProjectService(mockRepo)

	// テストデータ
	testProject := createTestProject()

	// モックの設定
	mockRepo.On("GetProjectByID", testProject.ID).Return(testProject, nil)

	// テスト実行
	result, err := projectService.GetProject(testProject.ID)

	// 検証
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, testProject.ID, result.ID)

	// モック呼び出しの検証
	mockRepo.AssertExpectations(t)
}

// ListProjectsのシンプルなテスト
func TestListProjects(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := new(test.MockProjectRepository)
	projectService := service.NewProjectService(mockRepo)

	// テストデータ
	testProjects := []*model.Project{createTestProject(), createTestProject()}

	// モックの設定
	mockRepo.On("ListProjects").Return(testProjects, nil)

	// テスト実行
	results, err := projectService.ListProjects()

	// 検証
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, results, "Results should not be nil")
	assert.Len(t, results, 2, "Should return 2 projects")

	// モック呼び出しの検証
	mockRepo.AssertExpectations(t)
}

// ListProjectsByOwnerのシンプルなテスト
func TestListProjectsByOwner(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := new(test.MockProjectRepository)
	projectService := service.NewProjectService(mockRepo)

	// テストデータ
	ownerID := uuid.New()
	testProjects := []*model.Project{createTestProject(), createTestProject()}

	// モックの設定
	mockRepo.On("GetProjectsByOwnerID", ownerID).Return(testProjects, nil)

	// テスト実行
	results, err := projectService.ListProjectsByOwner(ownerID)

	// 検証
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, results, "Results should not be nil")
	assert.Len(t, results, 2, "Should return 2 projects")

	// モック呼び出しの検証
	mockRepo.AssertExpectations(t)
}

// UpdateProjectのシンプルなテスト
func TestUpdateProject(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := new(test.MockProjectRepository)
	projectService := service.NewProjectService(mockRepo)

	// テストデータ
	testProject := createTestProject()
	newName := "Updated Project Name"

	// モックの設定
	mockRepo.On("GetProjectByID", testProject.ID).Return(testProject, nil)
	mockRepo.On("UpdateProject", mock.AnythingOfType("*model.Project")).Return(
		&model.Project{
			ID:          testProject.ID,
			Name:        newName,
			Description: testProject.Description,
			OwnerID:     testProject.OwnerID,
			CreatedAt:   testProject.CreatedAt,
			UpdatedAt:   time.Now(),
		}, nil)

	// テスト実行
	result, err := projectService.UpdateProject(testProject.ID, &newName, nil)

	// 検証
	assert.NoError(t, err, "Should not return an error")
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, newName, result.Name, "Name should be updated")

	// モック呼び出しの検証
	mockRepo.AssertExpectations(t)
}

// DeleteProjectのシンプルなテスト
func TestDeleteProject(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := new(test.MockProjectRepository)
	projectService := service.NewProjectService(mockRepo)

	// テストデータ
	testProject := createTestProject()

	// モックの設定
	mockRepo.On("GetProjectByID", testProject.ID).Return(testProject, nil)
	mockRepo.On("DeleteProject", testProject.ID).Return(nil)

	// テスト実行
	err := projectService.DeleteProject(testProject.ID)

	// 検証
	assert.NoError(t, err, "Should not return an error")

	// モック呼び出しの検証
	mockRepo.AssertExpectations(t)
}
