import React, { useEffect, useState } from 'react';
import useFetchGet from '../hooks/useFetchGet';
import { TransactionDetail } from '../types/types';
import { GetCookie } from '../helper/cookies';
import { notifyError } from './notification';
import queryString from 'query-string';
import '../styles/table/table.css';
import { QueryParams } from '../pages/home';

type DropdownProps = {
  QueryParams?: QueryParams;
};

const TransactionTable: React.FC<DropdownProps> = ({ QueryParams }) => {
  const [transactions, setTransactions] = useState<TransactionDetail[]>([]);
  const token = GetCookie('token');
  const [queryParams, setQueryParams] = useState('');
  const [paramChange, setParamChange] = useState(false);
  const { out, loading, error } = useFetchGet(
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
      notifyError(error.response?.data?.message || error.message);
      return;
    }

    if (out != null) {
      const transactionDetail: TransactionDetail[] = out.data.map(
        (item: any) => {
          const selfWallet: string | null =
            localStorage.getItem('wallet_number');
          return {
            TransactionId: item.TransactionId,
            Amount: item.Amount,
            Description: item.Description ? item.Description : '',
            FromTo: item.SourceId ? item.SourceId : item.TargetWalletNumber,
            Type:
              item.SourceId || item.TargetWalletNumber == selfWallet
                ? 'Credit'
                : 'Debit',
            DateTime: item.CreatedAt,
          };
        },
      );

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
