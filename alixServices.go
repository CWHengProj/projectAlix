package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
)

func add(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("removed task template: ", cmd.Args().First())
					return nil
}
func delete(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("removed task template: ", cmd.Args().First())
					return nil
}
func update(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("removed task template: ", cmd.Args().First())
					return nil
}
func list(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("removed task template: ", cmd.Args().First())
					return nil
}
func listDetailed(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("removed task template: ", cmd.Args().First())
					return nil
}
func detectShellType() (string,error) {
	return "hehe", nil
}