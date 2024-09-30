import api from "../../config/axios-instance";
import { LoginPayload, LoginResponse } from "./auth.type";

export const AuthService = {
  async login(loginPayload: LoginPayload) {
    try {
      const res = await api.post("/auth/login", loginPayload);

      return res.data;
    } catch (error) {
      return error;
    }
  },
};
