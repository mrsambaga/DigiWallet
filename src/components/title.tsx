import { ProfileResponse } from '../types/types';
import React, { useState, useEffect } from 'react';
import '../styles/home/home.css';

type TitleProps = {
  idName?: string;
  contentProps: ProfileResponse;
};

export const formatRupiah = (angka: string): string => {
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

const Title: React.FC<TitleProps> = ({ idName, contentProps }) => {
  const [formattedBalance, setFormattedBalance] = useState('0');
  useEffect(() => {
    const balanceStr = String(contentProps.Balance);
    const formatBalance = formatRupiah(balanceStr);
    setFormattedBalance(formatBalance);
  }, [contentProps]);

  return (
    <div className="home__container__title">
      <div
        className="home__container__title__left"
        id={idName ? 'games__container__title__left' : ''}
      >
        <h3 className="greeter">Good morning, {contentProps.UserName}</h3>
        <p className="walletAcc">Account: {contentProps.WalletNumber}</p>
      </div>
      <div
        className="home__container__title__right"
        id={idName ? 'games__container__title__right' : ''}
      >
        <p className="balance-title">Balance:</p>
        <h3 className="balance">IDR {formattedBalance},00</h3>
      </div>
    </div>
  );
};

export default Title;
