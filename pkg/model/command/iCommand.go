package command

type ICommand interface {
	Exist(cmd string) error
	Selector()
	Execute(fullCommand string) error
	Help(cmd string) //if cmd(command name is empty show all)
}
