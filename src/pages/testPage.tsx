import React, { useEffect, useState } from 'react';
import useFetchPost from '../hooks/useFetchPost';
import { notifySuccess, NotifContainer } from '../components/notification';
import 'react-toastify/dist/ReactToastify.css';

const Test: React.FC = () => {
  const [submit, setSubmit] = useState(false);
  const { data, loading, error } = useFetchPost(
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
    if (data != null) {
      notifySuccess(data.name);
    }
  }, [data]);

  return (
    <div className="post-container">
      <div>
        <button onClick={clickHandler}>Post</button>
      </div>
      {error && <div>Error: {error.message}</div>}
      {loading && <div>Loading...</div>}
      {data && (
        <div key={data.data.user_id}>
          <p>{data.data.name}</p>
          <p>{data.data.email}</p>
          <p>{data.data.password}</p>
          <p>{data.data.wallet_number}</p>
          <p>{data.data.balance}</p>
        </div>
      )}

      <NotifContainer />
    </div>
  );
};

export default Test;
