import React from "react";
import {
  Avatar,
  Button,
  Menu,
  MenuHandler,
  MenuItem,
  MenuList,
  Typography,
} from "@material-tailwind/react";

const profileMenuItems = [
  {
    label: "Sign Out",
  },
];

type AvatarWithUserDropdownProps = {
  avatarUrl: string;
  username: string;
  setLoginRes: React.Dispatch<React.SetStateAction<any>>;
};

export function AvatarWithUserDropdown({
  avatarUrl,
  username,
  setLoginRes,
}: AvatarWithUserDropdownProps) {
  const [isMenuOpen, setIsMenuOpen] = React.useState(false);

  const closeMenu = () => setIsMenuOpen(false);

  return (
    <Menu open={isMenuOpen} handler={setIsMenuOpen} placement="bottom-end">
      <MenuHandler>
        <Button
          variant="text"
          color="blue-gray"
          className="flex items-center p-0 rounded-full"
        >
          <Avatar
            variant="circular"
            size="md"
            alt={username}
            withBorder={true}
            color="blue-gray"
            className=" p-0.5"
            src={avatarUrl}
          />
        </Button>
      </MenuHandler>
      <MenuList className="p-1">
        <MenuItem
          onClick={() => {
            closeMenu();
            setLoginRes(null);

            localStorage.removeItem("token")
            localStorage.removeItem("userData")
          }}
          className={`flex items-center gap-2 rounded `}
        >
          <Typography
            as="span"
            variant="small"
            className="font-normal"
            color={"red"}
          >
            Sign out
          </Typography>
        </MenuItem>
      </MenuList>
    </Menu>
  );
}
