import React, { useEffect, useState } from 'react';
import useFetchGet from '../hooks/useFetchGet';
import { TransactionDetail } from '../types/types';
import { GetCookie } from '../function/cookies';
import { notifyError } from './notification';
import '../styles/table/table.css';

const TransactionTable: React.FC = () => {
  const [transactions, setTransactions] = useState<TransactionDetail[]>([]);
  const token = GetCookie('token');
  const { out, error } = useFetchGet(
    `http://localhost:8000/users/transaction`,
    token!,
  );

  useEffect(() => {
    if (error) {
      notifyError(error.response?.data?.message || error.message);
      return;
    }

    if (out != null) {
      const transactionDetail: TransactionDetail[] = out.data.map(
        (item: any) => {
          return {
            TransactionId: item.TransactionId,
            Amount: item.Amount,
            Description: item.Description ? item.Description : '',
            FromTo: item.SourceId ? item.SourceId : item.TargetWalletNumber,
            Type: item.SourceId ? 'Credit' : 'Debit',
            DateTime: item.CreatedAt,
          };
        },
      );

      setTransactions(transactionDetail);
    }
  }, [out, error]);

  return (
    <div>
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
    </div>
  );
};

export default TransactionTable;
