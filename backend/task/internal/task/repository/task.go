package repository

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/task/internal/task/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/task/pkg/db"
	"github.com/google/uuid"
)

func GetTaskByID(task_id uuid.UUID) (*model.Task, error) {
	var task model.Task
	if err := db.DB.Where("id = ?", task_id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func List(column_id uuid.UUID, assignee_id uuid.UUID, page int32, page_size int32) ([]*model.Task, int32, error) {
	var tasks []*model.Task
	var totalCount int64

	query := db.DB.Model(&model.Task{})
	if column_id != uuid.Nil {
		query = query.Where("column_id = ?", column_id)
	}
	if assignee_id != uuid.Nil {
		query = query.Where("assignee_id = ?", assignee_id)
	}
	// カウントを取得
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if page > 0 && page_size > 0 {
		query = query.Offset(int((page - 1) * page_size)).Limit(int(page_size))
	}

	query = query.Order("start_time DESC")

	if err := query.Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, int32(totalCount), nil
}

func Create(task *model.Task) error {
	if err := db.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func Update(task *model.Task) error {
	if err := db.DB.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func Delete(task_id uuid.UUID) error {
	if err := db.DB.Delete(&model.Task{}, task_id).Error; err != nil {
		return err
	}
	return nil
}
