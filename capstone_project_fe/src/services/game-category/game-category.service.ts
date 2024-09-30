import axiosInterceptorInstance from "../../config/axios-instance";
import { GameCategoryParams } from "./game-category.type";

export const GameCategoryService = {
  async all(params: GameCategoryParams) {
    try {
      const res = await axiosInterceptorInstance.get("/game-categories", {
        params,
        withCredentials: true,
        headers: {
          Accept: "application/json",
        },
      });
      console.log("🚀 ~ all ~ res:", res);

      return res.data;
    } catch (error) {
      console.log("🚀 ~ all ~ error:", error);
    }
  },
};
