package commands

import (
    "github.com/mix-go/console"
    "github.com/mix-go/mix-grpc-skeleton/commands"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "grpc:server",
            Usage: "Server demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"d", "daemon"},
                    Usage: "Run in the background",
                },
            },
            Command: &commands.GrpcServerCommand{},
        },
        console.CommandDefinition{
            Name:    "grpc:client",
            Usage:   "Client demo",
            Command: &commands.GrpcClientCommand{},
        },
    )
}
