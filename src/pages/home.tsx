import React, { useEffect } from 'react';
import '../styles/home/home.css';
import useFetchGet from '../hooks/useFetchGet';
import { decodeToken } from 'react-jwt';
import { notifyError } from '../components/notification';

type Claims = {
  id: number;
  email: string;
  exp: number;
  iat: number;
  iss: string;
};

const Home: React.FC = () => {
  const token = localStorage.getItem('token');
  const claims: Claims | null = token ? decodeToken(token!) : null;
  const userId = claims?.id;
  const { out, error } = useFetchGet(
    `http://localhost:8000/users/${userId}`,
    token!,
  );

  useEffect(() => {
    if (error) {
      notifyError(error.response?.data?.message || error.message);
      return;
    }

    localStorage.setItem('wallet_number', out?.data.wallet_number);
  }, [out, error]);

  return (
    <div className="home">
      <div className="home__container">
        <div className="home__container__title">
          <div className="home__container__title__left">
            <h3 className="greeter">Good morning, {out?.data.user_name}</h3>
            <p className="walletAcc">Account: {out?.data.wallet_number}</p>
          </div>
          <div className="home__container__title__right">
            <p className="balance-title">Balance:</p>
            <h3 className="balance">IDR {out?.data.balance},00</h3>
          </div>
        </div>
        <div className="home__container__table">Ini Table</div>
      </div>
    </div>
  );
};

export default Home;
