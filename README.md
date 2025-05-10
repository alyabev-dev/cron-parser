# Cron Parser

Cron Parser is a CLI tool for parsing cron strings.

## Usage

1. Build:
    - go build -o cron_parser
2. Use as CLI (optionally):
    - sudo cp cron_parser /usr/local/bin/
3. Run:
    - cron_parser -cron="<cron_string>" -cmd="<command>" || ./cron_parser -cron="<cron_string>" -cmd="<command>"

## Example

cron_parser -cron="*/15 0 1,15 * 1-5" -cmd="/usr/bin/find"