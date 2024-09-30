import { Typography } from "@material-tailwind/react";
import { getDateFormat, getHourFormat } from "../../../utils/handle-date-time";

type GameTimeProps = {
  time: Date;
};

const GameTime: React.FC<GameTimeProps> = ({ time }) => {
  return <div className="flex flex-col px-5 py-2 text-white bg-black rounded-lg gap-y-4">
    <Typography>{getHourFormat(time)}</Typography>
    <Typography>{getDateFormat(time)}</Typography>
  </div>;
};

export default GameTime;
