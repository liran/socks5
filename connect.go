package socks5

import (
	"fmt"
	"io"
	"log"
	"net"
)

func Dial(network, remoteAddr, localIP string) (net.Conn, error) {
	lAddr, err := net.ResolveTCPAddr(network, fmt.Sprintf("%s:0", localIP))
	if err != nil {
		return nil, err
	}
	rAddr, err := net.ResolveTCPAddr(network, remoteAddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP(network, lAddr, rAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Connect remote conn which u want to connect with your dialer
// Error or OK both replied.
func (r *Request) Connect(c *net.TCPConn) (*net.TCPConn, error) {
	w := io.Writer(c)
	if Debug {
		log.Println("Call:", r.Address())
	}
	localAddr := c.LocalAddr().String()
	localIP, _, err := net.SplitHostPort(localAddr)
	if err != nil {
		return nil, err
	}
	tmp, err := Dial("tcp", r.Address(), localIP)
	if err != nil {
		var p *Reply
		if r.Atyp == ATYPIPv4 || r.Atyp == ATYPDomain {
			p = NewReply(RepHostUnreachable, ATYPIPv4, []byte{0x00, 0x00, 0x00, 0x00}, []byte{0x00, 0x00})
		} else {
			p = NewReply(RepHostUnreachable, ATYPIPv6, []byte(net.IPv6zero), []byte{0x00, 0x00})
		}
		if _, err := p.WriteTo(w); err != nil {
			return nil, err
		}
		return nil, err
	}
	rc := tmp.(*net.TCPConn)

	a, addr, port, err := ParseAddress(rc.LocalAddr().String())
	if err != nil {
		var p *Reply
		if r.Atyp == ATYPIPv4 || r.Atyp == ATYPDomain {
			p = NewReply(RepHostUnreachable, ATYPIPv4, []byte{0x00, 0x00, 0x00, 0x00}, []byte{0x00, 0x00})
		} else {
			p = NewReply(RepHostUnreachable, ATYPIPv6, []byte(net.IPv6zero), []byte{0x00, 0x00})
		}
		if _, err := p.WriteTo(w); err != nil {
			return nil, err
		}
		return nil, err
	}
	p := NewReply(RepSuccess, a, addr, port)
	if _, err := p.WriteTo(w); err != nil {
		return nil, err
	}

	return rc, nil
}
