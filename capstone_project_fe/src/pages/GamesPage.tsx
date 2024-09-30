import GameList from "../components/game/GameList";

const GamesPage: React.FC = () => {
  return (
    <div className="flex-1 overflow-auto">
      <GameList />
    </div>
  );
};

export default GamesPage;
