package telegram


import (
	"os"
	"os/signal"
	"fmt"
	"sync"
)

type pollingMaster struct {
	workers int
	updatesQueue chan update
	wg sync.WaitGroup
}

func newPolilingMaster(workers int) *pollingMaster {
	return &pollingMaster{workers: workers, updatesQueue: make(chan update, workers)}
}

func (pollingMaster *pollingMaster) shutdown() {
	close(pollingMaster.updatesQueue)
	pollingMaster.wg.Wait()
}

func (bot telegramBot) StartPolling() {
	updatesOffset := -1
	fmt.Println("start polling")

	pollingMaster := newPolilingMaster(100)

	go func() {
		for {
			updates := bot.client.getUpdates(updatesOffset)
			if !updates.Ok {
				fmt.Println("getting updates failed")  // TODO log it
			}
			if len(updates.Result) > 0 {
				updatesOffset = updates.GetLastUpdateId() + 1
				for _, update := range updates.Result {
					pollingMaster.updatesQueue <- update
				}
			} else {
				fmt.Println("no updates yet")
			}
		}
	}()

	pollingMaster.wg.Add(pollingMaster.workers)

	for i := 0; i < pollingMaster.workers; i++ {
		go func() {
			update := <- pollingMaster.updatesQueue
			bot.processMessage(update.Message)
			pollingMaster.wg.Done()
		}()
	}

	<-listenForExit()
	pollingMaster.shutdown()
}

func (bot telegramBot) processMessage(message message) {
	command := message.getCommand()
	if !bot.handlersRegistry.handlerExists(command) {
		fmt.Println("Command handler not registered! Skiping message")
		return
	}

	handler := bot.handlersRegistry.getTextHandlerByCommand(command)
	context := Context{bot: bot, Message: message}
	handler.Handle(&context)
}

func listenForExit() <-chan struct{} {
	runC := make(chan struct{}, 1)

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, os.Interrupt)

	go func() {
		defer close(runC)

		<-sc
	}()

	return runC
}
