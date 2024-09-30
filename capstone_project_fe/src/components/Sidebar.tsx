import { useEffect, useState } from "react";
import { GameCategoryDetail } from "../services/game-category/game-category.type";
import GameCategoryList from "./game-category/GameCategoryList";
import { GameCategoryService } from "../services/game-category/game-category.service";
import Menu from "./Menu";

const Sidebar: React.FC = () => {
  const [gameCategoryList, setGameCategoryList] = useState<
    GameCategoryDetail[]
  >([]);
  useEffect(() => {
    const fetchGameCategories = async () => {
      try {
        const data = await GameCategoryService.all({
          page: 1,
          pageSize: 20,
        });

        setGameCategoryList(data);
      } catch (error) {
        console.log("error:", error);
      }
    };

    fetchGameCategories();
  }, []);

  return (
    <div className="relative flex w-full max-w-[16rem] flex-col rounded-xl bg-white bg-clip-border p-4 text-gray-700 shadow-xl shadow-blue-gray-900/5">
      <GameCategoryList gameCategories={gameCategoryList} />

      <Menu />
    </div>
  );
};

export default Sidebar;
