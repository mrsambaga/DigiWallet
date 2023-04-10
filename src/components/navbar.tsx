import React, { useContext } from 'react';
import { NavLink, Outlet, useNavigate } from 'react-router-dom';
import '../styles/navbar/navbar.css';
import { AuthContext } from '../context/authContext';
import { SetCookie } from '../function/cookies';

const Navbar: React.FC = () => {
  const { authenticated, setAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleClickLogout = () => {
    setAuthenticated(false);
    SetCookie('token', '', 0);
    navigate(`/login`);
  };

  return (
    <div className="app">
      <nav className="navbar">
        <h3 className="navbar__title">DigiWallet</h3>
        <ul className="navbar__list">
          {!authenticated && (
            <>
              <li className="navbar__list__item-first">Home</li>
              <li className="navbar__list__item">
                <NavLink
                  to="/login"
                  className={({ isActive }) =>
                    isActive ? 'active' : 'nav-link'
                  }
                >
                  Login
                </NavLink>
              </li>
              <li className="navbar__list__item">
                <NavLink
                  to="/register"
                  className={({ isActive }) =>
                    isActive ? 'active' : 'nav-link'
                  }
                >
                  Register
                </NavLink>
              </li>
            </>
          )}
          {authenticated && (
            <>
              <li className="navbar__list__item">
                <NavLink
                  to="/"
                  className={({ isActive }) =>
                    isActive ? 'active-main' : 'nav-link-main'
                  }
                >
                  Home
                </NavLink>
              </li>
              <li className="navbar__list__item">
                <NavLink
                  to="/transfer"
                  className={({ isActive }) =>
                    isActive ? 'active-main' : 'nav-link-main'
                  }
                >
                  Transfer
                </NavLink>
              </li>
              <li className="navbar__list__item">
                <NavLink
                  to="/topup"
                  className={({ isActive }) =>
                    isActive ? 'active-main' : 'nav-link-main'
                  }
                >
                  Topup
                </NavLink>
              </li>
              <li className="navbar__list__item">
                <NavLink
                  to="/games"
                  className={({ isActive }) =>
                    isActive ? 'active-main' : 'nav-link-main'
                  }
                >
                  Games
                </NavLink>
              </li>
              <li className="navbar__list__item">
                <button className="logout" onClick={handleClickLogout}>
                  Logout
                </button>
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
