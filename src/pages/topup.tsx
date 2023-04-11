import React, { useEffect, useState } from 'react';
import Form from '../components/form';
import Button from '../components/button';
import Dropdown from '../components/dropDown';
import useFetchPost from '../hooks/useFetchPost';
import SuccessCard from '../components/successCard';
import '../styles/topup/topup.css';
import { NotifContainer, notifyError } from '../components/notification';
import { TransactionResponse } from '../types/types';
import { GetCookie } from '../function/cookies';

type TopupForm = {
  amount: number | string;
  source_of_funds_id: number;
};

const Topup: React.FC = () => {
  const walletNumber = localStorage.getItem('wallet_number');
  const token = GetCookie('token');
  const [sourceFunds, setSourceFunds] = useState(1001);
  const [amount, setAmount] = useState('');
  const [submit, setSubmit] = useState(false);
  const [success, setSuccess] = useState(false);
  const [topupResponse, setTopupResponse] = useState<TransactionResponse>({
    Amount: 0,
    TransactionId: 0,
    From: 0,
    To: 0,
    Description: '',
  });

  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(event.target.value);
  };

  const handleCategoryChange = (
    event: React.ChangeEvent<HTMLSelectElement>,
  ) => {
    const category: string = event.target.value;
    if (category == 'bank transfer') {
      setSourceFunds(1001);
    } else if (category == 'credit card') {
      setSourceFunds(1002);
    } else if (category == 'cash') {
      setSourceFunds(1003);
    }
  };

  const handleClickTopup = () => {
    if (!submit) {
      setSubmit(true);
    }
  };

  const body: TopupForm = {
    amount: parseInt(amount),
    source_of_funds_id: sourceFunds,
  };

  const { out, error } = useFetchPost(
    'http://localhost:8000/users/topup',
    body,
    submit,
    () => setSubmit(false),
    token!,
  );

  useEffect(() => {
    if (error != null) {
      notifyError(error.response?.data?.message || error.message);
    } else if (out != null) {
      const topupResponse: TransactionResponse = {
        Amount: out.data.amount,
        TransactionId: out.data.transaction_id,
        From: out.data.source_of_funds,
        To: out.data.target_wallet_number,
        Description: out.data.description,
      };
      setTopupResponse(topupResponse);
      setSuccess(true);
    }
  }, [out, error]);

  const closeSuccessCard = () => {
    setSuccess(false);
  };

  return (
    <div className="topup" id={success ? 'topup-active' : ''}>
      {success ? (
        <SuccessCard
          toggleSuccess={closeSuccessCard}
          contentProps={topupResponse}
          type="topup"
        />
      ) : (
        <div className="topup__container">
          <h1>Top Up</h1>
          <Dropdown label="To" onChange={handleCategoryChange} />
          <Form
            label="To"
            placeholder={walletNumber ? walletNumber : ''}
            isReadOnly={true}
          />
          <Form
            label="Amount"
            placeholder="1.000.000.000"
            value={amount}
            onChangeHandler={handleAmountChange}
          />
          <Button label="Topup" onClickHandler={handleClickTopup} />
        </div>
      )}
      <NotifContainer />
    </div>
  );
};

export default Topup;
