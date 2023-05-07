# gousbip

This tool is for listing available USB devices that can be used in the WSL (Windows Subsystem for Linux) and allowing the user to select a device to connect or disconnect.

## Requirements

Please refer to the following for more information.

https://learn.microsoft.com/ja-jp/windows/wsl/connect-usb

## Install

```bash
go install github.com/kinpoko/gousbip@latest
```

## Usage

### Connecting a Device

To connect a device, use the `-a` flag.

```bash
gousbip -a
```

This command displays a list of available devices. The user can select a device from the list using fzf. Once the selection is complete, the selected device is connected to WSL.

### Disconnecting a Device

To disconnect a device, use the -d flag.

```bash
gousbip -d
```

This command displays a list of available devices. The user can select a device from the list using fzf. Once the selection is complete, the selected device is disconnected.
