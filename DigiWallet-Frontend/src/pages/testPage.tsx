import React, { useEffect, useState } from 'react';
import useFetchPost from '../hooks/useFetchPost';
import { notifySuccess, NotifContainer } from '../components/notification';
import 'react-toastify/dist/ReactToastify.css';

const Test: React.FC = () => {
  const [submit, setSubmit] = useState(false);
  const { out, loading, error } = useFetchPost(
    'http://localhost:8000/register',
    {
      name: '332432',
      email: '12231223@shopee.com',
      password: '123123123',
    },
    submit,
    () => {
      setSubmit(false);
    },
  );

  const clickHandler = () => {
    if (!submit) {
      setSubmit(true);
    }
  };

  useEffect(() => {
    if (out != null) {
      notifySuccess(out.name);
    }
  }, [out]);

  return (
    <div className="post-container">
      <div>
        <button onClick={clickHandler}>Post</button>
      </div>
      {error && <div>Error: {error.message}</div>}
      {loading && <div>Loading...</div>}
      {out && (
        <div key={out.data.user_id}>
          <p>{out.data.name}</p>
          <p>{out.data.email}</p>
          <p>{out.data.password}</p>
          <p>{out.data.wallet_number}</p>
          <p>{out.data.balance}</p>
        </div>
      )}

      <NotifContainer />
    </div>
  );
};

export default Test;
