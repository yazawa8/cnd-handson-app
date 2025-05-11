import React, { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../store';
import Column from './Column';
import { Column as ColumnType } from '../features/columns/types';
import { removeColumn } from '../features/columns/slice';
import { updateTaskColumn } from '../features/tasks/slice';
import { DndContext, DragEndEvent } from '@dnd-kit/core';
import AddButton from './AddButton';
import { v4 as uuidv4 } from 'uuid';

const KanbanBoard: React.FC = () => {
  const dispatch = useDispatch();
  const tasks = useSelector((state: RootState) => state.tasks.tasks);

  const initialColumns: ColumnType[] = useSelector((state: RootState) => state.columns.columns);

  // 初期は空のカラムリスト
  const [columns, setColumns] = useState<ColumnType[]>(initialColumns);

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;
    if (!over) return;
    const taskId = active.id.toString();
    const newColumnId = over.id.toString();
    dispatch(updateTaskColumn({ taskId, columnId: newColumnId }));
  };

  const handleAddColumn = () => {
    const newColumn = {
      id: uuidv4(),
      name: '',
    };
    setColumns([...columns, newColumn]);
  };

  const handleUpdateColumnName = (id: string, name: string) => {
    setColumns(columns.map(column => column.id === id ? { ...column, name } : column));
    if (handleUpdateColumnName) {
      handleUpdateColumnName(id, name);
    }
  };

  const handleDeleteColumn = (id: string) => {
    dispatch(removeColumn(id));
    setColumns(columns.filter(c => c.id !== id));
  };

  return (
    <DndContext onDragEnd={handleDragEnd}>
      <div style={{ display: 'flex', justifyContent: 'flex-end', padding: '8px' }}>
        <AddButton label="列を追加" onClick={handleAddColumn} />
      </div>
      <div style={{ display: 'flex', gap: '16px', padding: '16px' }}>
        {columns.map((column) => {
          const tasksInColumn = tasks.filter(task => task.columnId === column.id);
          return (
            <Column
              key={column.id}
              column={column}
              tasks={tasksInColumn}
              onUpdateColumnName={handleUpdateColumnName}
              onDeleteColumn={handleDeleteColumn}
              initiallyEditing={column.name === ''}
            />
          );
        })}
      </div>
    </DndContext>
  );
};

export default KanbanBoard;
