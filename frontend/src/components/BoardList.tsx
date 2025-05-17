import type React from "react";
import { useSelector, useDispatch } from "react-redux";
import type { RootState } from "../store";
import type { Board } from "../features/boards/types";
import {
  Box,
  Grid,
  Card,
  CardContent,
  Typography,
  CardActionArea,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import MoreMenu, { type MoreMenuOption } from "./MoreMenu";
import AddButton from "./AddButton";

const BoardList: React.FC = () => {
  const navigate = useNavigate();
  const boards = useSelector((state: RootState) => state.boards.boards);
  const dispatch = useDispatch();
  const handleBoardClick = (board: Board) => {
    navigate(`/boards/${board.id}`);
  };

  const options: MoreMenuOption<string>[] = [
    { label: "編集", onClick: (id) => navigate(`/boards/edit/${id}`) },
    { label: "削除", onClick: (id) => navigate(`/boards/delete/${id}`) },
  ];

  const onAdd = () => {
    navigate("/boards/new");
  };
  return (
    <div style={{ padding: "16px" }}>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          padding: "8px",
        }}
      >
        <Typography variant="h4" gutterBottom>
          Boards
        </Typography>
        <AddButton label="Boardを追加" onClick={onAdd} />
      </div>
      <Grid container spacing={2}>
        {boards.map((board: Board) => (
          <Grid item xs={12} sm={6} md={4} key={board.id}>
            <Card
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "space-between",
              }}
            >
              <CardActionArea onClick={() => handleBoardClick(board)}>
                <CardContent>
                  <Box>
                    <Typography variant="h5">{board.name}</Typography>
                    <Typography variant="body2" color="textSecondary">
                      {board.description || "No description available."}
                    </Typography>
                  </Box>
                </CardContent>
              </CardActionArea>
              <MoreMenu id={board.id} options={options} />
            </Card>
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export default BoardList;
