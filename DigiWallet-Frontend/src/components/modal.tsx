import React from 'react';
import '../styles/successCard/successCard.css';
import successLogo from '../img/icon-success.png';
import { Transaction } from '../types/types';
import { useNavigate } from 'react-router-dom';

type CardProps = {
  toggleSuccess: (success: boolean) => void;
  contentProps: Transaction;
  label: string;
};

const SuccessCard: React.FC<CardProps> = ({
  toggleSuccess,
  contentProps,
  label,
}) => {
  const navigate = useNavigate();
  const onClickModal = () => {
    navigate('/');
    toggleSuccess(false);
  };

  return (
    <div className="card">
      <div className="card__container">
        <h3 className="card__title">{label}</h3>
        <div className="card__container__content">
          <img src={successLogo} alt="success" className="card__img" />
          <h3>{label} Success</h3>
          <div className="card__container__content__p">
            <div className="card__container__content__p__left">
              <p>Amount</p>
              <p>Transaction Id</p>
              <p>From</p>
              <p>To</p>
              <p>Description</p>
            </div>
            <div className="card__container__content__p__right">
              <p>{contentProps.Amount}</p>
              <p>{contentProps.TransactionId}</p>
              <p>{contentProps.From}</p>
              <p>{contentProps.To}</p>
              <p>{contentProps.Description}</p>
            </div>
          </div>
          <div className="card__button">
            <button onClick={onClickModal}>Print</button>
            <button onClick={onClickModal}>Close</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SuccessCard;
