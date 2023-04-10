import React from 'react';
import '../styles/successCard/successCard.css';
import successLogo from '../img/icon-success.png';

const SuccessCard: React.FC = () => {
  return (
    <div className="card">
      <div className="card__container">
        <h3 className="card__title">Top Up</h3>
        <div className="card__container__content">
          <img src={successLogo} alt="success" />
          <h3>Top Up Success</h3>
          <p>Amount</p>
          <p>Transaction Id</p>
          <p>From</p>
          <p>To</p>
          <p>Description</p>
          <div className="card__button">
            <button>Print</button>
            <button>Close</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SuccessCard;
