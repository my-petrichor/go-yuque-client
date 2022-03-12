package internal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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

func (c *Client) CheckLogin() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if viper.GetString("token") == "" {
			fmt.Println(ErrNoLogin)
			os.Exit(1)
		}
	}
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
