import React from 'react';
import './App.css';
// import Register from './pages/register';
// import AuthProvider from './context/authContext';
// import { BrowserRouter, Route, Routes } from 'react-router-dom';
// import Login from './pages/login';
// import Navbar from './components/navbar';
// import { Navigate, Outlet } from 'react-router-dom';
// import { AuthContext } from './context/authContext';
// import Home from './pages/home';
// import Topup from './pages/topup';
// import Transfer from './pages/transfer';
// import Games from './pages/games';
// import Unauthenticated from './pages/unauthenticated';
import SuccessCard from './components/successCard';

const App: React.FC = () => {
  // const PrivateRoutes = () => {
  //   const { authenticated } = useContext(AuthContext);

  //   if (!authenticated) {
  //     return <Navigate to={'/unauthenticated'} replace />;
  //   }
  //   return <Outlet />;
  // };

  return (
    // <AuthProvider>
    //   <BrowserRouter>
    //     <Routes>
    //       <Route element={<Navbar />}>
    //         <Route path="/register" element={<Register />} />
    //         <Route path="/login" element={<Login />} />
    //         <Route element={<PrivateRoutes />}>
    //           <Route path="/" element={<Home />} />
    //           <Route path="/topup" element={<Topup />} />
    //           <Route path="/transfer" element={<Transfer />} />
    //           <Route path="/games" element={<Games />} />
    //         </Route>
    //       </Route>
    //       <Route path="*" element={<h1>404</h1>}></Route>
    //       <Route path="/unauthenticated" element={<Unauthenticated />}></Route>
    //     </Routes>
    //   </BrowserRouter>
    // </AuthProvider>
    <SuccessCard />
  );
};

export default App;
