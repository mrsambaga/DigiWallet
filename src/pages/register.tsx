import React, { useState } from 'react';
import Navbar from '../components/navbar';
import '../styles/register/register.css';
import registerLogo from '../img/register-logo.png';
import Form from '../components/form';
import Button from '../components/button';

// type RegisterProps = {
//   name: string;
//   email: string;
//   passwod: string;
// };

const Register: React.FC = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  //   const [registerForm, setRegisterForm] = useState(null);

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleClickSubmit = () => {};

  return (
    <div className="register">
      <Navbar />
      <div className="register__container">
        <div className="register__container__left">
          <div className="register__form">
            <h3>Register</h3>
            <Form
              label="Name"
              placeholder="Asep Budiantoro Chandradiman"
              value={name}
              onChangeHandler={handleNameChange}
            />
            <Form
              label="Email"
              placeholder="asep.bc@gmail.com"
              value={email}
              onChangeHandler={handleEmailChange}
            />
            <Form
              label="Password"
              placeholder="***********"
              value={password}
              onChangeHandler={handlePasswordChange}
            />
            <Button label="Submit" onClickHandler={handleClickSubmit} />
          </div>
        </div>
        <div className="register__container__right">
          <img
            src={registerLogo}
            alt="register-logo"
            className="register__logo"
          />
        </div>
      </div>
    </div>
  );
};

export default Register;
