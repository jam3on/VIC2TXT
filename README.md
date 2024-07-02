# VIC2TXT

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.MIT)
[![Language: Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#)
[![x64](https://img.shields.io/badge/Windows-64_bit-0078d7.svg)](#)
[![v1.0](https://img.shields.io/badge/Version-1.0-ff5733.svg)](#)

Simple tool for converting the vic database from json format to txt format

## Usage

To use this tool from the command line, run the following command:

```sh
./vic2txt.exe -input(-i) %userprofile%/Desktop/database.json(filepath) -output(-o) %userprofile%/Desktop//Desktop/(directory)
```

Replace `database.json` with the appropriate values.

## Creating an Executable Binary

To create an executable binary, you will need to install `golang`. 


```sh
go-winres make
go build -ldflags "-s -w" --trimpath -o vic2txt.exe
```

This command will generate a standalone executable file for Decrypt-FVEK.