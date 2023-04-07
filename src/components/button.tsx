import React from 'react';
import '../styles/button/button.css';

type ButtonProps = {
  label: string;
  onClickHandler: () => void;
};

const Button: React.FC<ButtonProps> = ({ label, onClickHandler }) => {
  return (
    <button onClick={onClickHandler} className="button">
      {label}
    </button>
  );
};

export default Button;
