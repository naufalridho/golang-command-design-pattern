package main

import "fmt"

type commandSwitch struct {
	commands map[string]Command
}

func NewCommandSwitch() *commandSwitch {
	return &commandSwitch{commands: make(map[string]Command)}
}

func (s *commandSwitch) registerCommand(commandName string, command Command) {
	s.commands[commandName] = command
}

func (s *commandSwitch) runCommand(commandName string, input interface{}) error {
	command := s.commands[commandName]
	if command == nil {
		return fmt.Errorf("command not found")
	}
	command.execute(input)
	command.display()

	return nil
}