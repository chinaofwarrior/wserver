package connection

import (
	"fmt"
	"reflect"
)

// Command represents a device with its associated data
type Command struct {
	DeviceID string
	Type     string
	Params   string
}

type CommandHandler interface {
	Execute(cmd Command)
}

// HandlerFactory is a relection-base factroy used create fucntion instance.
func HandlerFactory(handlerName string) (CommandHandler, error) {
	// use reflect
	t := reflect.TypeOf(nil)

	v := reflect.New(t).Interface()
	if handler, ok := v.(CommandHandler); ok {
		return handler, nil
	}
	return nil, fmt.Errorf("type docs not implement CommandHandler interface: %s", handlerName)
}
