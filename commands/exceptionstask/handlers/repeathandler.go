package handlers

import (
	"container/list"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/repeat"
)

func NewRepeatHandler(command *repeat.RepeatCommand, queue *list.List) *RepeatHandler {
	repeatHandler := &RepeatHandler{
		command: command,
		queue:   queue,
	}
	return repeatHandler
}

type RepeatHandler struct {
	command *repeat.RepeatCommand
	queue   *list.List
}

func (eh *RepeatHandler) Handle() {
	eh.queue.PushBack(eh.command)
}
