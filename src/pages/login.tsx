import React, { useState, useEffect, useContext } from 'react';
import '../styles/login/login.css';
import loginLogo from '../img/login-logo.png';
import Form from '../components/form';
import Button from '../components/button';
import useFetchPost from '../hooks/useFetchPost';
import { NotifContainer, notifyError } from '../components/notification';
import { AuthContext } from '../context/authContext';
import { useNavigate } from 'react-router-dom';
import { SetCookie } from '../helper/cookies';
import { Navigate } from 'react-router-dom';

type LoginForm = {
  email: string;
  password: string;
};

const Login: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [submit, setSubmit] = useState(false);
  const { authenticated, setAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  if (authenticated) {
    return <Navigate to="/" replace />;
  }

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

  const body: LoginForm = {
    email: email,
    password: password,
  };

  const { out, error } = useFetchPost(
    'http://localhost:8000/login',
    body,
    submit,
    () => setSubmit(false),
  );

  useEffect(() => {
    if (error != null) {
      notifyError(error.response?.data?.message || error.message);
    } else if (out != null) {
      SetCookie('token', out.data.token, 1);
      setAuthenticated(true);
      navigate(`/`);
    }
  }, [out, error]);

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
              inputType="text"
            />
            <Form
              label="Password"
              placeholder="***********"
              value={password}
              onChangeHandler={handlePasswordChange}
              inputType="text"
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
