import React, { useContext } from 'react';
import { NavLink, Outlet } from 'react-router-dom';
import '../styles/navbar/navbar.css';
import { AuthContext } from '../context/authContext';

const Navbar: React.FC = () => {
  const { authenticated } = useContext(AuthContext);

  return (
    <div className="app">
      <nav className="navbar">
        <h3 className="navbar__title">DigiWallet</h3>
        <ul className="navbar__list">
          <li className="navbar__list__item-first">Home</li>
          <li className="navbar__list__item">
            <NavLink
              to="/login"
              className={({ isActive }) => (isActive ? 'active' : 'nav-link')}
            >
              Login
            </NavLink>
          </li>
          <li className="navbar__list__item">
            <NavLink to="/register" className={'nav-link'}>
              Register
            </NavLink>
          </li>
          {authenticated && (
            <>
              <li className="navbar__list__item" id="main">
                <NavLink to="/">Home</NavLink>
              </li>
              <li className="navbar__list__item" id="main">
                Transfer
              </li>
              <li className="navbar__list__item" id="main">
                Topup
              </li>
              <li className="navbar__list__item" id="main">
                Games
              </li>
              <li className="navbar__list__item" id="main">
                Logout
              </li>
            </>
          )}
        </ul>
      </nav>
      <Outlet />
    </div>
  );
};

export default Navbar;
