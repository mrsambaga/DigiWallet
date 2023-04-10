import React, { useEffect, useState } from 'react';
import Title from '../components/title';
import useFetchGet from '../hooks/useFetchGet';
import { notifyError } from '../components/notification';
import { ProfileResponse } from '../types/types';
import { decodeToken } from 'react-jwt';
import { Claims } from '../types/types';

const Games: React.FC = () => {
  const [profileResponse, setProfileResponse] = useState<ProfileResponse>({
    Balance: 0,
    Email: '',
    UserId: 0,
    UserName: '',
    WalletId: 0,
    WalletNumber: 0,
  });
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

    if (out != null) {
      const profileResponse: ProfileResponse = {
        Balance: out.data.balance,
        Email: out.data.email,
        UserId: out.data.user_id,
        UserName: out.data.user_name,
        WalletId: out.data.wallet_id,
        WalletNumber: out.data.wallet_number,
      };

      localStorage.setItem('wallet_number', out.data.wallet_number);
      setProfileResponse(profileResponse);
    }
  }, [out, error]);

  return (
    <div className="home">
      <div className="home__container">
        <Title idName="games" contentProps={profileResponse!} />
        <div className="home__container__table">Ini Table</div>
      </div>
    </div>
  );
};

export default Games;