package command

import (
	"errors"
	"strings"
)

var (
	commands map[string]Command
)

type Command struct {
	id          int
	optional    []string
	required    []string
	fullCommand []string
}

func init() {
	commands = make(map[string]Command)

	commands["secret"] = Command{
		id:       1,
		optional: []string{"--separate"},
		required: []string{"set", "get", "-value="},
	}

	commands["crypto"] = Command{
		id:       1,
		optional: []string{"--hotkey", "--stored"},
		required: []string{"enc", "dec", "-path="},
	}
}

func assign(input string) error {
	flds := make([]string, 0)
	for _, v := range strings.Fields(strings.ToLower(input)) {
		flds = append(flds, strings.Trim(v, " "))
	}
	cmd, ok := commands[flds[0]]
	if !ok {
		return errors.New("was kiri")
	}
	cmd.fullCommand = flds
	return nil
}

func validation(c Command) {

}
