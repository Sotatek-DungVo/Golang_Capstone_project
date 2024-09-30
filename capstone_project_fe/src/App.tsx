import { useState } from "react";
import "./App.css";
import { IoGameController } from "react-icons/io5";
import { Button } from "@material-tailwind/react";
import LoginFormDialog from "./components/auth/LoginFormDialog";
import Sidebar from "./components/Sidebar";
import { Outlet, useNavigate } from "react-router-dom";
import { AvatarWithUserDropdown } from "./components/AvatarWithUserDropdown";
import { LoginResponse } from "./services/auth/auth.type";

function App() {
  const [openForm, setOpenForm] = useState<boolean>(false);
  const handleOpen = () => setOpenForm(!openForm);
  const navigate = useNavigate();
  const [loginRes, setLoginRes] = useState<LoginResponse | null>(null);

  return (
    <div className="flex flex-col min-h-screen">
      {/* Header */}
      <div className="flex items-center justify-between px-4 py-2">
        <div
          className="flex items-center cursor-pointer"
          onClick={() => navigate("/")}
        >
          <div className="p-1 bg-gray-400 rounded-full">
            <IoGameController className="w-8 h-8" />
          </div>
        </div>
        <div className="flex items-center space-x-4">
          {!loginRes && (
            <Button onClick={handleOpen} variant="gradient">
              Login
            </Button>
          )}

          {loginRes && (
            <AvatarWithUserDropdown
              avatarUrl={loginRes.avatarUrl}
              username={loginRes.username}
              setLoginRes={setLoginRes}
            />
          )}
        </div>
      </div>

      {/* Main Content */}
      <div className="flex flex-1">
        <Sidebar />

        <div className="flex-1 overflow-auto">
          <Outlet />
        </div>
      </div>

      <LoginFormDialog
        open={openForm}
        handleOpen={handleOpen}
        setLoginRes={setLoginRes}
      />
    </div>
  );
}

export default App;
