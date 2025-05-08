// src/components/ProjectList.tsx
import React, { useState, MouseEvent } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../store';
import { Grid, Card, CardContent, Typography, CardActionArea, Box, IconButton, Menu, MenuItem } from '@mui/material';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import { useNavigate } from 'react-router-dom';
import { setSelectedProject } from '../features/projects/slice';
import { deleteProject } from '../features/projects/slice';

const ProjectList: React.FC = () => {
  const projects = useSelector((state: RootState) => state.projects.projects);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [menuProjectId, setMenuProjectId] = useState<string>('');

  const openMenu = (e: MouseEvent<HTMLElement>, projectId: string) => {
    e.stopPropagation(); // カードクリックと競合しないように
    setAnchorEl(e.currentTarget);
    setMenuProjectId(projectId);
  };

  const closeMenu = () => {
    setAnchorEl(null);
    setMenuProjectId('');
  };

  const handleEdit = () => {
    // 編集時にはモーダルなど開く or 画面遷移
    closeMenu();
  };

  const handleDelete = () => {
    if (window.confirm('本当に削除しますか？')) {
        dispatch(deleteProject(menuProjectId));
    }
    closeMenu();
  };

  const handleClickCard = (projectId: string) => {
    dispatch(setSelectedProject(projectId));
    navigate('/boards');
  };

  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h4" gutterBottom>
        Projects
      </Typography>
      <Grid container spacing={2}>
        {projects.map((proj) => (
          <Grid item xs={12} sm={6} md={4} key={proj.id}>
            <Card>
              <CardActionArea onClick={() => handleClickCard(proj.id)}>
                <CardContent sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <Box>
                    <Typography variant="h5">{proj.name}</Typography>
                    {proj.description && (
                      <Typography variant="body2" color="textSecondary">
                        {proj.description}
                      </Typography>
                    )}
                  </Box>
                  <IconButton
                    edge="end"
                    aria-label="settings"
                    onClick={(e) => openMenu(e, proj.id)}
                  >
                    <MoreVertIcon />
                  </IconButton>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={closeMenu}
      >
        <MenuItem onClick={handleEdit}>編集</MenuItem>
        <MenuItem onClick={handleDelete}>削除</MenuItem>
      </Menu>
    </Box>
  );
};

export default ProjectList;
