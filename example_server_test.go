package neinp_test

import (
	"github.com/rbns/neinp"
	"fmt"
	"net"
)

func ExampleServer() {
	fs := &neinp.NopP2000{}

	l, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("listen:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept:", err)
			return
		}

		s := neinp.NewServer(fs)
		if err := s.Serve(c); err != nil {
			fmt.Println("serve:", err)
			return
		}
	}
}
