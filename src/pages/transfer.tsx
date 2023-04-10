import React, { useEffect, useState } from 'react';
import Form from '../components/form';
import Button from '../components/button';
import useFetchPost from '../hooks/useFetchPost';
import SuccessCard from '../components/successCard';
import '../styles/topup/topup.css';
import { NotifContainer, notifyError } from '../components/notification';

type TransferForm = {
  amount: number | string;
  source_of_funds_id: number;
};

export type TransferResponse = {
  Amount: number;
  TransactionId: number;
  From: number;
  To: number;
  Description: string;
};

const Transfer: React.FC = () => {
  const walletNumber = localStorage.getItem('wallet_number');
  const token = localStorage.getItem('token');
  const [sourceFunds] = useState(1);
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

  const handleClickTransfer = () => {
    if (!submit) {
      setSubmit(true);
    }
  };

  const submitForm: TransferForm = {
    amount: parseInt(amount),
    source_of_funds_id: sourceFunds,
  };

  const { out, error } = useFetchPost(
    'http://localhost:8000/users/Transfer',
    submitForm,
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
        From: out.data.source_of_funds,
        To: Number(walletNumber),
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
            value={amount}
            onChangeHandler={handleAmountChange}
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
            value={amount}
            onChangeHandler={handleAmountChange}
          />
          <Button label="Transfer" onClickHandler={handleClickTransfer} />
        </div>
      )}
      <NotifContainer />
    </div>
  );
};

export default Transfer;
