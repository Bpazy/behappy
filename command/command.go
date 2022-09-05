package command

import (
	"fmt"
)

type MsgType int32

const (
	TypeImage MsgType = iota
	TypeText  MsgType = iota
	TypeNil   MsgType = iota
)

type Commander interface {
	Keyword() string
	Run(event interface{}, arg string) (MsgType, string)
	Example() string
	Hidden() bool
}

type DefaultCommander struct {
}

func (d *DefaultCommander) Keyword() string {
	panic("implement me")
}

func (d *DefaultCommander) Run(event interface{}, arg string) (MsgType, string) {
	panic("implement me")
}

func (d *DefaultCommander) Example() string {
	panic("implement me")
}

func (d *DefaultCommander) Hidden() bool {
	return false
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
	i := 0
	for _, cmder := range c.RegisteredCommanders {
		if cmder.Hidden() {
			continue
		}
		i++
		result = result + fmt.Sprintf("%d. %s\n", i, cmder.Example())
	}
	return result
}
