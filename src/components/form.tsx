import React from 'react';
import '../styles/form/form.css';

type FormProps = {
  label: string;
  placeholder: string;
  value: string | number;
  onChangeHandler: (value: React.ChangeEvent<HTMLInputElement>) => void;
};

const Form: React.FC<FormProps> = ({
  label,
  placeholder,
  value,
  onChangeHandler,
}) => {
  return (
    <div className="form">
      <label className="form__label">{label}</label>
      <div className="form__container">
        <input
          type="text"
          placeholder={placeholder}
          value={value}
          onChange={onChangeHandler}
          className="form__container__input"
        />
      </div>
    </div>
  );
};

export default Form;
