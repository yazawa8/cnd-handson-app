// src/components/Column.tsx
import React from 'react';
import { Typography, Paper, Box } from '@mui/material';
import { Column as ColumnType } from '../features/columns/types';
import { useDroppable } from '@dnd-kit/core';
import { Task } from '../features/tasks/types';
import TaskCard from './TaskCrad';
import MoreMenu, { MoreMenuOption } from './MoreMenu';
import { useNavigate } from 'react-router-dom';

interface ColumnProps {
  column: ColumnType;
  tasks: Task[];
}

const Column: React.FC<ColumnProps> = ({ column, tasks }) => {
  const { isOver, setNodeRef } = useDroppable({
    id: column.id,
    data: { column },
  });
  const navigate = useNavigate();

    const options: MoreMenuOption<string>[] = [
      { label: '編集', onClick: (id) => navigate(`/boards/edit/${id}`) },
      { label: '削除', onClick: (id) => navigate('/') },
    ];

  return (
    <Paper
      ref={setNodeRef}
      sx={{
        width: '300px',
        padding: 2,}}
    > <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <Typography variant="h5" gutterBottom>
          {column.name}
        </Typography>
        <MoreMenu id={column.id} options={options} />
      </Box>
      {tasks.map((task) => (
        <TaskCard key={task.id} task={task} />
      ))}
    </Paper>
  );
};

export default Column;
