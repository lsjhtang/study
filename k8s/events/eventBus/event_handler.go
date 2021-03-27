package eventBus

import (
	"fmt"
	"reflect"
)

type EventHandler struct {
	fn reflect.Value
}

func NewEventHandler(fn interface{}) *EventHandler {
	rfn := reflect.ValueOf(fn)
	if rfn.Kind() != reflect.Func {
		panic("eventHandler kind error")
	}
	return &EventHandler{fn: rfn}
}

func (eh *EventHandler) Call(params ...interface{}) interface{} {
	values := eh.fn.Call(eh.Params(params))
	fmt.Println(values)
	dto := make([]interface{}, len(values))
	for k, v := range values {
		dto[k] = v.Interface()
	}
	return dto
}

func (eh *EventHandler) Params(params []interface{}) []reflect.Value {
	rv := make([]reflect.Value, len(params))
	for k, v := range params {
		if v == nil {
			rv[k] = reflect.New(eh.fn.Type().In(k)).Elem()
		} else {
			rv[k] = reflect.ValueOf(v)
		}
	}
	return rv
}
