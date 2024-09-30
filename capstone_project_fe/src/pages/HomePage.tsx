import { useEffect, useState } from "react";
import PlayerList from "../components/player/PlayerList";
import SearchPlayer from "../components/player/SearchPlayer";
import { PlayerService } from "../services/player/player.service";
import {
  PlayerInfoDetail,
  PlayerListParams,
} from "../services/player/player.type";

const HomePage: React.FC = () => {
  const [playerData, setPlayerData] = useState<PlayerInfoDetail[]>([]);
  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(20);
  const fetchPlayers = async (params: PlayerListParams) => {
    try {
      const data = await PlayerService.all(params);

      setPlayerData(data);
    } catch (error) {
      console.log("🚀 ~ fetchPlayers ~ error:", error);
    }
  };

  useEffect(() => {
    fetchPlayers({ page, pageSize });
  }, [page, pageSize]);

  return (
    <>
      <SearchPlayer fetchPlayer={fetchPlayers} />
      <h3 className="px-10 pt-10 font-bold text-red-500">Players</h3>
      <PlayerList playerData={playerData} />
    </>
  );
};

export default HomePage;
