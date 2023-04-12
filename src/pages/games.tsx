import React, { useEffect, useState } from 'react';
import Title from '../components/title';
import useFetchGet from '../hooks/useFetchGet';
import { notifyError } from '../components/notification';
import { ProfileResponse } from '../types/types';
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

  const [selectedId, setSelectedId] = useState(0);
  const [boxClass, setBoxClass] = useState('box-inactive');
  const [chance] = useState(5);
  const boxOnClickHandler = (id: number) => {
    if (chance > 0 && boxClass == 'box-inactive') {
      setBoxClass('box-active');
      setSelectedId(id);
      return;
    } else if (chance == 0 && boxClass == 'box-inactive') {
      setBoxClass('box-no-chance');
      return;
    }

    setBoxClass('box-inactive');
  };

  const boxContent: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9];

  return (
    <div className="games">
      <div className="games__container">
        <Title idName="games-title" contentProps={profileResponse!} />
        <div className="games__container__table">
          <h1>Games</h1>
          <p>Choose random box below to get reward!</p>
          <p>Chance : {chance}</p>
          <NavLink to="/leaderboard">Check Leaderboard</NavLink>
          <div className="games__container__table__box">
            <button
              onClick={() => boxOnClickHandler(1)}
              className={boxClass}
              id={
                selectedId == 1 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[0] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(2)}
              className={boxClass}
              id={
                selectedId == 2 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[1] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(3)}
              className={boxClass}
              id={
                selectedId == 3 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[2] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(4)}
              className={boxClass}
              id={
                selectedId == 4 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[3] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(5)}
              className={boxClass}
              id={
                selectedId == 5 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[4] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(6)}
              className={boxClass}
              id={
                selectedId == 6 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[5] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(7)}
              className={boxClass}
              id={
                selectedId == 7 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[6] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(8)}
              className={boxClass}
              id={
                selectedId == 8 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[7] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(9)}
              className={boxClass}
              id={
                selectedId == 9 && boxClass == 'box-active' ? 'selectedBox' : ''
              }
            >
              {boxClass == 'box-active' ? boxContent[8] : ''}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Games;
