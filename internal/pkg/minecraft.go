package pkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/mcstatus-io/mcutil"
)

var (
	errFailedToConnect = errors.New("Failed to connect the server")
)

type Server struct {
	Ip         string
	Port       uint16
	Version    string
	Motd       string
	Players    string
	PlayerList []string
}

func Status(ip string, port uint16, timeout time.Duration) (Server, error) {

	res := make(chan Server, 1)
	go func() {
		response, err := mcutil.Status(ip, port)

		if err != nil {
			res <- Server{}
			return
		}

		var _playerList []string
		for _, p := range response.Players.Sample {
			_playerList = append(_playerList, p.NameClean)
		}

		s := Server{
			Ip:         ip,
			Port:       port,
			Version:    response.Version.NameRaw,
			Motd:       response.MOTD.Clean,
			Players:    fmt.Sprintf("%d/%d", *response.Players.Online, *response.Players.Max),
			PlayerList: _playerList,
		}

		res <- s
	}()

	select {
	case result := <-res:
		return result, nil
	case <-time.After(timeout * time.Millisecond):
		return Server{}, errFailedToConnect
	}
}
