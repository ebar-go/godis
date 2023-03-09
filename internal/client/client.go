package client

type Client struct{}

func New() *Client {
	return &Client{}
}

func (cli *Client) Run() error {
	return nil
}
