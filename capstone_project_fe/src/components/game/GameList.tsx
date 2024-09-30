import {
  GameDetailAPI,
  GameInfoDetail,
  GameListParams,
} from "../../services/game/game.type";
import { GameService } from "../../services/game/game.service";
import GameCardInfo from "./game-card-info/GameCardInfo";
import { Button } from "@material-tailwind/react";
import { useEffect, useState } from "react";
import GameDialogForm from "./GameDialogForm";
import { FaArrowLeft, FaArrowRight } from "react-icons/fa";
import { toast } from "react-toastify";

const GameList: React.FC = () => {
  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(20);
  const [gameData, setGameData] = useState<GameDetailAPI[] | null>(null);
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen((cur) => !cur);

  const fetchGames = async (params: GameListParams) => {
    try {
      const data = await GameService.all(params);

      setGameData(data);
    } catch (error: any) {
      if (error.message) {
        toast.error(error.message)
      }
      console.log("🚀 ~ fetchGames ~ error:", error);
    }
  };

  useEffect(() => {
    fetchGames({
      page,
      pageSize,
    });
  }, [page, pageSize]);

  return (
    <>
      <div className="flex justify-between pr-20">
        <div className="flex items-center justify-between">
          <h3 className="px-10 font-bold text-red-500">Games</h3>

          <div className="flex">
          <Button
            variant="text"
            className="flex items-center gap-2"
            onClick={() => setPage(page - 1)}
            disabled={page === 1}
          >
            <FaArrowLeft strokeWidth={2} className="w-4 h-4" /> Previous
          </Button>

          <Button
            variant="text"
            className="flex items-center gap-2"
            onClick={() => setPage(page + 1)}
            disabled={!gameData}
          >
            <FaArrowRight strokeWidth={2} className="w-4 h-4" /> Next
          </Button>
        </div>
        </div>

        <Button onClick={handleOpen} className="ml-20">
          Create Game
        </Button>
      </div>
      <div className="grid grid-cols-2 px-10 pt-5 2xl:grid-cols-5 xl:grid-cols-4 lg:grid-cols-3 gap-x-5 gap-y-10">
        {gameData && gameData.length > 0 ? (
          <>
            {gameData.map((game, index) => {
              return <GameCardInfo key={index} gameDetail={game} />;
            })}
          </>
        ) : (
          <>
            <h3 className="px-10 pt-10 font-bold text-red-500">No Data</h3>
          </>
        )}
      </div>

      <GameDialogForm open={open} handleOpen={handleOpen} />
    </>
  );
};

export default GameList;
