import {
  Button,
  Dialog,
  Typography,
  Input,
  Card,
  CardBody,
  CardFooter,
} from "@material-tailwind/react";
import { useState } from "react";
import { AuthService } from "../../services/auth/auth.service";
import { LoginResponse } from "../../services/auth/auth.type";
import { toast } from "react-toastify";

type LoginFormDialogProps = {
  open: boolean;
  handleOpen: () => void;
  setLoginRes: React.Dispatch<React.SetStateAction<any>>;
};

const LoginFormDialog: React.FC<LoginFormDialogProps> = ({
  open,
  handleOpen,
  setLoginRes,
}) => {
  const [identifier, setIdentifier] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleLogin = async () => {
    try {
      const data: LoginResponse = await AuthService.login({
        identifier,
        password,
      });

      if (data && data.token) {
        toast.success("Login success");

        setIdentifier("");
        setPassword("");

        localStorage.setItem("token", data.token);
        const userData = {
          username: data.username,
          avatarUrl: data.avatarUrl,
        };

        localStorage.setItem("userData", JSON.stringify(userData));
        setLoginRes(userData);

        handleOpen();
      }

      if (data && data.error) {
        handleOpen();
        toast.error(data.error);
      }
    } catch (error) {
      toast.error("Something went wrong!");
    }
  };

  return (
    <Dialog
      size="xs"
      open={open}
      handler={handleOpen}
      className="bg-transparent shadow-none"
    >
      <Card className="mx-auto w-full max-w-[24rem]">
        <CardBody className="flex flex-col gap-4">
          <Typography variant="h4" color="blue-gray">
            Sign In
          </Typography>
          <Typography
            className="mb-3 font-normal"
            variant="paragraph"
            color="gray"
          >
            Enter your identifier and password to Sign In.
          </Typography>
          <Typography className="-mb-2" variant="h6">
            Your Email or Username
          </Typography>
          <Input
            value={identifier}
            onChange={(e) => setIdentifier(e.target.value)}
            crossOrigin={null}
            label="Email or Username"
            size="lg"
          />
          <Typography className="-mb-2" variant="h6">
            Your Password
          </Typography>
          <Input
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            crossOrigin={null}
            label="Password"
            size="lg"
            type="password"
          />
        </CardBody>
        <CardFooter className="pt-0">
          <Button
            loading={isLoading}
            variant="gradient"
            onClick={handleLogin}
            fullWidth
          >
            Sign In
          </Button>
          <Typography variant="small" className="flex justify-center mt-4">
            Don&apos;t have an account?
            <Typography
              as="a"
              href="#signup"
              variant="small"
              color="blue-gray"
              className="ml-1 font-bold"
              onClick={handleOpen}
            >
              Sign up
            </Typography>
          </Typography>
        </CardFooter>
      </Card>
    </Dialog>
  );
};

export default LoginFormDialog;
