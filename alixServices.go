package main

import (
	"context"
	"fmt"
	"os"
	// "os/exec"
	"github.com/urfave/cli/v3"
)
func debug(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("testing123", cmd.Args().Get(0), cmd.Args().Get(1))
	return nil
}

func list(ctx context.Context, cmd *cli.Command) error {
	// Translate this to golang "awk -F'[ =]' '/^alias / {print $2}' ~/.bashrc | sort""
	logger.Println("Successfully executed list function with the following parameters: ")
	return nil
}
func listDetailed(ctx context.Context, cmd *cli.Command) error {
	// Translate this to golang "grep "^alias " ~/.bashrc | sed 's/^alias //' | sort -t= -k2 | sed 's/^\([^=]*\)=/\x1b[1;32m\1\x1b[0m=/'"
	logger.Println("Successfully executed detailed list function with the following parameters: ")
	return nil
}
func add(ctx context.Context, cmd *cli.Command) error {
	alias_name := cmd.Args().Get(0)
	alias_cmd := cmd.Args().Get(1)
	/* Translate this to golang "
		# Check if the alias already exists
		if grep -q "^alias $alias_name=" ~/.bashrc; then
			echo "Alias '$alias_name' already exists. Use update instead."
		else
			# Add the new alias
			echo "alias $alias_name=\"$alias_cmd\"" >> ~/.bashrc
			echo "Alias '$alias_name' added as '$alias_cmd'."
		fi

		# Apply changes
		source ~/.bashrc
	"
	*/
	logger.Println("Successfully executed add function with the following parameters: ")
	return nil
}
func delete(ctx context.Context, cmd *cli.Command) error {
	alias_name := cmd.Args().Get(0)
	/* Translate this to golang "
		#!/bin/bash
		# Ask for the alias name to delete
		read -p "Enter alias name to delete: " alias_name

		# Check if the alias exists in ~/.bashrc
		if grep -q "^alias $alias_name=" ~/.bashrc; then
			# Delete the alias line
			sed -i "/^alias $alias_name=/d" ~/.bashrc
			echo "Alias '$alias_name' deleted."
		else
			echo "Alias '$alias_name' not found in ~/.bashrc."
		fi

		# Apply changes
		source ~/.bashrc
	"*/ 
	logger.Println("Successfully executed add function with the following parameters: ")
	return nil
}
func update(ctx context.Context, cmd *cli.Command) error {
	alias_name := cmd.Args().Get(0)
	alias_cmd := cmd.Args().Get(1)
	/* Translate this to golang "
		# Check if alias exists in ~/.bashrc
		if grep -q "^alias $alias_name=" ~/.bashrc; then
			# Update existing alias
			sed -i "s|^alias $alias_name=.*|alias $alias_name=\"$alias_cmd\"|" ~/.bashrc
			echo "Alias '$alias_name' updated to '$alias_cmd'."
		else
			# Add new alias
			echo "alias $alias_name=\"$alias_cmd\"" >> ~/.bashrc
			echo "Alias '$alias_name' added as '$alias_cmd'."
		fi

		# Apply changes
		source ~/.bashrc
	"*/ 
	logger.Println("Successfully executed update function with the following parameters: ")
	return nil
}

func detectShellType() (string,error) {
	home , err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
    bashrc := home + "/.bashrc"
    zshrc := home + "/.zshrc"
    if _, err := os.Stat(bashrc); err == nil {
        return bashrc, nil
    }
    if _, err := os.Stat(zshrc); err == nil {
        return zshrc, nil
    }
    return "", fmt.Errorf("no supported shell config found")
}