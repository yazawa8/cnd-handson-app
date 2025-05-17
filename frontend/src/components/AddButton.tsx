import type React from "react";
import { Button, type ButtonProps } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";

interface AddButtonProps extends Omit<ButtonProps, "onClick"> {
  onClick: () => void;
  label?: string;
}

const AddButton: React.FC<AddButtonProps> = ({
  onClick,
  label = "追加",
  variant = "contained",
  color = "primary",
  size = "medium",
  ...rest
}) => {
  return (
    <Button
      startIcon={<AddIcon />}
      variant={variant}
      color={color}
      size={size}
      onClick={onClick}
      {...rest}
    >
      {label}
    </Button>
  );
};

export default AddButton;
