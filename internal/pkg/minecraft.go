package pkg

import (
	"errors"
	"fmt"
	"time"
	"github.com/mcstatus-io/mcutil"
  "MineMonitor/internal/pkg/types"
)

var (
	errFailedToConnect = errors.New("Failed to connect the server")
)

func Status(ip string, port uint16, timeout time.Duration) (types.Server, error) {

	res := make(chan types.Server, 1)

	go func() {
		response, err := mcutil.Status(ip, port)

    // Return empty when unavailable
		if err != nil {
			res <- types.Server{}
			return
		}

		var _playerList []string
		for _, p := range response.Players.Sample {
			_playerList = append(_playerList, p.NameClean)
		}

		s := types.Server{
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
		return types.Server{}, errFailedToConnect
	}
}
