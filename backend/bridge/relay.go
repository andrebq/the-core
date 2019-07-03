package bridge

import (
	"encoding/binary"
	"io"
	"math"
	"net"
	"sync"
	"time"
)

type (
	// Relay exposes an TCP relay to nodes to allow them to exchange packets.
	//
	// Since the relay won't manage any connection between nodes themselves, the
	// messages sent should be considered unreliable but ordered.
	Relay struct {
	}

	nodeMap map[NodeAddress]*node

	// Packet holds relay information
	Packet struct {
		Size    uint16
		From    NodeAddress
		To      NodeAddress
		Payload []byte
	}

	node struct {
		addr   NodeAddress
		output chan Packet
	}

	signal struct{}

	NodeAddress [16]byte
)

const (
	MaxPayload = math.MaxUint16 - 32
)

func NewPacket(to, from NodeAddress, payload []byte) Packet {
	if len(payload) > MaxPayload {
		payload = payload[:MaxPayload]
	}
	return Packet{
		To:      to,
		From:    from,
		Payload: payload,
		Size:    uint16(len(payload) + 32), /* TODO MAKE THIS A CONSTANT, SEE COMMENT BELOW */
	}
}

func NewRelay() *Relay {
	return &Relay{}
}

func (r *Relay) RunTCP(addr string) error {
	lst, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return r.Run(lst)
}

func (r *Relay) Run(lst net.Listener) error {
	newconn := make(chan net.Conn)
	newnode := make(chan *node)
	output := make(chan Packet)
	nodes := make(nodeMap)
	done := make(chan signal)
	go r.accept(newconn, lst)
	defer lst.Close()
	defer close(done)
	for {
		select {
		case conn := <-newconn:
			go r.handleConn(newnode, conn, output, done)
		case node := <-newnode:
			nodes[node.addr] = node
		case packet := <-output:
			nodes[packet.To].push(packet)
		}
	}
}

func (r *Relay) accept(newconn chan net.Conn, lst net.Listener) {
	defer close(newconn)
	for {
		conn, err := lst.Accept()
		if err != nil {
			return
		}
		newconn <- conn
	}
}

func (r *Relay) handleConn(newnode chan *node, conn net.Conn, output chan Packet, done chan signal) error {
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	p, err := ReadPacket(conn)
	if err != nil {
		return err
	}
	thisNode := &node{
		addr:   p.From,
		output: make(chan Packet, 10),
	}
	select {
	case newnode <- thisNode:
	case <-done:
		return nil
	}
	exit := make(chan signal)
	var once sync.Once
	closeBody := func() {
		once.Do(func() {
			conn.Close()
			// TODO: handle disconnect notification
			close(exit)
		})
	}

	go thisNode.pipeInput(output, conn, closeBody)
	go thisNode.pipeOutput(conn, thisNode.output, closeBody)
	<-exit
	return nil
}

func (n *node) push(p Packet) {
	if n == nil {
		return
	}
	select {
	case n.output <- p:
	default:
	}
}

func (n *node) pipeInput(output chan Packet, in net.Conn, closeFn func()) error {
	defer closeFn()
	for {
		in.SetReadDeadline(time.Now().Add(time.Minute * 5))
		p, err := ReadPacket(in)
		if err != nil {
			return err
		}
		select {
		case output <- p:
		default:
		}
	}
}

func (n *node) pipeOutput(out net.Conn, input chan Packet, closeFn func()) error {
	defer closeFn()
	for p := range input {
		out.SetWriteDeadline(time.Now().Add(time.Minute * 5))
		err := WritePacket(out, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func WritePacket(out io.Writer, p Packet) error {
	if err := binary.Write(out, binary.BigEndian, p.Size); err != nil {
		return err
	}
	if _, err := out.Write(p.From[:]); err != nil {
		return err
	}
	if _, err := out.Write(p.To[:]); err != nil {
		return err
	}
	_, err := out.Write(p.Payload)
	return err
}

func ReadPacket(in io.Reader) (p Packet, err error) {
	err = binary.Read(in, binary.BigEndian, &p.Size)
	if err != nil {
		return
	}
	_, err = io.ReadFull(in, p.From[:])
	if err != nil {
		return
	}
	_, err = io.ReadFull(in, p.To[:])
	if err != nil {
		return
	}
	p.Payload = make([]byte, p.Size-32 /* MAKE THIS A CONSTANT, THIS IS THE SIZE OF THE HEADER */)
	_, err = io.ReadFull(in, p.Payload)
	return
}
