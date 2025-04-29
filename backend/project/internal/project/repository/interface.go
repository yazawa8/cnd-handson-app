package repository

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/google/uuid"
)

// ProjectRepository はプロジェクトリポジトリのインターフェース
type ProjectRepository interface {
	GetProjectByID(projectID uuid.UUID) (*model.Project, error)
	GetProjectsByOwnerID(ownerID uuid.UUID) ([]*model.Project, error)
	ListProjects() ([]*model.Project, error)
	CreateProject(project *model.Project) (*model.Project, error)
	UpdateProject(project *model.Project) (*model.Project, error)
	DeleteProject(projectID uuid.UUID) error
}

// DefaultProjectRepository は標準のプロジェクトリポジトリインスタンス
var DefaultProjectRepository ProjectRepository

// 注: NewProjectRepository関数の実装はproject.goファイルに移動しました