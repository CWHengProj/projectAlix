package main

import (
	"context"
	"fmt"
	"os"
	"bufio"
    "sort"
    "strings"
	"github.com/urfave/cli/v3"
)
// func debug(ctx context.Context, cmd *cli.Command) error {
// 	fmt.Println("testing123", cmd.Args().Get(0), cmd.Args().Get(1))
// 	return nil
// }

func list(ctx context.Context, cmd *cli.Command) error {
	file, err := os.Open(shellDir)
    if err != nil {
        return err
    }
    defer file.Close()

    var aliases []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "alias ") {
            // Split by space and '='
            parts := strings.FieldsFunc(line, func(r rune) bool {
                return r == ' ' || r == '='
            })
            if len(parts) > 1 {
                aliases = append(aliases, parts[1])
            }
        }
    }
    if err := scanner.Err(); err != nil {
        return err
    }

    sort.Strings(aliases)
    for _, alias := range aliases {
        fmt.Println(alias)
    }
	logger.Println("Successfully executed list function")
	return nil
}
func listDetailed(ctx context.Context, cmd *cli.Command) error {
	file, err := os.Open(shellDir)
    if err != nil {
        return err
    }
    defer file.Close()

    type aliasEntry struct {
        Name string
        Cmd  string
        Raw  string
    }
    var entries []aliasEntry

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "alias ") {
            aliasLine := strings.TrimPrefix(line, "alias ")
            parts := strings.SplitN(aliasLine, "=", 2)
            if len(parts) == 2 {
                name := parts[0]
                cmd := parts[1]
                entries = append(entries, aliasEntry{
                    Name: name,
                    Cmd:  cmd,
                    Raw:  aliasLine,
                })
            }
        }
    }
    if err := scanner.Err(); err != nil {
        return err
    }

    // Sort by command (after '=')
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Cmd < entries[j].Cmd
    })

    for _, entry := range entries {
        // Highlight alias name in green
        fmt.Printf("\x1b[1;32m%s\x1b[0m=%s\n", entry.Name, entry.Cmd)
    }

	logger.Println("Successfully executed detailed list function")
	return nil
}
func add(ctx context.Context, cmd *cli.Command) error {
	alias_name := cmd.Args().Get(0)
	alias_cmd := cmd.Args().Get(1)
	if alias_name == "" || alias_cmd == "" {
        fmt.Println("Usage: add <alias_name> <alias_cmd>")
        return nil
    }

    // Check if alias already exists
    file, err := os.Open(shellDir)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    exists := false
    search := fmt.Sprintf("alias %s=", alias_name)
    for scanner.Scan() {
        if strings.HasPrefix(scanner.Text(), search) {
            exists = true
            break
        }
    }
    if err := scanner.Err(); err != nil {
        return err
    }

    if exists {
        fmt.Printf("Alias '%s' already exists. Use update instead.\n", alias_name)
        return nil
    }

    // Add the new alias
    f, err := os.OpenFile(shellDir, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    aliasLine := fmt.Sprintf("alias %s=\"%s\"\n", alias_name, alias_cmd)
    if _, err := f.WriteString(aliasLine); err != nil {
        return err
    }

    fmt.Printf("Alias '%s' added as '%s'.\n", alias_name, alias_cmd)
    logger.Println("Successfully executed add function with the following parameters: ", alias_name, alias_cmd)
    fmt.Println("Successfully executed add function: ", alias_name, " please restart your shell to use the newly added aliases.")
    return nil
}
func delete(ctx context.Context, cmd *cli.Command) error {
	alias_name := cmd.Args().Get(0)
    if alias_name == "" {
        fmt.Println("Usage: delete <alias_name>")
        return nil
    }

    // Read all lines from shellDir
    file, err := os.Open(shellDir)
    if err != nil {
        return err
    }
    defer file.Close()

    var lines []string
    found := false
    search := fmt.Sprintf("alias %s=", alias_name)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, search) {
            found = true
            continue // skip this line (delete)
        }
        lines = append(lines, line)
    }
    if err := scanner.Err(); err != nil {
        return err
    }

    if !found {
        fmt.Printf("Alias '%s' not found in %s.\n", alias_name, shellDir)
        return nil
    }

    // Write back all lines except the deleted alias
    err = os.WriteFile(shellDir, []byte(strings.Join(lines, "\n")+"\n"), 0644)
    if err != nil {
        return err
    }

    fmt.Printf("Alias '%s' deleted.\n", alias_name)
    logger.Println("Successfully executed delete function with the following parameters: ", alias_name)
    fmt.Println("Please restart your shell to apply changes.")
	return nil
}
func update(ctx context.Context, cmd *cli.Command) error {
   	alias_name := cmd.Args().Get(0)
    alias_cmd := cmd.Args().Get(1)
    if alias_name == "" || alias_cmd == "" {
        fmt.Println("Usage: update <alias_name> <alias_cmd>")
        return nil
    }

    // Read all lines from shellDir
    file, err := os.Open(shellDir)
    if err != nil {
        return err
    }
    defer file.Close()

    var lines []string
    found := false
    search := fmt.Sprintf("alias %s=", alias_name)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, search) {
            // Replace the alias line
            newLine := fmt.Sprintf("alias %s=\"%s\"", alias_name, alias_cmd)
            lines = append(lines, newLine)
            found = true
        } else {
            lines = append(lines, line)
        }
    }
    if err := scanner.Err(); err != nil {
        return err
    }

    if !found {
        // Add new alias
        newLine := fmt.Sprintf("alias %s=\"%s\"", alias_name, alias_cmd)
        lines = append(lines, newLine)
        fmt.Printf("Alias '%s' added as '%s'.\n", alias_name, alias_cmd)
    } else {
        fmt.Printf("Alias '%s' updated to '%s'.\n", alias_name, alias_cmd)
    }

    // Write back all lines
    err = os.WriteFile(shellDir, []byte(strings.Join(lines, "\n")+"\n"), 0644)
    if err != nil {
        return err
    }

    logger.Println("Successfully executed update function with the following parameters: ", alias_name, alias_cmd)
    fmt.Println("Please restart your shell to apply changes.")
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