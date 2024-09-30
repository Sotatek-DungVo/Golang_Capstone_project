import {
  Button,
  Dialog,
  Card,
  CardBody,
  CardFooter,
  Typography,
  Input,
} from "@material-tailwind/react";
import { useEffect, useState } from "react";
import DatePicker from "react-datepicker";
import CreatableSelect from "react-select/creatable";
import { RequiredSkillsService } from "../../services/required-skills/required-skills.service";
import { GameCategoryService } from "../../services/game-category/game-category.service";
import { GameCategoryDetail } from "../../services/game-category/game-category.type";
import Select from "react-select";
import { GameService } from "../../services/game/game.service";
import { toast } from "react-toastify";

type GameDialogFormProps = {
  open: boolean;
  handleOpen: () => void;
};

export interface ColourOption {
  readonly value: string;
  readonly label: string;
}

const GameDialogForm: React.FC<GameDialogFormProps> = ({
  open,
  handleOpen,
}) => {
  const [startTime, setStartTime] = useState<null | Date>(null);
  const [endTime, setEndTime] = useState<null | Date>(null);
  const [maxMember, setMaxMember] = useState<number | undefined>(undefined);
  const [name, setName] = useState<string | undefined>(undefined);
  const [requiredSkillOptions, setRequiredSkillOptions] = useState([]);
  const [gameCategories, setGameCategories] = useState([]);

  const [categoryId, setCategoryId] = useState(null);
  const [requiredSkills, setRequiredSkills] = useState<number[]>([]);

  const handleCreateGame = async () => {
    try {
      const data = await GameService.create({
        startTime,
        endTime,
        gameCategoryId: categoryId,
        maxMember,
        requiredSkills,
        name,
      });

      if (!data) {
        toast.error("Error message");
      }

      toast.success("Create new game success");
      handleOpen();
    } catch (error: any) {
      handleOpen();
      if (error.message) {
        toast.error(error.message);
      }
    }
  };

  useEffect(() => {
    const fetchSkills = async () => {
      const data = await RequiredSkillsService.all();

      setRequiredSkillOptions(data);
    };

    const fetchGameCategories = async () => {
      const data = await GameCategoryService.all({
        page: 1,
        pageSize: 20,
      });

      const categoryOptions =
        data.length === 0
          ? data
          : data.map((item: GameCategoryDetail) => {
              return {
                label: item.name,
                value: item.id,
              };
            });

      setGameCategories(categoryOptions);
    };

    fetchSkills();
    fetchGameCategories();
  }, []);

  return (
    <Dialog
      size="lg"
      open={open}
      handler={handleOpen}
      className="bg-transparent shadow-none"
    >
      <Card className="mx-auto w-full max-w-[24rem]">
        <CardBody className="flex flex-col gap-4">
          <Typography variant="h4" color="blue-gray">
            Create new game
          </Typography>
          <Typography
            className="mb-3 font-normal"
            variant="paragraph"
            color="gray"
          >
            Enter your game information.
          </Typography>
          <Typography className="-mb-2" variant="h6">
            Name
          </Typography>
          <Input
            crossOrigin={null}
            onChange={(e) => setName(e.target.value)}
            value={name}
            label="Name"
            size="lg"
          />
          <Typography className="-mb-2" variant="h6">
            Max Member
          </Typography>
          <Input
            crossOrigin={null}
            onChange={(e) => setMaxMember(parseInt(e.target.value))}
            type="number"
            label="Max member"
            value={maxMember}
            size="lg"
          />
          <Typography className="-mb-2" variant="h6">
            Start time
          </Typography>
          <DatePicker
            selected={startTime}
            onChange={(date) => setStartTime(date)}
            showTimeSelect
            dateFormat="MMMM d, yyyy h:mm aa"
            placeholderText="Start time"
            required
          />
          <Typography className="-mb-2" variant="h6">
            End time
          </Typography>
          <DatePicker
            selected={endTime}
            onChange={(date) => setEndTime(date)}
            showTimeSelect
            dateFormat="MMMM d, yyyy h:mm aa"
            placeholderText="End time"
            required
          />

          <Typography className="-mb-2" variant="h6">
            Game Category
          </Typography>
          <Select
            className="basic-single"
            classNamePrefix="select"
            isClearable={true}
            isSearchable={true}
            name="gameCategory"
            options={gameCategories}
            onChange={(choice: any) => setCategoryId(choice.value)}
          />

          <Typography className="-mb-2" variant="h6">
            Required skills
          </Typography>
          <Select
            isMulti
            options={requiredSkillOptions}
            onChange={(choices) => {
              if (choices.length > 0) {
                const skillIds = choices.map((option: GameCategoryDetail) => {
                  return option.id;
                });

                setRequiredSkills(skillIds);
              }
            }}
          />
        </CardBody>
        <CardFooter className="pt-0">
          <Button variant="gradient" onClick={handleCreateGame} fullWidth>
            Create
          </Button>
        </CardFooter>
      </Card>
    </Dialog>
  );
};

export default GameDialogForm;
