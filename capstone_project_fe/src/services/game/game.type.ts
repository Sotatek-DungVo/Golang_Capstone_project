import { UserResponse } from "../auth/auth.type";

export type GameListParams = {
  page: number;
  pageSize: number;
};

export type GameInfoDetail = {
  id: number;
  name: string;
  maxMember: number;
  startTime: Date;
  endTime: Date;
  gameOwner: UserResponse;
  gameCategory: {
    id: number;
    name: string;
    imageUrl: string;
  };
};

enum GameRequestStatus {
  PENDING = "pending",
  APPROVE = "approve",
}

export type GameRequest = {
  id: number;
  status: GameRequestStatus;
  user: UserResponse;
};

export type GameDetailAPI = GameInfoDetail & {
  gameRequests: GameRequest[];
  createdAt: Date;
};

export type CreateGamePayload = {
  startTime: Date | null;
  endTime: Date | null;
  gameCategoryId: number | null;
  maxMember: number | undefined;
  requiredSkills: number[];
  name: string | undefined;
};
