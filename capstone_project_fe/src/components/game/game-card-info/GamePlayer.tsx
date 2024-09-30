import { Avatar, Tooltip, Typography } from "@material-tailwind/react";
import { GameRequest } from "../../../services/game/game.type";

type GamePlayerProps = {
  gameRequests: GameRequest[] | null;
};

const GamePlayer: React.FC<GamePlayerProps> = ({ gameRequests }) => {
  console.log("🚀 ~ gameRequests:", gameRequests)
  return (
    <div className="flex items-center justify-between">
      {/* already joined game */}
      <div className="flex flex-col items-center gap-y-2">
        {gameRequests && gameRequests.length > 0 && (
          <Typography variant="paragraph">Joined Player</Typography>
        )}
        <div className="flex items-center -space-x-3">
          {gameRequests && gameRequests.length > 0 && 
            gameRequests.slice(0, 4).map((req, index) => {
              return (
                <Tooltip
                  key={`${req.user.username}-${index}`}
                  content={req.user.username}
                >
                  <Avatar
                    size="sm"
                    variant="circular"
                    alt={req.user.username}
                    src={req.user.avatarUrl}
                    className="border-2 border-white hover:z-10"
                  />
                </Tooltip>
              );
            })}
        </div>
      </div>
    </div>
  );
};

export default GamePlayer;
