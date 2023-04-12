import React, { useState, useEffect, useContext } from 'react';
import '../styles/register/register.css';
import registerLogo from '../img/register-logo.png';
import Form from '../components/form';
import Button from '../components/button';
import useFetchPost from '../hooks/useFetchPost';
import {
  NotifContainer,
  notifySuccess,
  notifyError,
} from '../components/notification';
import { AuthContext } from '../context/authContext';
import { Navigate } from 'react-router-dom';

type RegisterBody = {
  name: string;
  email: string;
  password: string;
};

const Register: React.FC = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isSubmit, setIsSubmit] = useState(false);
  const { authenticated } = useContext(AuthContext);

  if (authenticated) {
    return <Navigate to="/" replace />;
  }

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
    if (!isSubmit) {
      setIsSubmit(true);
    }
  };

  const body: RegisterBody = {
    name: name,
    email: email,
    password: password,
  };

  const { out, error } = useFetchPost(
    'http://localhost:8000/register',
    body,
    isSubmit,
    () => setIsSubmit(false),
  );

  useEffect(() => {
    if (error != null) {
      notifyError(error.response?.data?.message || error.message);
    } else if (out != null) {
      notifySuccess(out.data.name);
    }
  }, [out, error]);

  return (
    <div className="register">
      <div className="register__container">
        <div className="register__container__left">
          <div className="register__form">
            <h3>Register</h3>
            <Form
              label="Name"
              placeholder="Asep Budiantoro Chandradiman"
              value={name}
              onChangeHandler={handleNameChange}
              inputType="text"
            />
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
        <div className="register__container__right">
          <img
            src={registerLogo}
            alt="register-logo"
            className="register__logo"
          />
        </div>
      </div>
      <NotifContainer />
    </div>
  );
};

export default Register;
