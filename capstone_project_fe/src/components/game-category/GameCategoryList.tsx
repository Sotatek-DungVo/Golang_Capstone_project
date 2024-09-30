import { GameCategoryDetail } from "../../services/game-category/game-category.type";

type GameCategoryListProps = {
  gameCategories: GameCategoryDetail[];
};
const GameCategoryList: React.FC<GameCategoryListProps> = ({
  gameCategories,
}) => {
  return (
    <>
      <div className="p-4 mb-2">
        <h5 className="block font-sans text-xl antialiased font-semibold leading-snug tracking-normal text-blue-gray-900">
          Game Category
        </h5>
      </div>
      <nav className="flex min-w-[240px] flex-col gap-1 p-2 font-sans text-base font-normal text-blue-gray-700">
        {gameCategories.map((category, index) => {
          return (
            <div
              key={index}
              className="flex items-center w-full p-3 leading-tight transition-all rounded-lg outline-none text-start hover:bg-blue-gray-50 hover:bg-opacity-80 hover:text-blue-gray-900 focus:bg-blue-gray-50 focus:bg-opacity-80 focus:text-blue-gray-900 active:bg-blue-gray-50 active:bg-opacity-80 active:text-blue-gray-900"
            >
              <div className="grid mr-4 place-items-center">
                <img
                  className="w-9 h-9"
                  src={category.imageUrl}
                  alt={category.name}
                />
              </div>
              {category.name}
            </div>
          );
        })}
      </nav>
    </>
  );
};

export default GameCategoryList;
