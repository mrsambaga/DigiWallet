import React from 'react';
import '../styles/navbar/navbar.css';

const Navbar: React.FC = () => {
  return (
    <div className="navbar">
      <h3 className="navbar__title">DigiWallet</h3>
      <ul className="navbar__list">
        <li className="navbar__list__item-first">Home</li>
        <li className="navbar__list__item">Login</li>
        <li className="navbar__list__item">Register</li>
      </ul>
    </div>
  );
};

export default Navbar;
