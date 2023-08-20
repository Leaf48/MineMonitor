package types

type Server struct {
	Ip         string
	Port       uint16
	Version    string
	Motd       string
	Players    string
	PlayerList []string
}
