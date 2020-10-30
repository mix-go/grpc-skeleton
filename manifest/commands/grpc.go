package commands

import (
    "github.com/mix-go/console"
    "github.com/mix-go/mix-grpc-skeleton/commands"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "grpc:server",
            Usage: "\tEcho demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"n", "name"},
                    Usage: "Your name",
                },
                {
                    Names: []string{"say"},
                    Usage: "\tSay ...",
                },
            },
            Command: &commands.GrpcServerCommand{},
        },
        console.CommandDefinition{
            Name:  "grpc:client",
            Usage: "\tEcho demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"n", "name"},
                    Usage: "Your name",
                },
                {
                    Names: []string{"say"},
                    Usage: "\tSay ...",
                },
            },
            Command: &commands.GrpcClientCommand{},
        },
    )
}
