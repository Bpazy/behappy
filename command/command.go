package command

import (
	"fmt"
)

type Commander interface {
	Keyword() string
	Run(event interface{}, arg string) string
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
	result := "我支持这些指令（需要@我）: \n"
	for i, cmder := range c.RegisteredCommanders {
		result = result + fmt.Sprintf("%d. %s\n", i+1, cmder.Example())
	}
	return result
}
