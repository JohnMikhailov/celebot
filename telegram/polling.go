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
			updates := bot.client.getUpdates(updatesOffset)
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

func (bot telegramBot) handleMessage(message message) {
	if message.IsReply() && message.ReplyToMessage.From.Username == "test_celebot" {
		if !bot.handlersRegistry.replyHandlerExist(message.ReplyToMessage.Text) {
			log.Println("Reply handler not registered! Skiping, original text was: " + message.Text)
			return
		}
		handler := bot.handlersRegistry.getReplyHandlerByMessageText(message.ReplyToMessage.Text)
		bundle := newBundle(&bot, &message)
		handler(bundle)
		return
	}

	command := message.getCommand()
	if !bot.handlersRegistry.handlerExists(command) {
		log.Println("Command handler not registered! Skiping, original text was: " + message.Text)
		return
	}

	handler := bot.handlersRegistry.getHandlerByCommand(command)
	bundle := newBundle(&bot, &message)
	handler(bundle)
}

func (bot telegramBot) processUpdateFromPrivateChat(update update) {
	message := update.Message
	bot.handleMessage(message)
}

func (bot telegramBot) processUpdateFromGroup(update update) {
	message := update.Message
	bot.handleMessage(message)
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
