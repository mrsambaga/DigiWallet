import React, { useEffect, useState } from 'react';
import useFetchGet from '../hooks/useFetchGet';
import { TransactionDetail, TransferResponse } from '../types/types';
import { GetCookie } from '../helper/cookies';
import { notifyError } from './notification';
import queryString from 'query-string';
import '../styles/table/table.css';
import { QueryParams } from '../pages/home';
import moment from 'moment';

type DropdownProps = {
  QueryParams?: QueryParams;
};

const TransactionTable: React.FC<DropdownProps> = ({ QueryParams }) => {
  const [transactions, setTransactions] = useState<TransactionDetail[]>([]);
  const token = GetCookie('token');
  const [queryParams, setQueryParams] = useState('');
  const [paramChange, setParamChange] = useState(false);
  const { out, loading, error } = useFetchGet<{ data: TransferResponse[] }>(
    `http://localhost:8000/profile/transaction?${queryParams}`,
    token!,
    paramChange,
  );

  useEffect(() => {
    const queryParams = queryString.stringify({
      search: QueryParams?.search,
      sortBy: QueryParams?.sortBy,
      sort: QueryParams?.sort,
    });
    setQueryParams(queryParams);
    setParamChange(!paramChange);
  }, [QueryParams]);

  useEffect(() => {
    if (error) {
      const errorMessage = error.response?.data || error.message;
      notifyError(JSON.stringify(errorMessage));
      return;
    }

    if (out != null && out.data != null) {
      console.log(out);
      const transactionDetail: TransactionDetail[] = out.data.map((item) => {
        const selfWallet: string | null = localStorage.getItem('wallet_number');
        const dateTime = moment(item.CreatedAt).format('HH:mm - DD MMMM YYYY');
        return {
          TransactionId: item.TransactionId,
          Amount: item.Amount,
          Description: item.Description ? item.Description : '',
          FromTo: item.SourceId ? item.SourceId : item.TargetWalletNumber,
          Type:
            item.SourceId || item.TargetWalletNumber == selfWallet
              ? 'Credit'
              : 'Debit',
          DateTime: dateTime,
        };
      });

      setTransactions(transactionDetail);
    }
  }, [out, error]);

  return (
    <div>
      {loading ? (
        <>
          <div>
            <h1>Loading Transaction History.....</h1>
          </div>
        </>
      ) : (
        <table className="table table-bordered table-striped">
          <thead className="table__head">
            <tr>
              <th>Date & Time</th>
              <th>Type</th>
              <th>From/To</th>
              <th>Description</th>
              <th>Amount</th>
            </tr>
          </thead>
          <tbody className="table__body">
            {transactions.map((transaction) => (
              <tr key={transaction.TransactionId}>
                <td>{transaction.DateTime}</td>
                <td>{transaction.Type}</td>
                <td>{transaction.FromTo}</td>
                <td>{transaction.Description}</td>
                <td>{transaction.Amount}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default TransactionTable;
