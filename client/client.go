package client

type Client struct {
}

type Option func(*Client)

func New(opts ...Option) *Client {
	s := &Client{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (svr *Client) Start() error {
	return nil
}

func (svr *Client) Stop() {

}
