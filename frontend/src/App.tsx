import React from 'react';
import { Provider, useSelector } from 'react-redux';
import { store } from './store';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import BoardList from './components/BoardList';
import Borad from './components/Borad';
import Login from './components/Login';
import Logout from './components/Logout';
import PrivateLayout from './components/PrivateLayout';
import ProjectList from './components/ProjectList';
import ProjectForm from './components/ProjectForm';
import { Project } from './features/projects/types';
import BoardForm from './components/BoardForm';
import NewTaskPage from './components/TaskNewPage';
import TaskEditPage from './components/TaskEditPage';



function App() {
  return (
    <Provider store={store}>
      <Router>
        <Routes> 
          <Route path="/login" element={<Login />} />
          <Route path="/logout" element={<Logout />} />
          <Route element={<PrivateLayout />}>
            <Route path="/boards" element={ <BoardList />} />
            <Route path="/" element={<ProjectList />} />
            <Route path="/projects/new" element={<ProjectForm open={false} project={null} onClose={function (): void {
              throw new Error('Function not implemented.');
            } } onSave={function (updated: Project): void {
              throw new Error('Function not implemented.');
            } } />} />
            <Route path="/projects/edit/:id" element={<ProjectForm open={false} project={null} onClose={function (): void {
              throw new Error('Function not implemented.');
            } } onSave={function (): void {
              throw new Error('Function not implemented.');
            } } />} />
            <Route path="/boards/:id" element={<Borad />} />
            <Route path="/boards/add" element={<BoardForm />} />
            <Route path="/boards/edit/:id" element={<BoardForm />} />
            <Route path="/tasks/new" element={<NewTaskPage open={false} task={null} onClose={function (): void {
              throw new Error('Function not implemented.');
            } } />} />
            <Route path="/tasks/edit/:id" element={<TaskEditPage open={false} task={null} onClose={function (): void {
              throw new Error('Function not implemented.');
            } } />} />
          </Route>
          <Route path="*" element={<Navigate to="/login" replace />} />
        </Routes>
      </Router>
    </Provider>
  );
}

export default App;
