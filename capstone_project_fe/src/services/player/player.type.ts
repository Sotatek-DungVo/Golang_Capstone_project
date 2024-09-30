export type PlayerListParams = {
  page: number;
  pageSize: number;
  gender?: string;
  username?: string;
};

export type PlayerInfoDetail = {
  username: string;
  email: string;
  description: string;
  gender: Gender;
  avatarUrl: string;
};

export enum Gender {
  FEMALE = 'FEMALE',
  MALE = 'MALE'
}