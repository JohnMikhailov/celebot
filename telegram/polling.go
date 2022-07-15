package telegram


import (
	"os"
	"os/signal"
	"log"
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
	log.Println("start polling")

	pollingMaster := newPolilingMaster(100)

	go func() {
		for {
			updates := bot.client.GetUpdates(updatesOffset)
			if !updates.Ok {
				log.Println("getting updates failed")
			}
			if len(updates.Result) > 0 {
				updatesOffset = updates.GetLastUpdateId() + 1
				for _, update := range updates.Result {
					pollingMaster.updatesQueue <- update
				}
			}
		}
	}()

	pollingMaster.wg.Add(pollingMaster.workers)

	for i := 0; i < pollingMaster.workers; i++ {
		go func() {
			update := <- pollingMaster.updatesQueue
			bot.processUpdate(update)
			pollingMaster.wg.Done()
		}()
	}

	<-listenForExit()
	pollingMaster.shutdown()
}

func (bot telegramBot) handleUpdate(update update) {
	message := update.Message
	bundle := newBundle(&bot, &message, &update)
	if message.IsReply() && message.ReplyToMessage.From.IsBot {
		// don't check bot name - work in groups is WIP
		if !bot.handlersRegistry.replyHandlerExist(message.ReplyToMessage.Text) {
			log.Println("Reply handler not registered! Skiping, original text was: " + message.Text)
			return
		}
		handler := bot.handlersRegistry.getReplyHandlerByMessageText(message.ReplyToMessage.Text)
		handler(bundle)
		return
	}

	command := message.getCommand()
	if !bot.handlersRegistry.handlerExists(command) {
		bot.handlersRegistry.getDefaultHandler()(bundle)
		return
	}

	handler := bot.handlersRegistry.getHandlerByCommand(command)
	handler(bundle)
}

func (bot telegramBot) handleGroupUpdate(update update) {
	message := update.Message
	bundle := newBundle(&bot, &message, &update)

	defaultHandler := bot.handlersRegistry.getDefaultHandler()
	defaultHandler(bundle)
}

func (bot telegramBot) processUpdateFromPrivateChat(update update) {
	bot.handleUpdate(update)
}

func (bot telegramBot) processUpdateFromGroup(update update) {
	bot.handleGroupUpdate(update)
}

func (bot telegramBot) processUpdate(update update) {
	if update.isFromGroup() {
		bot.processUpdateFromGroup(update)
		return
	}
	if update.isFromPrivateChat() {
		bot.processUpdateFromPrivateChat(update)
		return
	}
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
