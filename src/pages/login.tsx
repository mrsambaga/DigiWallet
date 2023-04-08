import React, { useState } from 'react';
import Navbar from '../components/navbar';
import '../styles/login/login.css';
import loginLogo from '../img/login-logo.png';
import Form from '../components/form';
import Button from '../components/button';

// type loginProps = {
//   name: string;
//   email: string;
//   passwod: string;
// };

const Login: React.FC = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  //   const [loginForm, setLoginForm] = useState<loginProps | null>(null);

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleClickSubmit = () => {
    // const submitForm: loginProps = {
    //   name: name,
    //   email: email,
    //   passwod: password,
    // };
    // setLoginForm(submitForm);
  };

  return (
    <div className="login">
      <Navbar type="login" />
      <div className="login__container">
        <div className="login__container__left">
          <div className="login__form">
            <h3>Login</h3>
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
        <div className="login__container__right">
          <img src={loginLogo} alt="login-logo" className="login__logo" />
        </div>
      </div>
    </div>
  );
};

export default Login;
