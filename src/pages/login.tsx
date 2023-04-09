import React, { useState, useEffect, useContext } from 'react';
import '../styles/login/login.css';
import loginLogo from '../img/login-logo.png';
import Form from '../components/form';
import Button from '../components/button';
import useFetchPost from '../hooks/useFetchPost';
import { NotifContainer, notifyError } from '../components/notification';
import { AuthContext } from '../context/authContext';
import { useNavigate } from 'react-router-dom';

type LoginForm = {
  email: string;
  password: string;
};

const Login: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [submit, setSubmit] = useState(false);
  const { setAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleClickSubmit = () => {
    if (!submit) {
      setSubmit(true);
    }
  };

  const submitForm: LoginForm = {
    email: email,
    password: password,
  };

  const { data, error } = useFetchPost(
    'http://localhost:8000/login',
    submitForm,
    submit,
    () => setSubmit(false),
  );

  useEffect(() => {
    if (error != null) {
      notifyError(error.response?.data?.message || error.message);
    } else if (data != null) {
      localStorage.setItem('token', data.token);
      setAuthenticated(true);
      navigate(`/home`);
    }
  }, [data, error]);

  return (
    <div className="login">
      <div className="login__container">
        <div className="login__container__left">
          <div className="login__form">
            <h3>Login</h3>
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
      <NotifContainer />
    </div>
  );
};

export default Login;
