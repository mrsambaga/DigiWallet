import React, { useEffect, useState } from 'react';
import Title from '../components/title';
import useFetchGet from '../hooks/useFetchGet';
import { notifyError } from '../components/notification';
import { LeaderboardResp, Profile, ProfileResponse } from '../types/types';
import { GetCookie } from '../helper/cookies';
import '../styles/games/games.css';
import { NavLink } from 'react-router-dom';

const Leaderboard: React.FC = () => {
  const [profileResponse, setProfileResponse] = useState<Profile>({
    Balance: 0,
    Email: '',
    UserId: 0,
    UserName: '',
    WalletNumber: '',
  });
  const token = GetCookie('token');
  const { out: outProfile, error: outError } = useFetchGet<{
    data: ProfileResponse;
  }>(`http://localhost:8000/profile`, token!);

  useEffect(() => {
    if (outError) {
      const errorMessage = outError.response?.data || outError.message;
      notifyError(JSON.stringify(errorMessage));
      return;
    }

    if (outProfile != null && outProfile.data != null) {
      const profileResponse: Profile = {
        Balance: outProfile.data.balance,
        Email: outProfile.data.email,
        UserId: outProfile.data.user_id,
        UserName: outProfile.data.user_name,
        WalletNumber: outProfile.data.wallet_number,
      };

      localStorage.setItem('wallet_number', outProfile.data.wallet_number);
      setProfileResponse(profileResponse);
    }
  }, [outProfile, outError]);

  const [leaderboard, setLeaderBoard] = useState<LeaderboardResp[]>([]);
  const { out: outLeaderboard, error: errorLeaderboard } = useFetchGet<{
    data: LeaderboardResp[];
  }>(`http://localhost:8000/games/leaderboard`, token!);

  useEffect(() => {
    if (errorLeaderboard) {
      const errorMessage =
        errorLeaderboard.response?.data || errorLeaderboard.message;
      notifyError(JSON.stringify(errorMessage));
      return;
    }

    if (outLeaderboard != null && outLeaderboard.data != null) {
      setLeaderBoard(outLeaderboard.data);
    }
  }, [outLeaderboard, errorLeaderboard]);

  return (
    <div className="games">
      <div className="games__container">
        <Title idName="games-title" contentProps={profileResponse!} />
        <div className="games__container__table">
          <h1>Leaderboard</h1>
          <NavLink to="/games">Back to Games</NavLink>
          {leaderboard.map((board) => (
            <h3 key={board.Name}>{board.Name}</h3>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Leaderboard;
