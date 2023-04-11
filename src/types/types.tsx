export type TransactionResponse = {
  Amount: number;
  TransactionId: number;
  From: number;
  To: number;
  Description: string;
};

export type TransactionDetail = {
  TransactionId: number;
  Amount: string;
  Description: string;
  FromTo: string;
  Type: string;
  DateTime: string;
};

export type ProfileResponse = {
  Balance: number;
  Email: string;
  UserId: number;
  UserName: string;
  WalletNumber: number;
};

export type TransferResponse = {
  Amount: number;
  TransactionId: number;
  From: number;
  To: number;
  Description: string;
};

export type Claims = {
  id: number;
  email: string;
  exp: number;
  iat: number;
  iss: string;
};
