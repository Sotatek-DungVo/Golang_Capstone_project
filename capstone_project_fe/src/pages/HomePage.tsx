import { useEffect, useState } from "react";
import PlayerList from "../components/player/PlayerList";
import SearchPlayer from "../components/player/SearchPlayer";
import { PlayerService } from "../services/player/player.service";
import {
  PlayerInfoDetail,
  PlayerListParams,
} from "../services/player/player.type";
import { Button } from "@material-tailwind/react";
import { FaArrowLeft } from "react-icons/fa";
import { FaArrowRight } from "react-icons/fa";
import { toast } from "react-toastify";

const HomePage: React.FC = () => {
  const [playerData, setPlayerData] = useState<PlayerInfoDetail[] | null>(null);
  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(20);
  const fetchPlayers = async (params: PlayerListParams) => {
    try {
      const data = await PlayerService.all(params);

      setPlayerData(data);
    } catch (error: any) {
      if (error.message) {
        toast.error(error.message);
      }
      console.log("🚀 ~ fetchPlayers ~ error:", error);
    }
  };

  useEffect(() => {
    fetchPlayers({ page, pageSize });
  }, [page, pageSize]);

  return (
    <>
      <SearchPlayer fetchPlayer={fetchPlayers} />
      <div className="flex justify-between">
        <h3 className="px-10 pt-10 font-bold text-red-500">Players</h3>

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
            disabled={!playerData}
          >
            <FaArrowRight strokeWidth={2} className="w-4 h-4" /> Next
          </Button>
        </div>
      </div>
      <PlayerList playerData={playerData} />
    </>
  );
};

export default HomePage;
