import { Avatar, Tooltip } from "@material-tailwind/react";
import { GameRequest } from "../../services/game/game.type";

type JoinedPlayerListProps = {
  gameRequests: GameRequest[];
};

const JoinedPlayerList: React.FC<JoinedPlayerListProps> = ({
  gameRequests,
}) => {
  return (
    <div className="flex flex-wrap justify-start w-2/3 gap-4">
      {gameRequests.length > 0 &&
        gameRequests.map((gameRequests, index) => {
          return (
            <Tooltip
              content={gameRequests.user.username}
              key={`${gameRequests.user.username}-${index}`}
            >
              <Avatar
                size="xl"
                variant="rounded"
                src={gameRequests.user.avatarUrl}
                alt={gameRequests.user.username}
                withBorder={true}
                color="green"
                className="p-0.5"

              />
            </Tooltip>
          );
        })}
    </div>
  );
};

export default JoinedPlayerList;
