package transport

type ServerState int

const (
	ServerStateReady ServerState = iota
	ServerStateInGracePeriod
	ServerStateInCleanupPeriod
)
