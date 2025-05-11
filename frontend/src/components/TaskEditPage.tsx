import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../store';
import TaskForm from './TaskForm';

const TaskEditPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const task = useSelector((state: RootState) =>
    state.tasks.tasks.find((t) => t.id === id)
  );

  const columnOptions = [
    { id: 'col-1', name: 'Todo' },
    { id: 'col-2', name: 'In Progress' },
    { id: 'col-3', name: 'Done' },
  ];

  if (!task) {
    return <div>タスクが見つかりません</div>;
  }

  const handleSubmit = (updatedTask: any) => {
    console.log('Updated Task:', updatedTask);
    // 保存処理やステート更新、リダイレクト処理をここに追加
    navigate('/');
  };

  return (
    <TaskForm
      initialTask={task}
      columnOptions={columnOptions}
      onSubmit={handleSubmit}
      onCancel={() => navigate(-1)}
    />
  );
};

export default TaskEditPage;
