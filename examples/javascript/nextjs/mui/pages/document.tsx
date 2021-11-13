import type {NextPage} from 'next'
import styles from '../styles/Home.module.css'
import {Popover, Typography} from "@mui/material";
import React from "react";

const Home: NextPage = () => {
  const [anchorEl, setAnchorEl] = React.useState(null);

  const handlePopoverOpen: React.MouseEventHandler = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl((event.currentTarget as any));
  };

  const handlePopoverClose = () => {
    setAnchorEl(null);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const open = Boolean(anchorEl);
  const id = 'mouse-over-popover';

  return (
    <div className={styles.container}>
      <main className={styles.main}>
        <Typography
          aria-owns={id}
          aria-haspopup="true"
          onMouseEnter={handlePopoverOpen}
          onMouseLeave={handlePopoverClose}
        >
          Hover with Popover
        </Typography>
        <Popover
          id="mouse-over-popover"
          sx={{
            pointerEvents: 'none',
          }}
          open={open}
          anchorEl={anchorEl}
          onClose={handlePopoverClose}
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'left',
          }}
          transformOrigin={{
            vertical: 'top',
            horizontal: 'left',
          }}
          disableRestoreFocus
        >
          The content of the Popover.
        </Popover>
      </main>
    </div>
  )
}

export default Home
