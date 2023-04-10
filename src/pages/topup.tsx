import React, { useEffect, useState } from 'react';
import Form from '../components/form';
import Button from '../components/button';
import Dropdown from '../components/dropDown';
import useFetchPost from '../hooks/useFetchPost';
import '../styles/topup/topup.css';

type TopupForm = {
  amount: number | string;
  source_of_funds_id: number;
};

const Topup: React.FC = () => {
  const walletNumber = localStorage.getItem('wallet_number');
  const token = localStorage.getItem('token');
  const [sourceFunds, setSourceFunds] = useState(1);
  const [amount, setAmount] = useState('');
  const [submit, setSubmit] = useState(false);

  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(event.target.value);
  };

  const handleCategoryChange = (
    event: React.ChangeEvent<HTMLSelectElement>,
  ) => {
    const category: string = event.target.value;
    if (category == 'bank transfer') {
      setSourceFunds(1);
    } else if (category == 'credit card') {
      setSourceFunds(2);
    } else if (category == 'cash') {
      setSourceFunds(3);
    }
  };

  const handleClickTopup = () => {
    if (!submit) {
      setSubmit(true);
    }
  };

  const submitForm: TopupForm = {
    amount: parseInt(amount),
    source_of_funds_id: sourceFunds,
  };

  const { out } = useFetchPost(
    'http://localhost:8000/users/topup',
    submitForm,
    submit,
    () => setSubmit(false),
    token!,
  );

  useEffect(() => {
    console.log(out);
  }, [out]);
  return (
    <div className="topup">
      <div className="topup__container">
        <h1>Top Up</h1>
        <Dropdown label="From" onChange={handleCategoryChange} />
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
    </div>
  );
};

export default Topup;
