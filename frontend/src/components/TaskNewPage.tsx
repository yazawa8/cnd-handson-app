import type React from "react";
import TaskForm from "./TaskForm";
import { useSearchParams } from "react-router-dom";
import type { Task } from "../features/tasks/types";

const NewTaskPage: React.FC = () => {
  const [searchParams] = useSearchParams();
  const columnIdFromParams = searchParams.get("columnId") || "";

  const columnOptions = [
    { id: "col-1", name: "Todo" },
    { id: "col-2", name: "In Progress" },
    { id: "col-3", name: "Done" },
  ];

  const handleSubmit = (task: Task) => {
    console.log("New Task Submitted:", task);
    // ここで保存処理やリダイレクトを行う
  };

  return (
    <TaskForm
      initialTask={{ columnId: columnIdFromParams }}
      columnOptions={columnOptions}
      onSubmit={handleSubmit}
    />
  );
};

export default NewTaskPage;
