package command

import (
	"fmt"
)

type Commander interface {
	Keyword() string
	Run(event interface{}, arg string)
	Example() string
}

// CommandersChain defines a Commander array.
type CommandersChain struct {
	RegisteredCommanders []Commander
}

var Commanders = CommandersChain{}

func (c *CommandersChain) Register(cmd Commander) {
	c.RegisteredCommanders = append(c.RegisteredCommanders, cmd)
}

func (c *CommandersChain) GetHelpMessage() string {
	result := "我支持这些查询: \n\n--------\n\n"
	for i, cmder := range c.RegisteredCommanders {
		result = result + fmt.Sprintf("%d. %s\n", i, cmder.Example())
	}
	return result
}
