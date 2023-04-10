export type TransactionResponse = {
  Amount: number;
  TransactionId: number;
  From: number;
  To: number;
  Description: string;
};

export type ProfileResponse = {
  Balance: number;
  Email: string;
  UserId: number;
  UserName: string;
  WalletId: number;
  WalletNumber: number;
};

export type Claims = {
  id: number;
  email: string;
  exp: number;
  iat: number;
  iss: string;
};
