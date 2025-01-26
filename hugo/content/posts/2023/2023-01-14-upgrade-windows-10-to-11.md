---
date: "2023-01-14T19:00:00Z"
tags:
- windows
title: Upgrade Windows 10 to Windows 11
---

I used to use Windows 11, but for some reasons, the OS stopped working and I needed to clean-install it from Windows 10 from windows recovery environment.

After installing and trying to install Windows 11, there are a few things I needed to do.

Note that my motherboard was AMD B450.


# Enable UEFI boot on BIOS

In order to enable to UEFI Boot on BIOS, I needed to do

1. Convert the disk from Master Boot Record, MBR, into GUID Partition Table, GPT, partition style.
1. Disable Compatible Support Module, CSM on BIOS to enable UEFI boot

Note that I didn't have to do at the end, but I tried to enable Secure Boot on BIOS, but it prevented to boot Windows and needed to reset the bios configuration by taking its battery, described in [this article](https://appuals.com/windows-11-does-not-start-after-enabling-secure-boot/).

## Convert the MBR to GPT

See following documents for how to
* [an article](https://www.windowscentral.com/how-convert-mbr-disk-gpt-move-bios-uefi-windows-10)
* [Microsoft official document](https://learn.microsoft.com/en-us/windows-server/storage/disk-management/change-an-mbr-disk-into-a-gpt-disk)

### Confirm if my disk is MBR or GPT

1. Start a Disk Management tool
1. Right click on your disk, click **Properties > Volumes**
1. See if "Partition Style" is MBR or GPT

If it's GPT, the partition of your disk is already GPT.

##  Convert the MBR to GPT

1. Open **Settings > Update & Security > Recovery** and click "Restart Now" under "Advanced Setup"
1. Click **Troubleshoot > Advanced Options > Command Prompt**
1. After restarting the computer and start a command prompt, type `mbr2gpt /validate`
1. If succeeded, run `mbr2gpt /convert`

When I ran `mbr2gpt /convert`, I got next error for some reasons
```
MBR2GPT: Conversion completed successfully
Call WinReReapir to repair WinRE
MBR2GPT: Failed to update ReAgent.xml, please try to manually disable and enable WinRE
```

In order to manually disable and enable, I needed to run these commands, but they are not possible to run on the windows recovery environment.

```bat
reagentc /info
reagentc /disable
reagentc /enable
```

So, at first, I needed to
1. Restart a computer
1. Disable CMS on BIOS
    - It turned out this step was unnecessary
1. Start a windows and Powershell
1. Run the above command


# Enable TPM 2.0

After enabling UEFI and reran PC Health check, it showed the next message

```
TPM 2.0 must be supported and enabled on this PR
TPM: TPM not detected
```

By following [this document](https://support.microsoft.com/en-us/windows/enable-tpm-2-0-on-your-pc-1fd5a332-360d-4f46-a1e7-ae6b0c90645c),

1. Open **Settings > Update & Security > Windows Security > Device Security**, and see if there is a section for a "Security Processor". I didn't find it, so TPM is disabled on my PC.
1. Click **Settings > Update & Security > Recovery > Restart now**.
1. Click **Troubleshoot > Advanced options > UEFI Firmware Settings > Restart**
1. Go to **Peripherals** tab and enable **AMD CPU fTPM** option

# Install Windows 11 installation assisstant manually

Afterward, PC health check app shows "This PC meets Windows 11 requirements" finally.
But when I opened **Settings > Update & Security > Windows Update**, it still shows "This PC doesn't currently meet the minimum system requirements to run Windows 11".

It seems it's because Windows doesn't detect the current hardware statuses.
So following [this answer](https://answers.microsoft.com/en-us/windows/forum/all/clear-false-this-pc-doesnt-currently-meet-the/edf9e698-1549-4d57-b955-41d63976e04d), download windows 11 manually from [here](https://www.microsoft.com/en-us/software-download/windows11)

# Fix Ethernet connection issue after installing Windows 11

After installing Windows 11, I got a `networking connection error` saying `Default gateway isn't available` repeatedly, and this kept happening even after it was fixed by networking diagnostics.

In order to fix them, following [the answer](https://answers.microsoft.com/en-us/windows/forum/all/default-gateway-repeatedly-isnt-available/d7bdf938-fbf6-460d-918c-00b4ebd5da23) on this document.

1. Start a command prompt and run following commands. Note that there was no command `netsh int reset all`, so I didn't include it.

    ```bat
    ipconfig /flushdns
    nbtstat -R
    nbtstat -RR
    netsh int ip reset
    netsh winsock reset
    ```

1. Open **Settings > Network & Internet** and see "Description" of the Ethernet and see the network adapter name.
1. "Check Run a device manager and uninstall the network adapter shown on the above.
1. Restarted my computer
1. After rerunning the network diagnostics tool for a networking error, it was fixed

But this didn't work.

I also tried

```
ipconfig /release
ipconfig /renew
```

But it caused an error, which is the timeout to connect to the DHCP server.

I also tried
* to disable either an IPv4 or IPv6 network adapter
* to reset network

But I couldn't fix it.
At the end, I decided to use WiFi instead of Ethernet.
