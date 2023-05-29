package ports_input

type Command interface {
	Execute(command interface{}) (interface{}, error)
}
