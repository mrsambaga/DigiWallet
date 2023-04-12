import React, { useEffect, useState } from 'react';
import Form from '../components/form';
import Button from '../components/button';
import useFetchPost from '../hooks/useFetchPost';
import SuccessCard from '../components/modal';
import '../styles/topup/topup.css';
import { NotifContainer, notifyError } from '../components/notification';
import { GetCookie } from '../function/cookies';
import { TransferResponse } from '../types/types';

type TransferForm = {
  amount: number | string;
  target_wallet_number: number;
  description: string;
};

const Transfer: React.FC = () => {
  const walletNumber = localStorage.getItem('wallet_number');
  const token = GetCookie('token');
  const [targetWallet, setTargetWallet] = useState<number | null>(null);
  const [description, setDescription] = useState('');
  const [amount, setAmount] = useState('');
  const [submit, setSubmit] = useState(false);
  const [success, setSuccess] = useState(false);
  const [transferResponse, setTransferResponse] = useState<TransferResponse>({
    Amount: 0,
    TransactionId: 0,
    From: 0,
    To: 0,
    Description: '',
  });

  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(event.target.value);
  };

  const handleDescChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setDescription(event.target.value);
  };

  const handleTargetWalletChange = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const targetWalletInt = parseInt(event.target.value);
    setTargetWallet(targetWalletInt);
  };

  const handleClickTransfer = () => {
    if (!submit) {
      setSubmit(true);
    }
  };

  const transferBody: TransferForm = {
    amount: parseInt(amount),
    target_wallet_number: targetWallet!,
    description: description,
  };

  const { out, error } = useFetchPost(
    'http://localhost:8000/transfer',
    transferBody,
    submit,
    () => setSubmit(false),
    token!,
  );

  useEffect(() => {
    if (error != null) {
      notifyError(error.response?.data?.message || error.message);
    } else if (out != null) {
      const transferResponse: TransferResponse = {
        Amount: out.data.amount,
        TransactionId: out.data.transaction_id,
        From: Number(walletNumber),
        To: out.data.target_wallet_number,
        Description: out.data.description,
      };
      setTransferResponse(transferResponse);
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
          contentProps={transferResponse}
          type="transfer"
        />
      ) : (
        <div className="topup__container">
          <h1>Transfer</h1>
          <Form
            label="From"
            placeholder={walletNumber ? walletNumber : ''}
            isReadOnly={true}
          />
          <Form
            label="To"
            placeholder="1230000000001"
            value={targetWallet!}
            onChangeHandler={handleTargetWalletChange}
          />
          <Form
            label="Amount"
            placeholder="1.000.000.000"
            value={amount}
            onChangeHandler={handleAmountChange}
          />
          <Form
            label="Description"
            placeholder="Bayar Hutang"
            value={description}
            onChangeHandler={handleDescChange}
          />
          <Button label="Transfer" onClickHandler={handleClickTransfer} />
        </div>
      )}
      <NotifContainer />
    </div>
  );
};

export default Transfer;
