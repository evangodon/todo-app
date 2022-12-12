package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func (cmd Cmd) Add() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add a todos",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "The name of the todo",
			},
		},
		Action: func(ctx *cli.Context) error {
			filename := ctx.String("file")
			todosList := newTodos(filename)
			todosList.parseFile()

			p := tea.NewProgram(initialModel())

			m, err := p.Run()
			if err != nil {
				log.Fatal(err)
			}

			todoName := m.(textinputModel).textInput.Value()

			if todoName == "" {
				return nil
			}

			todosList.uncompleted.addTodo(newTodo(todoName, uncompletedStatus))

			todosList.writeToFile()

			return nil
		},
	}
}