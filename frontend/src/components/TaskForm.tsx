import React, { useState } from 'react';
import { Box, TextField, Button, Paper, Typography, MenuItem, Select, InputLabel, FormControl } from '@mui/material';
import { Task } from '../features/tasks/types';

interface TaskFormProps {
  initialTask?: Partial<Task>;
  columnOptions: { id: string; name: string }[];
  onSubmit: (task: Task) => void;
  onCancel?: () => void;
}

const TaskForm: React.FC<TaskFormProps> = ({ initialTask = {}, columnOptions, onSubmit, onCancel }) => {
  const [title, setTitle] = useState(initialTask.title || '');
  const [description, setDescription] = useState(initialTask.description || '');
  const [columnId, setColumnId] = useState(initialTask.columnId || (columnOptions[0]?.id || ''));

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;

    onSubmit({
        id: initialTask.id || '',
        title,
        description,
        columnId,
        createdAt: initialTask.createdAt || new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        assigneeId: initialTask.assigneeId || '',
        status: ''
    });
  };

  return (
    <Paper sx={{ p: 4, width: '100%', maxWidth: 600, margin: 'auto' }}>
      <Typography variant="h5" gutterBottom>
        {initialTask.id ? 'タスクを編集' : '新規タスク作成'}
      </Typography>
      <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <TextField
          label="タイトル"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />

        <TextField
          label="説明"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          multiline
          rows={4}
        />

        <FormControl fullWidth>
          <InputLabel id="column-select-label">カラム</InputLabel>
          <Select
            labelId="column-select-label"
            value={columnId}
            onChange={(e) => setColumnId(e.target.value)}
            label="カラム"
          >
            {columnOptions.map((option) => (
              <MenuItem key={option.id} value={option.id}>
                {option.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        <Box sx={{ display: 'flex', justifyContent: 'flex-end', gap: 2 }}>
          {onCancel && <Button onClick={onCancel}>キャンセル</Button>}
          <Button type="submit" variant="contained" color="primary">
            保存
          </Button>
        </Box>
      </Box>
    </Paper>
  );
};

export default TaskForm;
