import type { CSSProperties, FC, ReactNode } from "react";

type ButtonProps = {
  children: ReactNode;
  onClick?: () => void;
  variant?: "primary" | "secondary" | "danger";
  disabled?: boolean;
  type?: "button" | "submit" | "reset";
};

const Button: FC<ButtonProps> = ({
  children,
  onClick,
  variant = "primary",
  disabled = false,
  type = "button",
}) => {
  const getButtonStyle = (): CSSProperties => {
    const baseStyle: CSSProperties = {
      padding: "8px 16px",
      borderRadius: "4px",
      border: "none",
      cursor: disabled ? "not-allowed" : "pointer",
      opacity: disabled ? 0.7 : 1,
    };

    switch (variant) {
      case "primary":
        return {
          ...baseStyle,
          backgroundColor: "#007bff",
          color: "white",
        };
      case "secondary":
        return {
          ...baseStyle,
          backgroundColor: "#6c757d",
          color: "white",
        };
      case "danger":
        return {
          ...baseStyle,
          backgroundColor: "#dc3545",
          color: "white",
        };
      default:
        return baseStyle;
    }
  };

  return (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled}
      style={getButtonStyle()}
    >
      {children}
    </button>
  );
};

export default Button;
