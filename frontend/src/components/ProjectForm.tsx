import type React from "react";
import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useSelector, useDispatch } from "react-redux";
import type { RootState } from "../store";
import { addProject, updateProject } from "../features/projects/slice";
import type { Project } from "../features/projects/types";
import { Box, TextField, Button, Paper, Typography } from "@mui/material";
import { v4 as uuidv4 } from "uuid";

const ProjectForm: React.FC = () => {
  const { id } = useParams<{ id?: string }>();
  const isEdit = Boolean(id);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const project = useSelector((state: RootState) =>
    isEdit ? state.projects.projects.find((p) => p.id === id) : undefined,
  );

  const [name, setName] = useState("");
  const [description, setDescription] = useState("");

  useEffect(() => {
    if (isEdit) {
      if (project) {
        setName(project.name);
        setDescription(project.description || "");
      } else {
        navigate("/projects", { replace: true });
      }
    }
  }, [isEdit, project, navigate]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;
    const timestamp = new Date().toISOString();
    if (isEdit && project) {
      dispatch(
        updateProject({
          id: project.id,
          name,
          description,
          updatedAt: timestamp,
        }),
      );
    } else {
      const newProject: Project = {
        id: uuidv4(),
        name,
        description,
        createdAt: timestamp,
        updatedAt: timestamp,
      };
      dispatch(addProject(newProject));
    }
    navigate("/projects");
  };

  return (
    <Box sx={{ display: "flex", justifyContent: "center", mt: 4 }}>
      <Paper sx={{ p: 4, width: "100%", maxWidth: 600 }}>
        <Typography variant="h5" gutterBottom>
          {isEdit ? "プロジェクトを編集" : "プロジェクトを追加"}
        </Typography>
        <Box
          component="form"
          onSubmit={handleSubmit}
          sx={{ display: "flex", flexDirection: "column", gap: 2 }}
        >
          <TextField
            label="プロジェクト名"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <TextField
            label="説明"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            multiline
            rows={4}
          />
          <Box sx={{ display: "flex", justifyContent: "flex-end", gap: 2 }}>
            <Button onClick={() => navigate("/projects")}>キャンセル</Button>
            <Button type="submit" variant="contained" color="primary">
              {isEdit ? "保存" : "追加"}
            </Button>
          </Box>
        </Box>
      </Paper>
    </Box>
  );
};

export default ProjectForm;
