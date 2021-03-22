---
title: Notes when I updated a WSL version from 1 to 2
date: 2021-03-17T00:00:00Z
---

I mostly followed [this article](https://docs.microsoft.com/en-us/windows/wsl/install-win10#set-your-distribution-version-to-wsl-1-or-wsl-2) to update a WSL version, except that I didn't enable Hyper-V until then and got an error `Please enable the Virtual Machine Platform Windows feature and ensure virtualization is enabled in the BIOS.` while I was trying to update the vesrion.

These are steps I did to update the WSL version.

1. Ran PowerShell as Administrator.
1. Enabled the Windows Subsystem for Linux by next command.
    ```
    PS C:\Windows\system32> dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
    Deployment Image Servicing and Management tool
    Version: 10.0.19041.844

    Image Version: 10.0.19042.868

    Enabling feature(s)
    [==========================100.0%==========================]
    The operation completed successfully.
    ```

1. Enabled Virtual Machine Feature by next command.
    ```
    PS C:\Windows\system32> dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart

    Deployment Image Servicing and Management tool
    Version: 10.0.19041.844

    Image Version: 10.0.19042.868

    Enabling feature(s)
    [==========================100.0%==========================]
    The operation completed successfully.
    ```

1. Tried to set the default vesrion of WSL to 2 and got an error.
    ```
    PS C:\Windows\system32> wsl --set-default-version 2
    Please enable the Virtual Machine Platform Windows feature and ensure virtualization is enabled in the BIOS.
    For information please visit https://aka.ms/wsl2-install
    ```

1. Checked if Hypber-V was enabled on my machine.

   In order to use WSL 2 on Windows 10, a CPU has to support Hyper-V.
   We can check whether CPU enables it or not on the result of `msinfo32`, which I found on [this article](https://www.zdnet.com/article/windows-10-tip-find-out-if-your-pc-can-run-hyper-v/).
   My result was

    - Hyper-V - VM Monitor Mode Extensions: Yes
    - Hyper-V - Second Level Address Translation Extensions: Yes
    - Hyper-V - Virtualization Enabled in Firmware: **No**
    - Hyper-V - Data Execution Protection: Yes

1. Enabled Hyper-V on a BIOS setting.

    Because my CPU was Ryzen 5, I followed [this comment](https://superuser.com/a/1248572).
    So, I restarted my computer, went to **Advanced Frequency Settings > Advanced CPU Settings** and enabled on **SVM mode**.
    Then restarted my computer again.

1. Set default version of WSL to 2.

    Now it worked to me.
    ```
    PS C:\Windows\system32> wsl --set-default-version 2
    For information on key differences with WSL 2 please visit https://aka.ms/wsl2
    ```

1. Updated existing distribution to 2.

    I also updated Ubuntu which I installed before.
    We can check the version of installed distributions by `wsl -l -v` and can also update them by `wsl --set-version [Name] [Version]`.

    ```
    PS C:\Windows\system32> wsl -l -v
      NAME            STATE           VERSION
      * Ubuntu          Running         1
        Ubuntu-20.04    Stopped         1
    PS C:\Windows\system32> wsl --set-version Ubuntu 2
    Conversion in progress, this may take a few minutes...
    For information on key differences with WSL 2 please visit https://aka.ms/wsl2
    Conversion complete.
    PS C:\Windows\system32> wsl -l -v
      NAME            STATE           VERSION
      * Ubuntu          Stopped         2
        Ubuntu-20.04    Stopped         1
    ```

Edited
===

On March 22, 2021 - Very slow network issue
---

I had an issue of very slow network on WSL 2 and I had to wait for a long time while I was using apt or git.
After seeing [this comment](https://github.com/microsoft/WSL/issues/4901#issuecomment-748531438), I was able to solve the issue by disabling "Large Send Offload".
I did it on `Device Manager > Network adapters > Hyper-V Virtual Ethernet Adapter > Properties > Advanced`.
