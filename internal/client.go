package internal

import (
	"github.com/spf13/viper"
)

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) IsLogin() bool {
	return viper.GetString("token") != ""
}

func (c *Client) Login(token string) error {
	c.Token = token
	viper.Set("token", token)

	return viper.WriteConfig()
}

func (c *Client) Logout() error {
	c.Token = ""
	viper.Set("token", "")

	return viper.WriteConfig()
}
