package qlog

type ParametersSet struct {
	Owner string

	ResumptionAllowed bool
	EarlyDataEnabled  bool
	TLSCipher         string
	AEADTagLength     uint8

	OriginalDestinationConnectionID []byte
	InitialSourceConnectionID       []byte
	RetrySourceConnectionID         []byte
	StatelessResetToken             Token
	DisableActiveMigration          bool

	MaxIdleTimeout          uint64
	MaxUdpPayloadSize       uint32
	AckDelayExponent        uint16
	MaxAckDelay             uint16
	ActiveConnectionIdLimit uint32

	InitialMaxData                 uint64
	InitialMaxStreamDataBidiLocal  uint64
	InitialMaxStreamDataBidiRemote uint64
	InitialMaxStreamDataUni        uint64
	InitialMaxStreamsBidi          uint64
	InitialMaxStreamsUni           uint64

	PreferredAddress PreferredAddress
}

type PreferredAddress struct {
	IpV4 IPAddress
	IpV6 IPAddress

	PortV4 uint16
	PortV6 uint16

	ConnectionID        []byte
	StatelessResetToken Token
}
