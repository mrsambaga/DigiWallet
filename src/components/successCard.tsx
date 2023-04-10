import React from 'react';
import '../styles/successCard/successCard.css';
import successLogo from '../img/icon-success.png';
import { TransactionResponse } from '../types/types';

type ButtonProps = {
  toggleSuccess: (success: boolean) => void;
  contentProps: TransactionResponse;
};

const SuccessCard: React.FC<ButtonProps> = ({
  toggleSuccess,
  contentProps,
}) => {
  const onClickClose = () => {
    toggleSuccess(false);
  };

  return (
    <div className="card">
      <div className="card__container">
        <h3 className="card__title">Top Up</h3>
        <div className="card__container__content">
          <img src={successLogo} alt="success" className="card__img" />
          <h3>Top Up Success</h3>
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
