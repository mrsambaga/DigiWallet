export type Transaction = {
  Amount: number;
  TransactionId: number;
  From: number;
  To: number;
  Description: string;
};

export type TransferResponse = {
  TransactionId: number;
  Amount: string;
  Description: string;
  SourceId: string;
  TargetWalletNumber: string;
  CreatedAt: number;
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
  balance: number;
  email: string;
  user_id: number;
  user_name: string;
  wallet_number: string;
};

export type Profile = {
  Balance: number;
  Email: string;
  UserId: number;
  UserName: string;
  WalletNumber: string;
};

export type Claims = {
  id: number;
  email: string;
  exp: number;
  iat: number;
  iss: string;
};

export type DropdownOption = {
  content: string;
  value: string;
};

export type Boxes = {
  BoxId: number;
  Prize: number;
};

export type LeaderboardResp = {
  Id: number;
  Name: string;
  TotalWinning: number;
};
