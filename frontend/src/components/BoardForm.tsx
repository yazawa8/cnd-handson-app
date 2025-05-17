import { Box, Button, Paper, TextField, Typography } from "@mui/material";
import type React from "react";
import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate, useParams } from "react-router-dom";
import { addBoard, updateBoard } from "../features/boards/slice";
import type { RootState } from "../store";

const BoardForm: React.FC = () => {
  const { id } = useParams<{ id?: string }>();
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const board = useSelector((state: RootState) =>
    state.boards.boards.find((b) => b.id === id),
  );

  const [name, setName] = useState<string>(board?.name || "");
  const [description, setDescription] = useState<string>(
    board?.description || "",
  );

  useEffect(() => {
    if (id && !board) {
      navigate("/boards", { replace: true });
    }
  }, [board, id, navigate]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;

    if (id) {
      dispatch(
        updateBoard({
          id,
          name,
          description,
          updatedAt: new Date().toISOString(),
          createdAt: "",
        }),
      );
    } else {
      dispatch(
        addBoard({
          name,
          description,
          id: "",
          createdAt: "",
          updatedAt: "",
        }),
      );
    }
    navigate("/boards");
  };

  return (
    <Box sx={{ display: "flex", justifyContent: "center", mt: 4 }}>
      <Paper sx={{ p: 4, width: "100%", maxWidth: 600 }}>
        <Typography variant="h5" gutterBottom>
          {id ? "ボードを編集" : "ボードを追加"}
        </Typography>
        <Box
          component="form"
          onSubmit={handleSubmit}
          sx={{ display: "flex", flexDirection: "column", gap: 2 }}
        >
          <TextField
            label="ボード名"
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
            <Button onClick={() => navigate("/boards")}>キャンセル</Button>
            <Button type="submit" variant="contained" color="primary">
              保存
            </Button>
          </Box>
        </Box>
      </Paper>
    </Box>
  );
};

export default BoardForm;
