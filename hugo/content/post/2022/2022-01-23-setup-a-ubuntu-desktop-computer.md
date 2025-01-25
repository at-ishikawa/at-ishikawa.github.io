---
date: "2022-01-23T00:00:00Z"
tags:
- ubuntu
title: setup a ubuntu desktop computer
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


### Support Multi gestures of Touchpad

**Note that most of gestures still doesn't work.**

Folllow [this article](https://www.omgubuntu.co.uk/2018/09/linux-touchpad-gestures-app) to set up this.
Install [Gestures](https://gitlab.com/cunidev/gestures) and its dependencies.

```
> sudo apt install python3 python3-gi meson xdotool libinput-tools gettext
```
Then install [libinput-gestures](https://github.com/bulletmark/libinput-gestures).

```
sudo gpasswd -a $USER input
ghq get github.com/bulletmark/libinput-gestures
cd /path/to/libinput-gestures
sudo make install
libinput-gestures-setup autostart start
```

Next, install a flatpak to download Gestures.
```
sudo apt install flatpak gnome-software-plugin-flatpak
flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
```
Before next step, restart your computer for the flatpak setup.

Then download Gestures from FLATHUB.


#### Troubleshooting
##### Error `No module named 'packaging`

I got an error `No module named 'packaging` when I ran `libinput-gestures-setup autostart start` for some reasons.
Following this [stackdriver question](https://stackoverflow.com/questions/42222096/no-module-named-packaging), I ran `sudo apt install python3-packaging` and restarted my desktop, and then it started successfully.

##### Swipe left and right to go back and forward on a browser doesn't work

I haven't been able to solve it yet.
The configuration uses xdotool, but it looks it has an issue to get a window on the tool.
I couldn't even get an active window on the CLI.
```
> xdotool getactivewindow
XGetWindowProperty[_NET_ACTIVE_WINDOW] failed (code=1)
xdo_get_active_window reported an error
```

### Keyboard
Before setting a keyboard, install [Gnome tweaks](https://wiki.gnome.org/action/show/Apps/Tweaks?action=show&redirect=Apps%2FGnomeTweakTool) by next command.
```
sudo apt install gnome-tweaks
```

#### Use Emacs shortcut keys everywhere
* Run `gnome-tweaks`
* Open **Keyboard & Mouse** tab and enable **Emacs Input**
There are other ways described in [this thread](https://www.reddit.com/r/emacs/comments/c22ff1/gtk_4_support_for_key_themes_does_not_affect/)

#### Swap Caps Lock key with a Ctrl key.
* Run `gnome-tweaks`
* Open **Keyboard & Mouse > Additional Layout Options > Caps Lock behavior** and choose **Caps Lock is also a Ctrl**.

See [this page](https://askubuntu.com/questions/33774/how-do-i-remap-the-caps-lock-and-ctrl-keys) for more details.


### Region & Language
In order install a language other than English, go to **Manage Installed Languages > Install/Remove Languages** and choose a language you wanna install.
Then restart a computer


### Use clipboard history
Follow [this article](https://ubuntuhandbook.org/index.php/2021/10/access-copy-paste-history-ubuntu-gpaste/).

```
sudo apt install gnome-shell-extension-gpaste
```
Also, to manage extensions, install next tool if you haven't installed it yet.

```
sudo apt install gnome-shell-extension-prefs
```

Once you restart your GNOME shell, you can use them.

There are shortcut keys for this tool to use this clipboard more efficiently.
See the full list in [this article](https://www.linuxuprising.com/2018/08/gpaste-is-great-clipboard-manager-for.html)

* Ctrl + Alt + H: Open the GPaste history on the top bar
* Ctrl + Alt + S: Mask the current item as a password

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

### Install gh
Note that it's not recommended to install it via snap according to [this comment](https://github.com/cli/cli/issues/3185#issuecomment-797596234)
I also got the issue on the above.
So, follow an [official document](https://github.com/cli/cli/blob/trunk/docs/install_linux.md) to install this cli.
```
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo gpg --dearmor -o /usr/share/keyrings/githubcli-archive-keyring.gpg
echo deb [arch=(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
sudo apt update
sudo apt install gh
```

### Install docker

First, install docker. I forgot how I did, so please check the Internet.
Then in order to run docker commands without `sudo`, add your user into docker group, by following [this article](https://www.cloudsavvyit.com/10623/how-to-install-docker-and-docker-compose-on-linux/)

```
sudo usermod -aG docker $USER
```

Then restart your machine.


Install 1Password GUI and CLI
---

Follow official documents to install them
* [GUI](https://support.1password.com/install-linux/)
* [CLI](https://1password.com/downloads/command-line/)

To set up CLI, run the next command at first
```
> op signin my.1password.com example@example.com
```

ULauncher
---
I installed [ULauncher](https://ulauncher.io/) as a launcher in Ubuntu.
To download it, run next command.
```
sudo add-apt-repository ppa:agornostal/ulauncher && sudo apt update && sudo apt install ulauncher
```

But it didn't work as expected.
In order to make it work, I have to update a few things by following [this document](https://github.com/Ulauncher/Ulauncher/wiki/Hotkey-In-Wayland)


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



Grub
===

- In order to set the default boot option as the last OS you chose, edit `/etc/default/grub` and set `GRUB_DEFAULT=saved` and `GRUB_SAVEDEFAULT=true`.
    - See [this article](https://www.howtogeek.com/196655/how-to-configure-the-grub2-boot-loaders-settings/) for more details.
- In order to fix grub boot errors, [boot-repair](https://help.ubuntu.com/community/Boot-Repair) may be helpful to install and use it, run

    ```
    sudo add-apt-repository ppa:yannubuntu/boot-repair && sudo apt update
    sudo apt install -y boot-repair && boot-repair
    ```


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
