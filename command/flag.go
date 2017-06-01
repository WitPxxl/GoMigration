package command

type Flag struct {
	Name         string
	DefaultValue interface{}
	HelpMessage  string
	Value        string
}
