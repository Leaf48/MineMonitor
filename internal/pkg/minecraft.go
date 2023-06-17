package pkg

import (
	"errors"
	"time"

	"github.com/mcstatus-io/mcutil"
)

var (
	errFailedToConnect = errors.New("Failed to connect the server")
)

type server struct {
	Ip      string
	Port    uint16
	Version string
	Motd    string
}

func Status(ip string, port uint16, timeout time.Duration) (server, error) {

	res := make(chan server, 1)
	go func() {
		response, err := mcutil.Status(ip, port)

		if err != nil {
			res <- server{}
		}

		s := server{
			Ip:      ip,
			Port:    port,
			Version: response.Version.NameRaw,
			Motd:    response.MOTD.Clean,
		}

		res <- s
	}()

	select {
	case result := <-res:
		return result, nil
	case <-time.After(timeout * time.Millisecond):
		return server{}, errFailedToConnect
	}
}
