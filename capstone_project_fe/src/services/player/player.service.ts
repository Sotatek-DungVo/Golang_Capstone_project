import api from "../../config/axios-instance";
import { PlayerListParams } from "./player.type";

export const PlayerService = {
  async all(params: PlayerListParams) {
    try {
      const res = await api.get("/players", {
        params
      });

      return res.data;
    } catch (error) {
    console.log("🚀 ~ all ~ error:", error)
    }
  },
};
