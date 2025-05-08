import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Project } from './types';

interface ProjectsState {
  projects: Project[];
  selectedId: string | '';
}

const initialState: ProjectsState = {
  projects: [
    {
      id: 'proj-1',
      name: 'Project Alpha',
      description: 'Description for Project Alpha',
    },
    {
      id: 'proj-2',
      name: 'Project Beta',
      description: 'Description for Project Beta',
    },
    {
      id: 'proj-3',
      name: 'Project Gamma',
      description: 'Description for Project Gamma',
    },
  ],
  selectedId: '',
};

const projectsSlice = createSlice({
  name: 'projects',
  initialState,
  reducers: {
    setProjects(state, action: PayloadAction<Project[]>) {
      state.list = action.payload;
    },
    setSelectedProject(state, action: PayloadAction<string>) {
      state.selectedId = action.payload;
    },
  },
});

export const { setProjects, setSelectedProject } = projectsSlice.actions;
export default projectsSlice.reducer;
