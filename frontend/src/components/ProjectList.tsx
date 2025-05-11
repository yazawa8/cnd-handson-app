// src/components/ProjectList.tsx
import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../store';
import { Grid, Card, CardContent, Typography, CardActionArea, Box, IconButton, Menu, MenuItem } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { setSelectedProject } from '../features/projects/slice';
import { deleteProject } from '../features/projects/slice';
import MoreMenu, { MoreMenuOption } from './MoreMenu';
import AddButton from './AddButton';
import ProjectForm from './ProjectForm';
import { Project } from '../features/projects/types';

const ProjectList: React.FC = () => {
  const projects = useSelector((state: RootState) => state.projects.projects);
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [editOpen, setEditOpen] = React.useState(false);
  const [editTarget, setEditTarget] = React.useState<Project | null>(null);

  const openEdit = (proj: Project) => {
    setEditTarget(proj);
    setEditOpen(true);
  };
  const closeEdit = () => {
    setEditOpen(false);
    setEditTarget(null);
  };
  const handleSave = (updated: Project) => {
    dispatch(updateProject(updated));
    closeEdit();
  };

  const options: MoreMenuOption<string>[] = [
      { label: '編集', onClick: (id) => openEdit(proj) },
      { label: '削除', onClick: (id) => {
        if (window.confirm('本当に削除しますか？')) {
          dispatch(deleteProject(id));
       }}},];

  const handleClickCard = (projectId: string) => {
    dispatch(setSelectedProject(projectId));
    navigate('/boards');
  };

  const onAdd = () => {
    navigate('/projects/new');
  }

  return (
    <Box sx={{ p: 2 }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', padding: '8px' }}>
        <Typography variant="h4" gutterBottom>
          Projects
        </Typography>
        <AddButton label="Projectを追加" onClick={onAdd} />
      </div>
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
                  <MoreMenu id={proj.id} options={options} />
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default ProjectList;
function updateProject(updated: Project): any {
  throw new Error('Function not implemented.');
}

