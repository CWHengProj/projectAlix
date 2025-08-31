package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/urfave/cli/v3"
)

var logger *log.Logger
var shellDir string

func main() {
	// EXAMPLE: Append to an existing template
	cli.RootCommandHelpTemplate = fmt.Sprintf(`%s

	WEBSITE: https://github.com/CWHengProj/projectAlix

	`, cli.RootCommandHelpTemplate)
	cmd := &cli.Command{
		Name:  "alix",
		Usage: "summons alix, the manager of aliases",
		Commands: []*cli.Command{
			// {
			// 	Name:  "debug",
			// 	Usage: "Testing purposes",
			// 	Action: debug,
			// },
			{
				Name:    "add",
				Usage:   "Creates a new alias",
				Aliases: []string{"new", "create", "insert"},
				Action:  add,
			},
			{
				Name:    "delete",
				Usage:   "Remove an existing alias",
				Aliases: []string{"remove", "rm", "del"},
				Action:  delete,
			},
			{
				Name:    "update",
				Usage:   "Updates an existing alias",
				Aliases: []string{"edit", "modify", "change"},
				Action:  update,
			},
			{
				Name:   "ls",
				Usage:  "Shows all existing aliases",
				Action: list,
				Commands: []*cli.Command{
					{
						Name:    "detailed",
						Usage:   "Shows existing aliases in detail",
						Aliases: []string{"show-all", "details", "info"},
						Action:  listDetailed,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	logDir := filepath.Join(home, ".alix")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err)
	}
	logPath := filepath.Join(logDir, "alix.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Started new instance of alix")

	shellDir, err = detectShellType()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Shell type detected, path: ", shellDir)
}
