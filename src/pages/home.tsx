import React from 'react';
import '../styles/home/home.css';

const Home: React.FC = () => {
  return (
    <div className="home">
      <div className="home__container">
        <div className="home__container__title">
          <div className="home__container__title__left">
            <h3 className="greeter">Good morning, Asep!</h3>
            <p className="walletAcc">Account: 123123123</p>
          </div>
          <div className="home__container__title__right">
            <p className="balance-title">Balance:</p>
            <h3 className="balance">IDR 13.750.000,00</h3>
          </div>
        </div>
        <div className="home__container__table">Ini Table</div>
      </div>
    </div>
  );
};

export default Home;
