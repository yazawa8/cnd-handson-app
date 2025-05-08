
import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../store';
import { Board } from '../features/board/types';
import { Box, Grid, Card, CardContent, Typography, CardActionArea } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import MoreMenu, { MoreMenuOption } from './MoreMenu';

const BoardList: React.FC = () => {
  const navigate = useNavigate();
  const boards = useSelector((state: RootState) => state.boards.boards);
  const dispatch = useDispatch();
  const handleBoardClick = (board: Board) => {
    navigate(`/boards/${board.id}`);
  };

  const options: MoreMenuOption<string>[] = [
    { label: '編集', onClick: (id) => navigate(`/boards/edit/${id}`) },
    { label: '削除', onClick: (id) => navigate('/') },
  ];

  return (
    <div style={{ padding: '16px' }}>
      <Typography variant="h4" gutterBottom>
        Boards
      </Typography>
      <Grid container spacing={2}>
        {boards.map((board: Board) => (
          <Grid item xs={12} sm={6} md={4} key={board.id}>
            <Card>
              <CardActionArea onClick={() => handleBoardClick(board)}>
                <CardContent sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <Box>
                    <Typography variant="h5">{board.name}</Typography>
                    <Typography variant="body2" color="textSecondary">
                      {board.description || 'No description available.'}
                    </Typography>
                  </Box>
                  <MoreMenu id={board.id} options={options} />
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export default BoardList;
