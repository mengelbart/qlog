package qlog

type IPAddress []byte

type PacketType int

const (
	Initial PacketType = iota
	Handshake
	ZeroRTT
	OneRTT
	Retry
	VersionNegotiation
	StatelessReset
	UnknownPacketType
)

func (p PacketType) String() string {
	switch p {
	case Initial:
		return "initial"
	case Handshake:
		return "handshake"
	case ZeroRTT:
		return "0RTT"
	case OneRTT:
		return "1RTT"
	case Retry:
		return "retry"
	case VersionNegotiation:
		return "version_negotiation"
	case StatelessReset:
		return "stateless_reset"
	case UnknownPacketType:
		fallthrough
	default:
		return "unknown"
	}
}

type Token struct {
	Type   string
	Length uint32
	Data   []byte
	//Details interface{}
}

type KeyType int

const (
	ServerInitialSecret KeyType = iota
	ClientInitialSecret

	ServerHandshakeSecret
	ClientHandshakeSecret

	Server0rttSecret
	Client0rttSecret

	Server1rttSecret
	Client1rttSecret
)

func (k KeyType) String() string {
	switch k {
	case ServerInitialSecret:
		return "server_initial_secret"
	case ClientInitialSecret:
		return "client_initial_secret"
	case ServerHandshakeSecret:
		return "server_handshake_secret"
	case ClientHandshakeSecret:
		return "client_handshake_secret"
	case Server0rttSecret:
		return "server_0rtt_secret"
	case Client0rttSecret:
		return "client_0rtt_secret"
	case Server1rttSecret:
		return "server_1rtt_secret"
	case Client1rttSecret:
		fallthrough
	default:
		return "client_1rtt_secret"
	}
}
