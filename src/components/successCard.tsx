import React from 'react';
import '../styles/successCard/successCard.css';
import successLogo from '../img/icon-success.png';
import { TransactionResponse } from '../types/types';

type CardProps = {
  toggleSuccess: (success: boolean) => void;
  contentProps: TransactionResponse;
  type: string;
};

const SuccessCard: React.FC<CardProps> = ({
  toggleSuccess,
  contentProps,
  type,
}) => {
  const onClickClose = () => {
    toggleSuccess(false);
  };

  let cardType = '';
  if (type == 'topup') {
    cardType = 'Top Up';
  } else if (type == 'transfer') {
    cardType = 'Transfer';
  }

  return (
    <div className="card">
      <div className="card__container">
        <h3 className="card__title">{cardType}</h3>
        <div className="card__container__content">
          <img src={successLogo} alt="success" className="card__img" />
          <h3>{cardType} Success</h3>
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
            <button>Print</button>
            <button onClick={onClickClose}>Close</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SuccessCard;
