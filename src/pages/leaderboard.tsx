import React, { useEffect, useState } from 'react';
import Title from '../components/title';
import useFetchGet from '../hooks/useFetchGet';
import { notifyError } from '../components/notification';
import { ProfileResponse } from '../types/types';
import { GetCookie } from '../function/cookies';
import '../styles/games/games.css';
import { NavLink } from 'react-router-dom';

const Leaderboard: React.FC = () => {
  const [profileResponse, setProfileResponse] = useState<ProfileResponse>({
    Balance: 0,
    Email: '',
    UserId: 0,
    UserName: '',
    WalletNumber: 0,
  });
  const token = GetCookie('token');
  const { out, error } = useFetchGet(`http://localhost:8000/profile`, token!);

  useEffect(() => {
    if (error) {
      notifyError(error.response?.data?.message || error.message);
      return;
    }

    if (out != null) {
      const profileResponse: ProfileResponse = {
        Balance: out.data.balance,
        Email: out.data.email,
        UserId: out.data.user_id,
        UserName: out.data.user_name,
        WalletNumber: out.data.wallet_number,
      };

      localStorage.setItem('wallet_number', out.data.wallet_number);
      setProfileResponse(profileResponse);
    }
  }, [out, error]);

  return (
    <div className="games">
      <div className="games__container">
        <Title idName="games-title" contentProps={profileResponse!} />
        <div className="games__container__table">
          <h1>Leaderboard</h1>
          <NavLink to="/games">Back to Games</NavLink>
          <div className="games__container__table__box"></div>
        </div>
      </div>
    </div>
  );
};

export default Leaderboard;
