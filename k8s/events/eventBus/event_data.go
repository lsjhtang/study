package eventBus

type EventData struct {
	Data interface{}
}

func NewEventData(data interface{}) *EventData {
	return &EventData{Data: data}
}

type EventDataChannel chan *EventData
