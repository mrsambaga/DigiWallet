import React, { useEffect, useState } from 'react';
import Title from '../components/title';
import useFetchGet from '../hooks/useFetchGet';
import { notifyError } from '../components/notification';
import { ProfileResponse } from '../types/types';
import { decodeToken } from 'react-jwt';
import { Claims } from '../types/types';
import { GetCookie } from '../function/cookies';
import '../styles/games/games.css';
import { NavLink } from 'react-router-dom';

const Games: React.FC = () => {
  const [profileResponse, setProfileResponse] = useState<ProfileResponse>({
    Balance: 0,
    Email: '',
    UserId: 0,
    UserName: '',
    WalletNumber: 0,
  });
  const token = GetCookie('token');
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
          <h1>Games</h1>
          <p>Choose random box below to get reward!</p>
          <p>Chance : 3</p>
          <NavLink to="/leaderboard">Check Leaderboard</NavLink>
          <div className="games__container__table__box">
            <div className="games__container__table__box__row">
              <button>1</button>
              <button>2</button>
              <button>3</button>
            </div>
            <div className="games__container__table__box__row">
              <button>4</button>
              <button>5</button>
              <button>6</button>
            </div>
            <div className="games__container__table__box__row">
              <button>7</button>
              <button>8</button>
              <button>9</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Games;
