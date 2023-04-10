import React, { useEffect, useState } from 'react';
import '../styles/home/home.css';
import useFetchGet from '../hooks/useFetchGet';
import { decodeToken } from 'react-jwt';
import { notifyError } from '../components/notification';

type Claims = {
  id: number;
  email: string;
  exp: number;
  iat: number;
  iss: string;
};

const formatRupiah = (angka: string) => {
  const number_string = angka.replace(/[^,\d]/g, '').toString();
  const split = number_string.split(',');
  const sisa = split[0].length % 3;
  let rupiah = split[0].substr(0, sisa);
  const ribuan = split[0].substr(sisa).match(/\d{3}/gi);

  // tambahkan titik jika yang di input sudah menjadi angka ribuan
  if (ribuan) {
    const separator = sisa ? '.' : '';
    rupiah += separator + ribuan.join('.');
  }

  rupiah = split[1] != undefined ? rupiah + ',' + split[1] : rupiah;
  return rupiah;
};

const Home: React.FC = () => {
  const [formattedBalance, setFormattedBalance] = useState('');
  const token = localStorage.getItem('token');
  const claims: Claims | null = token ? decodeToken(token!) : null;
  const userId = claims?.id;
  const { out, error } = useFetchGet(
    `http://localhost:8000/users/${userId}`,
    token!,
  );

  useEffect(() => {
    if (error) {
      notifyError(error.response?.data?.message || error.message);
      return;
    }

    if (out != null) {
      localStorage.setItem('wallet_number', out.data.wallet_number);
      const balanceStr = String(out.data.balance);
      const formatBalance = formatRupiah(balanceStr);
      setFormattedBalance(formatBalance);
    }
  }, [out, error]);

  return (
    <div className="home">
      <div className="home__container">
        <div className="home__container__title">
          <div className="home__container__title__left">
            <h3 className="greeter">Good morning, {out?.data.user_name}</h3>
            <p className="walletAcc">Account: {out?.data.wallet_number}</p>
          </div>
          <div className="home__container__title__right">
            <p className="balance-title">Balance:</p>
            <h3 className="balance">IDR {formattedBalance},00</h3>
          </div>
        </div>
        <div className="home__container__table">Ini Table</div>
      </div>
    </div>
  );
};

export default Home;
