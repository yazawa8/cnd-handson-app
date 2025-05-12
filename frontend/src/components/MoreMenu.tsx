// src/components/MoreMenu.tsx
import { useState, MouseEvent } from 'react';
import { IconButton, Menu, MenuItem } from '@mui/material';
import MoreVertIcon from '@mui/icons-material/MoreVert';

export interface MoreMenuOption<T> {
  label: string;
  onClick: (id: T) => void;
}

interface MoreMenuProps<T> {
  id: T;
  options: MoreMenuOption<T>[];
  // 省略可: ボタンの aria-label
  ariaLabel?: string;
}

function MoreMenu<T extends string | number>({ id, options, ariaLabel = 'settings' }: MoreMenuProps<T>) {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleOpen = (e: MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    setAnchorEl(e.currentTarget);
  };
  const handleClose = () => setAnchorEl(null);

  const handleClick = (fn: (id: T) => void) => {
    fn(id);
    handleClose();
  };

  return (
    <>
      <IconButton edge="end" aria-label={ariaLabel} onClick={handleOpen} size="small">
        <MoreVertIcon fontSize="small" />
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        {options.map((opt, i) => (
          <MenuItem key={i} onClick={() => handleClick(opt.onClick)}>
            {opt.label}
          </MenuItem>
        ))}
      </Menu>
    </>
  );
}

export default MoreMenu;
