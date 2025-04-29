package repository

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/pkg/db"
	"github.com/google/uuid"
)

// projectRepository はProjectRepositoryインターフェースの実装
type projectRepository struct{}

// GetProjectByID はIDによるプロジェクトの取得
func (r *projectRepository) GetProjectByID(projectID uuid.UUID) (*model.Project, error) {
	var project model.Project
	if err := db.DB.Where("id = ?", projectID).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

// GetProjectsByOwnerID はオーナーIDによるプロジェクト一覧の取得
func (r *projectRepository) GetProjectsByOwnerID(ownerID uuid.UUID) ([]*model.Project, error) {
	var projects []*model.Project
	if err := db.DB.Where("owner_id = ?", ownerID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// ListProjects は全プロジェクトの取得
func (r *projectRepository) ListProjects() ([]*model.Project, error) {
	var projects []*model.Project
	if err := db.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// CreateProject は新規プロジェクトの作成
func (r *projectRepository) CreateProject(project *model.Project) (*model.Project, error) {
	if err := db.DB.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

// UpdateProject はプロジェクト情報の更新
func (r *projectRepository) UpdateProject(project *model.Project) (*model.Project, error) {
	if err := db.DB.Model(&model.Project{}).Where("id = ?", project.ID).Updates(project).Error; err != nil {
		return nil, err
	}
	// 更新後のプロジェクト情報を取得
	updatedProject, err := r.GetProjectByID(project.ID)
	if err != nil {
		return nil, err
	}
	return updatedProject, nil
}

// DeleteProject はプロジェクトの削除
func (r *projectRepository) DeleteProject(projectID uuid.UUID) error {
	if err := db.DB.Where("id = ?", projectID).Delete(&model.Project{}).Error; err != nil {
		return err
	}
	return nil
}

// 以下は後方互換性のための関数群
// GetProjectByID はIDによるプロジェクトの取得（後方互換性用）
func GetProjectByID(projectID uuid.UUID) (*model.Project, error) {
	return DefaultProjectRepository.GetProjectByID(projectID)
}

// GetProjectsByOwnerID はオーナーIDによるプロジェクト一覧の取得（後方互換性用）
func GetProjectsByOwnerID(ownerID uuid.UUID) ([]*model.Project, error) {
	return DefaultProjectRepository.GetProjectsByOwnerID(ownerID)
}

// ListProjects は全プロジェクトの取得（後方互換性用）
func ListProjects() ([]*model.Project, error) {
	return DefaultProjectRepository.ListProjects()
}

// CreateProject は新規プロジェクトの作成（後方互換性用）
func CreateProject(project *model.Project) (*model.Project, error) {
	return DefaultProjectRepository.CreateProject(project)
}

// UpdateProject はプロジェクト情報の更新（後方互換性用）
func UpdateProject(project *model.Project) (*model.Project, error) {
	return DefaultProjectRepository.UpdateProject(project)
}

// DeleteProject はプロジェクトの削除（後方互換性用）
func DeleteProject(projectID uuid.UUID) error {
	return DefaultProjectRepository.DeleteProject(projectID)
}