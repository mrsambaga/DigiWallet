import React, { useEffect, useState } from 'react';
import '../styles/home/home.css';
import useFetchGet from '../hooks/useFetchGet';
import { notifyError } from '../components/notification';
import { DropdownOption, ProfileResponse } from '../types/types';
import Title from '../components/title';
import { GetCookie } from '../helper/cookies';
import TransactionTable from '../components/table';
import Dropdown from '../components/dropDown';
import Form from '../components/form';

export type QueryParams = {
  sort: string;
  sortBy: string;
  search: string;
};

const Home: React.FC = () => {
  const [, setShow] = useState('');
  const changeShow = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setShow(event.target.value);
  };

  const [sort, setSort] = useState('');
  const changeSort = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSort(event.target.value);
  };

  const [sortBy, setSortBy] = useState('');
  const changeSortBy = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSortBy(event.target.value);
  };

  const [search, setSearch] = useState('');
  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

  const query: QueryParams = {
    sort: sort,
    sortBy: sortBy,
    search: search,
  };

  const [profileResponse, setProfileResponse] = useState<ProfileResponse>({
    Balance: 0,
    Email: '',
    UserId: 0,
    UserName: '',
    WalletNumber: 0,
  });
  const token = GetCookie('token');
  const { out, error } = useFetchGet(`http://localhost:8000/profile`, token);

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

  const showDropdown: DropdownOption[] = [
    {
      value: 'last-10-transaction',
      content: 'Last 10 Transaction',
    },
    {
      value: 'this month',
      content: 'This Month',
    },
    {
      value: 'last month',
      content: 'Last Month',
    },
    {
      value: 'this year',
      content: 'This Year',
    },
    {
      value: 'last year',
      content: 'Last year',
    },
  ];

  const sortByDropdown: DropdownOption[] = [
    {
      value: 'created_at',
      content: 'Date',
    },
    {
      value: 'amount',
      content: 'Amount',
    },
  ];

  const sortDropdown: DropdownOption[] = [
    {
      value: 'ASC',
      content: 'Ascending',
    },
    {
      value: 'DESC',
      content: 'Descending',
    },
  ];

  return (
    <div className="home">
      <div className="home__container">
        <Title contentProps={profileResponse!} />
        <div className="home__container__sorting">
          <Dropdown
            label="Show"
            onChange={changeShow}
            dropdownOptions={showDropdown}
          />
          <div className="home__container__sorting__right">
            <Dropdown
              label="Sort by"
              onChange={changeSortBy}
              dropdownOptions={sortByDropdown}
            />
            <Dropdown
              label=" "
              onChange={changeSort}
              dropdownOptions={sortDropdown}
            />
            <Form
              label=""
              placeholder="Search"
              onChangeHandler={handleSearchChange}
              inputType="text"
            />
          </div>
        </div>
        <div className="home__container__table">
          <TransactionTable QueryParams={query} />
        </div>
      </div>
    </div>
  );
};

export default Home;
