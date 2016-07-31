package main

import "fmt"

type Action func(args ...interface{})

type Events struct {
  registered map[string][]Action
}

func NewEvents() *Events {
  return &Events{
    registered: make(map[string][]Action),
  }
}

func (e *Events) Subscribe(name string, a Action) {
  e.registered[name] = append(e.registered[name], a)
}

func (e *Events) Raise(name string, args ...interface{}) {
  for _, x := range e.registered[name] {
    x(args...)
  }
}

func main() {
  events := NewEvents()
  events.Subscribe("test", func(args ...interface{}) {
    fmt.Printf("Got event \"test\" with args: %v\n", args)
  })

  events.Raise("test", true, 5, "hello", 35)
}
