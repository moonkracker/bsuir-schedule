# BSUIR-SCHEDULE

---

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e66106cfb7234eda8b4fb516dc872c55)](https://www.codacy.com/gh/moonkracker/bsuir-schedule/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=moonkracker/bsuir-schedule&amp;utm_campaign=Badge_Grade) [![Go](https://github.com/moonkracker/bsuir-schedule/actions/workflows/bsuir-schedule.yml/badge.svg)](https://github.com/moonkracker/bsuir-schedule/actions/workflows/bsuir-schedule.yml) [![Coverage Status](https://coveralls.io/repos/github/moonkracker/bsuir-schedule/badge.svg?branch=master)](https://coveralls.io/github/moonkracker/bsuir-schedule?branch=master) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=moonkracker_bsuir-schedule&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=moonkracker_bsuir-schedule)

This application is a simple CLI application that allows you to get the schedule of the BSUIR.

## Installation

```bash
brew install moonkracker/tap/bsuir-schedule
```

## Usage

```go
Get BSUIR schedule

Usage:
  bsuir-schedule [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  group-schedule    Get group schedule
  help              Help about any command
  teacher-schedule  Get teacher schedule
  version           Print the version number of bsuir-schedule

Flags:
  -h, --help      help for bsuir-schedule
  -v, --version   version for bsuir-schedule

Use "bsuir-schedule [command] --help" for more information about a command.
```

## Docker

```bash
docker run --rm -it ghcr.io/moonkracker/bsuir-schedule:latest
```
