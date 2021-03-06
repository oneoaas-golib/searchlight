package hipchat

import (
	"github.com/appscode/go-notify"
	"github.com/kelseyhightower/envconfig"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

const Uid = "hipchat"

type Options struct {
	AuthToken string   `envconfig:"AUTH_TOKEN" required:"true"`
	To        []string `envconfig:"TO" required:"true"`
}

type client struct {
	opt  Options
	to   []string
	body string
}

var _ notify.ByChat = &client{}

func New(opt Options) *client {
	return &client{
		opt: opt,
		to:  opt.To,
	}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process(Uid, &opt)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func (c *client) WithBody(body string) notify.ByChat {
	c.body = body
	return c
}

func (c *client) To(to string, cc ...string) notify.ByChat {
	c.to = append([]string{to}, cc...)
	return c
}

func (c *client) Send() error {
	h := hipchat.NewClient(c.opt.AuthToken)

	for _, room := range c.to {
		_, err := h.Room.Notification(room, &hipchat.NotificationRequest{Message: c.body})
		if err != nil {
			return err
		}
	}
	return nil
}
