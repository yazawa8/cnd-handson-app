
import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../store';
import { Grid, Card, CardContent, Typography, CardActionArea, Box } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { setSelectedProject } from '../features/projects/slice';

const ProjectList: React.FC = () => {
  const projects = useSelector((state: RootState) => state.projects.projects);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleClick = (projectId: string) => {
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
              <CardActionArea onClick={() => handleClick(proj.id)}>
                <CardContent>
                  <Typography variant="h5">{proj.name}</Typography>
                  {proj.description && (
                    <Typography variant="body2" color="textSecondary">
                      {proj.description}
                    </Typography>
                  )}
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
