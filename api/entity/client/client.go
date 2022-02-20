package client

import "errors"

type Client struct {
	IdClient       int    `json:"id_client"`
	Name           string `json:"name"`
	UniqueClientID int    `json:"unique_client_id"`
	MainEmail      string `json:"main_email"`
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) IsValid() error {
	if c.Name == "" {
		return errors.New("Client's name is empty")
	}

	if c.UniqueClientID == 0 {
		return errors.New("Client is invalid")
	}

	return nil
}
