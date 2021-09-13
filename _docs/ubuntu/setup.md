---
title: Setup a Ubuntu desktop computer
date: 2021-09-12 00:00:00
---

I set up Ubuntu for the 1st time since several years ago.

Set up basic configuration
===

Settings
---
Open **Show Applications > Settings** and set each configuration.

### Displays
#### Set a resolution
I ended up to set 1920x1080.
- At first, I set 3920x2160 but, there was a performance issue with my graphics card.
- I reduced the resolution to 2880x1620 (16:9), but there was a garbled issue on many applications.

#### Enable Night light mode
* Open **Night Light** tab
* Enable **Night Light** toggle

### Mouse & Touchpad
Change the scroll direction on the middle button of a mouse by turning on **Mouse > Natural Scrolling**.

### Keyboard
Before setting a keyboard, install [Gnome tweaks](https://wiki.gnome.org/action/show/Apps/Tweaks?action=show&redirect=Apps%2FGnomeTweakTool) by next command.
```
sudo apt install gnome-tweaks
```

#### Swap Caps Lock key with a Ctrl key.
* Run `gnome-tweaks`
* Open **Keyboard & Mouse > Additional Layout Options > Caps Lock behavior** and choose **Caps Lock is also a Ctrl**.

See [this page](https://askubuntu.com/questions/33774/how-do-i-remap-the-caps-lock-and-ctrl-keys) for more details.


### Region & Language
In order install a language other than English, go to **Manage Installed Languages > Install/Remove Languages** and choose a language you wanna install.
Then restart a computer


Themes
---

### Change themes based on time
From [this article](https://www.linuxuprising.com/2019/11/change-gtk-theme-to-dark-variant-when.html).
First of all, GNOME shell extension and "chrome-gnome-shell" need to be installed, following [this document](https://itsfoss.com/gnome-shell-extensions/).
Then run `gnome-tweaks` and see there is **Extensions > User themes**.
If not, restart a computer and check it again.
Then install a browser extension for gnome extension and install [Night Theme Switcher](https://extensions.gnome.org/extension/2236/night-theme-switcher/).


Install applications
===

Install development tools
---
Run next commands to install git, emacs, and tmux
```
sudo apt git emacs tmux fish
```

Install 1Password GUI and CLI
---

Follow official documents to install them
* [GUI](https://support.1password.com/install-linux/)
* [CLI](https://1password.com/downloads/command-line/)

To set up CLI, run the next command at first
```
> op signin my.1password.com example@example.com
```

YouTube
---
In order to play live streaming videos on Firefox, we need to install "ffmpeg" by `sudo apt install ffmpeg` and restart Firefox.
I found this on [this page](https://askubuntu.com/questions/1298248/cant-play-live-youtube-videos-on-ubuntu-20-04).

Google Drive
---

Set up Google account on online accounts at first.

Then run next commands to install, make a mounted directory, and mount a Google Drive to the directory.

```
sudo add-apt-repository ppa:alessandro-strada/ppa
sudo apt update && sudo apt install google-drive-ocamlfuse
google-drive-ocamlfuse
mkdir ~/GoogleDrive
google-drive-ocamlfuse ~/GoogleDrive
```
This is from [this page](https://linuxhint.com/google_drive_installation_ubuntu/).

However, this mount is only active until I shutdown a computer.
In order to keep mounting even after shutdown, the command `google-drive-ocamlfuse /home/username/GoogleDrive` to be added on **Startup Applications**.
[This page](https://www.fosslinux.com/500/how-to-add-auto-startup-applications-in-ubuntu-16-04.htm) describes about Startup Applications on Ubuntu.


Troubleshotings
===

Solve garbled fonts on a console
---
According to [this question](https://askubuntu.com/questions/72023/why-are-letters-overlapping-in-the-terminal), "ttf-ubuntu-font-family" is required.
Install it by `sudo apt ttf-ubuntu-font-family`.

.deb file cannot be installed on Nautilus or Firefox
---

Install **gdebi** by next command and open **.deb** file using **GDebi Package Instaler**.
```
sudo apt install gdebi
```

I followed [this question and answers](https://askubuntu.com/questions/1232868/problem-installing-deb-in-software-install-ubuntu-20-04).


TODOs
===

- [ ] On tmux, copy the text to a system clipboard
- [ ] Set up a copy text history
