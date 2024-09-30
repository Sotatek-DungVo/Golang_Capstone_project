import { Gender } from "../../services/player/player.type"
import { MdFemale } from "react-icons/md";
import { MdMale } from "react-icons/md";

type GenderInfoProps = {
    gender: Gender
}

const GenderInfo: React.FC<GenderInfoProps> = ({ gender }) => {
    return gender === Gender.FEMALE ? <MdMale className="w-8 h-8" /> : <MdFemale className="w-8 h-8" color="#EE66A6" />
}

export default GenderInfo;