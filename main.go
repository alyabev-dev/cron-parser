package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	cronString := flag.String("cron", "", "Cron string to parse (e.g., '*/15 0 1,15 * 1-5')")
	command := flag.String("cmd", "", "Command to execute (e.g., '/usr/bin/find')")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *cronString == "" || *command == "" {
		fmt.Println("Error: -cron and -cmd flags are required")
		return
	}

	fields := strings.Fields(*cronString)
	if len(fields) != 5 {
		fmt.Println("Error: invalid cron format - expected 5 fields")
		os.Exit(1)
	}

	fields = append(fields, *command)

	parser := &Parser{}
	for i, f := range fields {
		if i >= CronFieldsAmount {
			break
		}
		field := parser.Parse(f, fieldLimits[i].Min, fieldLimits[i].Max, fieldLimits[i].Title)
		parser.Print(field)
	}
	fmt.Printf("%-14s %s\n", "command", fields[5])
}

func printHelp() {
	fmt.Println("Usage: cron_parser -cron=\"<cron_string>\" -cmd=\"<command>\"")
	fmt.Println("Example: cron_parser -cron=\"*/15 0 1,15 * 1-5\" -cmd=\"/usr/bin/find\"")
	fmt.Println("Flags:")
	fmt.Println("-cron string (cron string to parse (e.g., '*/15 0 1,15 * 1-5'))")
	fmt.Println("-cmd string (command to execute (e.g., '/usr/bin/find'))")
	fmt.Println("-help (show this help message)")
}
