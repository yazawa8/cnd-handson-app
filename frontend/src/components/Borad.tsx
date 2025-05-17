import {
  DndContext,
  type DragEndEvent,
  MouseSensor,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import type React from "react";
import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { v4 as uuidv4 } from "uuid";
import { removeColumn } from "../features/columns/slice";
import type { Column as ColumnType } from "../features/columns/types";
import { updateTaskColumn } from "../features/tasks/slice";
import type { RootState } from "../store";
import AddButton from "./AddButton";
import Column from "./Column";

const KanbanBoard: React.FC = () => {
  const dispatch = useDispatch();
  const tasks = useSelector((state: RootState) => state.tasks.tasks);

  const initialColumns: ColumnType[] = useSelector(
    (state: RootState) => state.columns.columns,
  );

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
      name: "",
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
    setColumns([...columns, newColumn]);
  };

  const handleUpdateColumnName = (id: string, name: string) => {
    setColumns(
      columns.map((column) =>
        column.id === id ? { ...column, name } : column,
      ),
    );
    if (handleUpdateColumnName) {
      handleUpdateColumnName(id, name);
    }
  };

  const handleDeleteColumn = (id: string) => {
    dispatch(removeColumn(id));
    setColumns(columns.filter((c) => c.id !== id));
  };

  const sensors = useSensors(
    useSensor(MouseSensor, { activationConstraint: { distance: 5 } }),
  );

  return (
    <DndContext sensors={sensors} onDragEnd={handleDragEnd}>
      <div
        style={{ display: "flex", justifyContent: "flex-end", padding: "8px" }}
      >
        <AddButton label="列を追加" onClick={handleAddColumn} />
      </div>
      <div style={{ display: "flex", gap: "16px", padding: "16px" }}>
        {columns.map((column) => {
          const tasksInColumn = tasks.filter(
            (task) => task.columnId === column.id,
          );
          return (
            <Column
              key={column.id}
              column={column}
              tasks={tasksInColumn}
              onUpdateColumnName={handleUpdateColumnName}
              onDeleteColumn={handleDeleteColumn}
              initiallyEditing={column.name === ""}
            />
          );
        })}
      </div>
    </DndContext>
  );
};

export default KanbanBoard;
