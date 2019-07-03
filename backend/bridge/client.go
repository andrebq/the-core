package bridge

import (
	uuid "github.com/satori/go.uuid"
	"net"
)

type (
	// Client exposes methods to interact with relays
	Client struct {
		conn net.Conn
		addr NodeAddress
	}
)

func DialTCP(addr string) (*Client, error) {
	myAddr := NodeAddress(uuid.NewV4())
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	err = WritePacket(conn, NewPacket(myAddr, myAddr, []byte("hello")))
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn,
		addr: myAddr}, nil
}

func (c *Client) Addr() NodeAddress {
	return c.addr
}

func (c *Client) Recv() (Packet, error) {
	return ReadPacket(c.conn)
}

func (c *Client) Send(to NodeAddress, payload []byte) error {
	return WritePacket(c.conn, NewPacket(to, c.addr, payload))
}
