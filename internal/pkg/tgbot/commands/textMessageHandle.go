package commands

import (
	"github.com/golover/anonymousbot/internal/pkg/usermanagement"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)
type userobject struct {
	userid string
}
func (c userobject)Recipient() string{
	return c.userid
}
func TextMessageHandle(m *tb.Message) {
	handlerBot.Send(userobject{usermanagement.GetCurrentContact(strconv.Itoa(m.Sender.ID))},m.Text)
	handlerBot.Send(m.Sender,"your message has been sent successfully.")
}
