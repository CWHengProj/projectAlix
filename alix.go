package main
import (
	"os"
	"log"
	"github.com/urfave/cli/v3"
	"context"
)
var logger *log.Logger
var shellType string

func main(){
	cmd := &cli.Command{
		Name:    "alix",
		Usage:   "summons alix, the manager of aliases",
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Creates a new alias",
				Action: add,
			},
						{
				Name:  "delete",
				Usage: "Remove an existing alias",
				Action: delete,
			},			
			{
				Name:  "update",
				Usage: "Updates an existing alias",
				Action: update,
			},			
			{
				Name:  "ls",
				Usage: "Shows all existing aliases",
				Action: list,
				Commands: []*cli.Command{
                    {
                        Name:  "-la",
                        Usage: "Shows existing aliases in detail",
                        Action: listDetailed,
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
		
	shellType, err = detectShellType()
	if err != nil {
		log.Fatal(err)
	}
	logger.Println("Shell type detected: ", shellType)
}