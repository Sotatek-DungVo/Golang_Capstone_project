import api from "../../config/axios-instance";
import { CreateGamePayload, GameListParams } from "./game.type";


export const GameService = {
  async all(params: GameListParams) {
    try {
      const res = await api.get("/games", {
        params,
      });

      return res.data;
    } catch (error: any) {
      throw new Error(error.message);
    }
  },

  async get(gameId: string) {
    try {
      const res = await api.get(`/games/${gameId}`);

      return res.data;
    } catch (error: any) {
      throw new Error(error.message);
    }
  },

  async joinGame(gameId: string | number) {
    try {
      const res = await api.post(`/game-requests`, {
        gameId,
      });

      return res.data;
    } catch (error: any) {
      throw new Error(error.message);
    }
  },

  async create(payload: CreateGamePayload) {
    try {
      const res = await api.post("/games", payload);

      return res.data;
    } catch (error: any) {
      throw new Error(error.message);
    }
  },
};
