package command

import (
	"errors"
	"strings"
)

var (
	commands map[string]Command
)

type Parameter struct {
	ID         int
	IsOptional bool
}
type Command struct {
	ID          int
	FullCommand string
	Parameters  map[string]Parameter
	action      func(command *Command) error
}

func init() {

	commands = make(map[string]Command)

	commands["secret"] = Command{
		ID:          10001,
		FullCommand: "",
		action:      action,
		Parameters: map[string]Parameter{
			"--separated": {
				ID:         20001,
				IsOptional: true,
			}, "--central": {
				ID:         20002,
				IsOptional: true,
			}, "-key": {
				ID:         20003,
				IsOptional: false,
			},
		},
	}
}
func find(cmd, key string) (string, error) {
	parts := strings.Split(cmd, key)
	if len(parts) > 1 {
		return strings.Fields(parts[1])[0], nil
	}
	return "", errors.New("key not found")
}
func HasAnyKey(m map[string]Parameter, keys ...string) (bool, string) {
	for _, k := range keys {
		if _, ok := m[k]; ok {
			return true, k
		}
	}
	return false, ""
}
func action(c *Command) error {

	cmd, err := find(c.FullCommand, "-key")
	if err != nil {
		return err
	}
	exist, option := HasAnyKey(c.Parameters, "--separated", "--central")
	if !exist {
		return errors.New("option not found")
	}
	if option == "" {

	}
	if cmd == "" {

	}
	return nil
}
