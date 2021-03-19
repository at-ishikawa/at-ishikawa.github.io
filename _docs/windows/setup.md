---
title: Setup a windows desktop computer
---

This is written on March 2021.

I set up Windows for the 1st time since several years ago and there was a lot of difference between recent version and the version I used.
This page collects information for something it would have been helpful for setting up windows for me.

Setup for software development
===

Join Windows Insider Program
---
[Windows Insider Program](https://insider.windows.com/en-us/getting-started) provides users to access applications from their dev, beta, or preview releases.
This is helpful to use some applications before it becomes stable releases, like WSL, and many documents on the Internet is also written for these pre-stable versions.

Once you joined, you wanna make sure turnning on automatic updates for OS version.
See [this page](https://www.sheffield.ac.uk/it-services/information-security/windows-updat) to turn it on.


Windows Subsystem for Linux (WSL)
---
[WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) is to run Linux on a windows as a native Windows application, unlike running a virtual machine or dual boot.
It's really helpful for developers to develop their applications on Windows like running apt or git.
You can find some Linux distributions on Microsoft Store.

### Enable copy and paste
As a default, copy and paste isn't enabled, but you can easily do it on **Properties > Options > Edit Options > User Ctrl+Shirt+C/V as Copy/Paste**.
See [this article](https://devblogs.microsoft.com/commandline/copy-and-paste-arrives-for-linuxwsl-consoles/) for more details.

### Update Fonts to solve garbled text
The default font might not support showing your languages.
In that case, updating a font might solve your issue.
You can update it by **Properties > Font > Font**.

### Limitations or issues I haven't been solved yet

1. The directory of a home directory on WSL (/home/$USER) is different from one of Windows (C:\Users\$USER), and if I put something like a git repository under Windows directory, then permissions of it was changed to 777 and all changes are changed and hard to check each change by git.
    - I wanted to develop an application using IDE on Windows but manage it by tools like git on WSL. I couldn't do it because of this.
1. Supporting GUI of Linux is under development and I haven't tried, though there is an [article](https://medium.com/@japheth.yates/the-complete-wsl2-gui-setup-2582828f4577) to do it.

### Troubleshootings
1. How to upgrade WSL from 1 to 2?
    - I wrote [an article](/2021/03/17/update-wsl-version/) for how I did it.

winget
---
[winget](https://docs.microsoft.com/en-us/windows/package-manager/winget/) is a CLI to install and manage applications, just like HomeBrew on Mac, apt or yum on Linux.
You need to install [App Installer](https://www.microsoft.com/en-us/p/app-installer/9nblggh4nns1?activetab=pivot:overviewtab) to use this tool, and also need to join Windows Insider Program to install it.



Setup for daily use
===

UI
---

### Enable night light
You can enable a night light under **Settings > System > Display** to reduce blue light and also can set a schedule.

### Enable Dark mode
You can enable Dark mode under **Settings > Personalization > Colors**.
See [this article](https://www.pcmag.com/how-to/how-to-enable-dark-mode-in-windows-10) for more details.

#### To turn on Dark mode at night automatically
Unfortunately, there is no built-in feature to switch a light and dark mode automatically.
In order to enable Dark mode only during sunset, you need to set up something.
The way I did was using a Task Scheduler to enable Dark mode automatically at night and disable it on the morning, by following [this article](https://www.howtogeek.com/356087/how-to-automatically-enable-windows-10s-dark-theme-at-night/).
