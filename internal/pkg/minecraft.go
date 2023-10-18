package pkg

import (
	"MineMonitor/internal/pkg/types"
	"errors"
	"fmt"
	"time"

	"github.com/mcstatus-io/mcutil"
)

var (
	errFailedToConnect = errors.New("Failed to connect the server")
)

func Status(ip string, port uint16, timeout time.Duration) (types.Server, error) {

	res := make(chan types.Server, 1)

	go func() {
		response, err := mcutil.Status(ip, port)

		// Return empty when unavailable
		if err != nil || response == nil {
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
			PlayerList: _playerList,
		}

		if response.Version.NameRaw != "" {
			s.Version = response.Version.NameRaw
		}

		if response.MOTD.Clean != "" {
			s.Motd = response.MOTD.Clean
		}

		if response.Players.Online != nil || response.Players.Max != nil {
			s.Players = fmt.Sprintf("%d/%d", *response.Players.Online, *response.Players.Max)
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
