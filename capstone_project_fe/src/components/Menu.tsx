import { FaGamepad } from "react-icons/fa";
import { FaUser } from "react-icons/fa";
import { useNavigate } from "react-router-dom";

type MenuItem = {
  icon: React.ReactNode;
  name: string;
  url: string;
};
const MENU: MenuItem[] = [
  {
    icon: <FaGamepad className="w-5 h-5" />,
    name: "Games",
    url: "/games",
  },
  {
    icon: <FaUser className="w-5 h-5" />,
    name: "Home",
    url: "/",
  },
];

const Menu: React.FC = () => {
  const navigate = useNavigate();

  return (
    <>
      <div className="p-4 mb-2">
        <h5 className="block font-sans text-xl antialiased font-semibold leading-snug tracking-normal text-blue-gray-900">
          Menu
        </h5>
      </div>

      <nav className="flex min-w-[240px] flex-col gap-1 p-2 font-sans text-base font-normal text-blue-gray-700">
        {MENU.map((menuItem, index) => {
          return (
            <div
              key={`${menuItem.name}-${index}`}
              role="button"
              className="flex items-center w-full p-3 leading-tight transition-all rounded-lg outline-none text-start hover:bg-blue-gray-50 hover:bg-opacity-80 hover:text-blue-gray-900 focus:bg-blue-gray-50 focus:bg-opacity-80 focus:text-blue-gray-900 active:bg-blue-gray-50 active:bg-opacity-80 active:text-blue-gray-900"
              onClick={() => navigate(menuItem.url)}
            >
              <div className="grid mr-4 place-items-center">
                {menuItem.icon}
              </div>
              {menuItem.name}
            </div>
          );
        })}
      </nav>
    </>
  );
};

export default Menu;
