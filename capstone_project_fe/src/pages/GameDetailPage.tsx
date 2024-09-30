import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { GameService } from "../services/game/game.service";
import { GameDetailAPI } from "../services/game/game.type";
import { Typography } from "@material-tailwind/react";
import GameDetail from "../components/game/GameDetail";

const GameDetailPage: React.FC = ({}) => {
  const { gameId } = useParams();
  const [gameDetail, setGameDetail] = useState<GameDetailAPI | null>(null);

  const fetchGameDetail = async (gameId: string) => {
    try {
      const data = await GameService.get(gameId);

      setGameDetail(data);
    } catch (error) {
      console.log("🚀 ~ fetchGameDetail ~ error:", error);
    }
  };

  useEffect(() => {
    gameId && fetchGameDetail(gameId);
  }, [gameId]);

  if (!gameDetail) {
    return (
      <Typography variant="h3" className="text-center">
        Game not found with this id
      </Typography>
    );
  }

  return <GameDetail game={gameDetail} />;
};

export default GameDetailPage;
