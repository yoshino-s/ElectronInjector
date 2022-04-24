# Electron Injector
[![CodeQL](https://github.com/Yoshino-s/ElectronInjector/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Yoshino-s/ElectronInjector/actions/workflows/codeql-analysis.yml)


Another injector on encrypted electron app.

## Warning

> Only for study and research. DO NOT USE THIS TOOL FOR CRACKING.

## Feature

- Supports Some Software Version 1.0.0+ (At least for now.)
- Support Some Software Windows And Linux
- Verified On Some Software
  - Linux: <!-- Linux-Begin -->1.0.0, 1.0.3, 1.1.3, 1.1.4, 1.1.5, 1.2.3, 1.2.4<!-- Linux-End --><!-- Linux-End --><!-- Linux-End -->
  - Windows: <!-- Windows-Begin -->1.1.4, 1.1.5, 1.2.3, 1.2.4<!-- Windows-End --><!-- Windows-End --><!-- Windows-End --><!-- Windows-End --><!-- Windows-End -->

## Usage

```sh
pip install -r requirements.txt
python3 main.py --help
```

```
Usage: main.py [OPTIONS] [PROG]

Options:
  -I, --inject PATH  Inject file
  -B, --from-bak     Use bak file as raw program
  --help             Show this message and exit.
```

## Example

```sh
python3 main.py .../resources/app.asar.unpacked/main.node
```

## How does it work?

Read this: https://blog.yoshino-s.online/2022/03/31/electron-enc-crack/
