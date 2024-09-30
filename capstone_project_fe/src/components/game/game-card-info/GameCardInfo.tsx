import {
  GameDetailAPI,
  GameInfoDetail,
} from "../../../services/game/game.type";
import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Typography,
  Tooltip,
  Avatar,
} from "@material-tailwind/react";
import { formatDateTime } from "../../../utils/handle-date-time";
import GamePlayer from "./GamePlayer";
import GameTime from "./GameTime";
import { useNavigate } from "react-router-dom";

type GameCardInfoProps = {
  gameDetail: GameDetailAPI;
};

const GameCardInfo: React.FC<GameCardInfoProps> = ({ gameDetail }) => {
  const {
    id,
    endTime,
    startTime,
    maxMember,
    name,
    gameCategory,
    gameOwner,
    gameRequests,
  } = gameDetail;
  const navigate = useNavigate();

  return (
    <Card
      onClick={() => {
        navigate(`${id}`);
      }}
      className="transition-all duration-300 cursor-pointer w-72 hover:scale-110"
    >
      <CardHeader floated={false} className="h-30">
        <img src={gameCategory.imageUrl} alt={gameCategory.name} />
      </CardHeader>
      <CardBody className="text-center">
        <Typography variant="h5" color="blue-gray" className="mb-2">
          {name}
        </Typography>
        <div className="flex justify-between">
          <GameTime time={startTime} />
          <GameTime time={endTime} />
        </div>
      </CardBody>
      <CardFooter className="flex justify-center pt-2 gap-7">
        <GamePlayer gameRequests={gameRequests} />

        <div className="flex flex-col">
          <Typography variant="paragraph" color="blue-gray" className="mb-2">
            Created by
          </Typography>
          <Tooltip content={gameOwner.username}>
            <Avatar
              size="xl"
              variant="circular"
              className="border-2 border-white"
              src={gameOwner.avatarUrl}
              alt={gameOwner.username}
            />
          </Tooltip>
        </div>
      </CardFooter>
    </Card>
  );
};

export default GameCardInfo;
