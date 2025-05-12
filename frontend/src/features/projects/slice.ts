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
      state.projects = action.payload;
    },
    setSelectedProject(state, action: PayloadAction<string>) {
      state.selectedId = action.payload;
    },
    deleteProject(state, action: PayloadAction<string>) {
      state.projects = state.projects.filter((project) => project.id !== action.payload);
      if (state.selectedId === action.payload) {
        state.selectedId = '';
      }
    },
    addProject(state, action: PayloadAction<Project>) {
      state.projects.push(action.payload);
    },
    updateProject(state, action: PayloadAction<Project>) {
      const index = state.projects.findIndex((project) => project.id === action.payload.id);
      if (index !== -1) {
        state.projects[index] = action.payload;
      }
    }
  },
});

export const { setProjects, setSelectedProject, deleteProject, addProject, updateProject } = projectsSlice.actions;
export default projectsSlice.reducer;
