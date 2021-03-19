package lane

func (l *lane) TransportKey() string {
	return "__lane-" + l.Name
}
