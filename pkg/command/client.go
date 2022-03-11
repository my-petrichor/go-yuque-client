package command

import "io"

type Cli interface {
	Err() io.Writer
}

type Client struct {
	err io.Writer
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Err() io.Writer {
	return c.err
}
