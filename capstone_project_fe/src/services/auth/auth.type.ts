export type LoginPayload = {
  identifier: string;
  password: string;
};

export type LoginResponse = {
  token: string;
  username: string;
  avatarUrl: string;
}

export type UserResponse = {
  id: number;
  username: string;
  avatarUrl: string;
}