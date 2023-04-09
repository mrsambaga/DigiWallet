import React from 'react';
import './App.css';
import Register from './pages/register';
import AuthProvider from './context/authContext';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Login from './pages/login';
import Navbar from './components/navbar';

const App: React.FC = () => {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route element={<Navbar />}>
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
          </Route>
          <Route path="*" element={<h1>404</h1>}></Route>
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  );
};

export default App;
