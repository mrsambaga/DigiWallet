import React, { useState } from 'react';
import '../styles/form/form.css';

type FormProps = {
  label: string;
  placeholder: string;
  value?: string | number;
  isReadOnly?: boolean;
  onChangeHandler?: (event: React.ChangeEvent<HTMLInputElement>) => void;
};

const Form: React.FC<FormProps> = ({
  label,
  placeholder,
  value,
  onChangeHandler,
  isReadOnly = false,
}) => {
  const [activity, setActivity] = useState(value);
  const [showError, setShowError] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');

  const handleBlur = () => {
    if (!activity && label == 'Amount') {
      setShowError(true);
      setErrorMessage('Please insert amount of transfer');
    } else if (!activity && label == 'Email') {
      setShowError(true);
      setErrorMessage('Please enter your email');
    } else if (!activity && label == 'Password') {
      setShowError(true);
      setErrorMessage('Please enter your password');
    }
  };

  const handleOnChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value;
    setActivity(newValue);
    setShowError(false);
    if (onChangeHandler) {
      onChangeHandler(event);
    }
  };

  return (
    <div className="form">
      <label className="form__label">{label}</label>
      <div className="form__container">
        <input
          type="text"
          placeholder={placeholder}
          value={value}
          onChange={handleOnChange}
          className={
            isReadOnly
              ? 'form__container__input-read'
              : 'form__container__input'
          }
          id={showError ? 'show-error' : ''}
          readOnly={isReadOnly}
          onBlur={handleBlur}
        />
      </div>
      {showError && <div className="form__error">{errorMessage}</div>}
    </div>
  );
};

export default Form;
