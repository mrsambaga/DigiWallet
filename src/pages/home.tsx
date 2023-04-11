import React, { useEffect, useState } from 'react';
import '../styles/home/home.css';
import useFetchGet from '../hooks/useFetchGet';
import { decodeToken } from 'react-jwt';
import { notifyError } from '../components/notification';
import { ProfileResponse } from '../types/types';
import Title from '../components/title';
import { Claims } from '../types/types';
import { GetCookie } from '../function/cookies';
import TransactionTable from '../components/table';
import Dropdown from '../components/dropDown';
import Form from '../components/form';

const Home: React.FC = () => {
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
    token,
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

  const changeShow = () => {};

  return (
    <div className="home">
      <div className="home__container">
        <Title contentProps={profileResponse!} />
        <div className="home__container__sorting">
          <Dropdown label="" onChange={changeShow} />
          <Dropdown label="" onChange={changeShow} />
          <Form label="" placeholder="Search" />
        </div>
        <div className="home__container__table">
          <TransactionTable />
        </div>
      </div>
    </div>
  );
};

export default Home;
