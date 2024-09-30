import api from "../../config/axios-instance";

export const RequiredSkillsService = {
  async all() {
    try {
      const res = await api.get("/required-skills");

      return res.data;
    } catch (error) {
        console.log("🚀 ~ all ~ error:", error)
    }
  },
};
