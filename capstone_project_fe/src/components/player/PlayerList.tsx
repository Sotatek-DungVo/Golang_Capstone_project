import PlayerInfoCard from "./PlayerInfoCard";
import { PlayerInfoDetail } from "../../services/player/player.type";

type PlayerListProps = {
  playerData: PlayerInfoDetail[] | null;
};

const PlayerList: React.FC<PlayerListProps> = ({ playerData }) => {
  return (
    <div className="grid grid-cols-2 px-10 pt-5 2xl:grid-cols-5 xl:grid-cols-4 lg:grid-cols-3 gap-x-5 gap-y-10">
      {
        playerData && playerData.length > 0 && playerData.map((player, index) => {
          return (
            <PlayerInfoCard
              key={index}
              avatarUrl={player.avatarUrl}
              name={player.username}
              description={player.description}
              gender={player.gender}
            />
          );
        })
      }
    </div>
  );
};

export default PlayerList;
