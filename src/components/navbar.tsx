import React from 'react';
import '../styles/navbar/navbar.css';

type NavbarProps = {
  type: string;
};

const Navbar: React.FC<NavbarProps> = ({ type }) => {
  return (
    <div className="navbar">
      <h3 className="navbar__title">DigiWallet</h3>
      <ul className="navbar__list">
        {type == 'login' ? (
          <>
            <li className="navbar__list__item-first">Home</li>
            <li className="navbar__list__item">Login</li>
            <li className="navbar__list__item">Register</li>
          </>
        ) : type == 'main' ? (
          <>
            <li className="navbar__list__item" id="main">
              Home
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
        ) : null}
      </ul>
    </div>
  );
};

export default Navbar;
