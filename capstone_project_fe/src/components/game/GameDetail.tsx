import { Button, Typography } from "@material-tailwind/react";
import { GameDetailAPI } from "../../services/game/game.type";
import {
  formatDateTime,
  getDateFormat,
  getHourFormat,
} from "../../utils/handle-date-time";
import JoinedPlayerList from "./JoinedPlayerList";
import { GameService } from "../../services/game/game.service";
import { useState } from "react";
import { toast } from "react-toastify";

type GameDetailProps = {
  game: GameDetailAPI;
};

const GameDetail: React.FC<GameDetailProps> = ({ game }) => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const handleClickJoinGame = async () => {
    try {
      setIsLoading(true);
      const data = await GameService.joinGame(game.id);

      if (data) {
        toast.success("Request to join this game is created");
      }

      setIsLoading(false);
    } catch (error: any) {
      toast.error(error.message);
      setIsLoading(false);
    }
  };

  return (
    <div className="flex items-start justify-start px-60 gap-x-5">
      {/* Game Images */}
      <div className="flex flex-col max-w-80 gap-y-4">
        <img
          src={game.gameCategory.imageUrl}
          alt={game.gameCategory.name}
          className="object-cover object-center w-full rounded-lg h-80"
        />

        <div className="flex items-center justify-start gap-x-2">
          <Typography variant="paragraph" color="blue-gray">
            CREATED AT :
          </Typography>
          <Typography variant="paragraph">
            {getDateFormat(game.createdAt)}
          </Typography>
          <Typography variant="paragraph">
            {getHourFormat(game.createdAt)}
          </Typography>
        </div>
      </div>
      {/* Game Info */}

      <div className="flex flex-col flex-1 max-w-full gap-y-4">
        <div className="flex justify-between">
          <Typography variant="h3">{game.name}</Typography>

          <Button
            loading={isLoading}
            className="w-32"
            color="red"
            onClick={handleClickJoinGame}
          >
            Join
          </Button>
        </div>

        <div className="flex justify-start gap-x-10">
          <div className="flex flex-col gap-y-1">
            <Typography className="text-xs font-bold" color="gray">
              MAX MEMBER
            </Typography>
            <Typography color="red" className="text-sm">
              {game.maxMember} people
            </Typography>
          </div>
          <div className="flex flex-col gap-y-1">
            <Typography className="text-xs font-bold" color="gray">
              START AT
            </Typography>
            <Typography color="red" className="text-sm">
              {formatDateTime(game.startTime)}
            </Typography>
          </div>
          <div className="flex flex-col gap-y-1">
            <Typography className="text-xs font-bold" color="gray">
              END AT{" "}
            </Typography>
            <Typography color="red" className="text-sm">
              {formatDateTime(game.endTime)}
            </Typography>
          </div>{" "}
        </div>

        <div className="w-2/3 border-t-2 gray-500 border-"></div>
        <Typography variant="h6">Game player</Typography>

        <JoinedPlayerList gameRequests={game.gameRequests} />

        <div className="w-2/3 border-t-2 gray-500 border-"></div>
        <Typography variant="h6">Required skills</Typography>
      </div>

      {/*  */}
    </div>
  );
};

export default GameDetail;
