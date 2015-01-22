package client

import (
	"fmt"
	"github.com/getlantern/balancer"
	"github.com/getlantern/fronted"
	"log"
)

type frontedServer struct {
	Host string
	Port int
}

// Wraps a fronted.Dialer with a balancer.Dialer.
func (s *frontedServer) dialer() *balancer.Dialer {
	fd := fronted.NewDialer(&fronted.Config{
		Host: s.Host,
		Port: s.Port,
	})

	masqueradeQualifier := ""

	return &balancer.Dialer{
		Label:  fmt.Sprintf("fronted proxy at %s:%d%s", s.Host, s.Port, masqueradeQualifier),
		Weight: 1,
		QOS:    0,
		Dial:   fd.Dial,
		OnClose: func() {
			err := fd.Close()
			if err != nil {
				log.Printf("Unable to close fronted dialer: %s", err)
			}
		},
	}
}
