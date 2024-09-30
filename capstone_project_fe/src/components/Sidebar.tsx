import { useEffect, useState } from "react";
import { GameCategoryDetail } from "../services/game-category/game-category.type";
import GameCategoryList from "./game-category/GameCategoryList";
import { GameCategoryService } from "../services/game-category/game-category.service";
import Menu from "./Menu";
import { toast } from "react-toastify";

const Sidebar: React.FC = () => {
  const [gameCategoryList, setGameCategoryList] = useState<
    GameCategoryDetail[] | null
  >(null);
  useEffect(() => {
    const fetchGameCategories = async () => {
      try {
        const data = await GameCategoryService.all({
          page: 1,
          pageSize: 20,
        });

        setGameCategoryList(data);
      } catch (error: any) {
        if (error.message) {
          toast.error(error.message);
        }
        console.log("error:", error);
      }
    };

    fetchGameCategories();
  }, []);

  return (
    <div className="relative flex w-full max-w-[16rem] flex-col rounded-xl bg-white bg-clip-border p-4 text-gray-700 shadow-xl shadow-blue-gray-900/5">
      <Menu />

      {gameCategoryList && (
        <GameCategoryList gameCategories={gameCategoryList} />
      )}
    </div>
  );
};

export default Sidebar;
