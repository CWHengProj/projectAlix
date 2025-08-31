package main
import (
	"os"
	"log"
	"github.com/urfave/cli/v3"
	"context"
	"fmt"
)
var logger *log.Logger

func main(){
	cmd := &cli.Command{
		Name:    "alix",
		Usage:   "summons alix, the manager of aliases",
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Creates a new alias",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("removed task template: ", cmd.Args().First())
					return nil
				},
			},
						{
				Name:  "delete",
				Usage: "Remove an existing alias",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("removed task template: ", cmd.Args().First())
					return nil
				},
			},			{
				Name:  "update",
				Usage: "Updates an existing alias",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("removed task template: ", cmd.Args().First())
					return nil
				},
			},			
			{
				Name:  "ls",
				Usage: "Shows all existing aliases",
				Commands: []*cli.Command{
                    {
                        Name:  "-la",
                        Usage: "Shows existing aliases in detail",
                        Action: func(ctx context.Context, cmd *cli.Command) error {
                            fmt.Println("new task template: ", cmd.Args().First())
                            return nil
                        },
                    },
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
func init(){
	file, err := os.OpenFile("alix.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //TODO: might want to change the permissions in the future
	if err!= nil {
		log.Fatal(err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Started new instance of app...")
	//TODO: detect the type of .*shrc user is using
}