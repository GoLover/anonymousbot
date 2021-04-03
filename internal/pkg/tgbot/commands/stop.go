package commands

import tb "gopkg.in/tucnak/telebot.v2"

func Stop(m *tb.Message) {
	if !m.Private() {
		return
	}
	_,err := handlerBot.Send(m.Sender,"goodbye dear")
	if err != nil {
		panic(err)
	}
}