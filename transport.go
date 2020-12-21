package qlog

type VersionInformation struct {
	ServerVersions ByteString `json:"server_versions"`
	ClientVersions ByteString `json:"client_versions"`
	ChosenVersion  ByteString `json:"chosen_version"`
}

type ALPNInformation struct {
	ServerALPNs []string `json:"server_alp_ns"`
	ClientALPNs []string `json:"client_alp_ns"`
	ChosenALPN  string   `json:"chosen_alpn"`
}

type ParametersSet struct {
	Owner string `json:"owner"`

	ResumptionAllowed bool   `json:"resumption_allowed"`
	EarlyDataEnabled  bool   `json:"early_data_enabled"`
	TLSCipher         string `json:"tls_cipher"`
	AEADTagLength     uint8  `json:"aead_tag_length"`

	OriginalDestinationConnectionID ByteString `json:"original_destination_connection_id"`
	InitialSourceConnectionID       ByteString `json:"initial_source_connection_id"`
	RetrySourceConnectionID         ByteString `json:"retry_source_connection_id"`
	StatelessResetToken             Token      `json:"stateless_reset_token"`
	DisableActiveMigration          bool       `json:"disable_active_migration"`

	MaxIdleTimeout          uint64 `json:"max_idle_timeout"`
	MaxUdpPayloadSize       uint32 `json:"max_udp_payload_size"`
	AckDelayExponent        uint16 `json:"ack_delay_exponent"`
	MaxAckDelay             uint16 `json:"max_ack_delay"`
	ActiveConnectionIdLimit uint32 `json:"active_connection_id_limit"`

	InitialMaxData                 uint64 `json:"initial_max_data"`
	InitialMaxStreamDataBidiLocal  uint64 `json:"initial_max_stream_data_bidi_local"`
	InitialMaxStreamDataBidiRemote uint64 `json:"initial_max_stream_data_bidi_remote"`
	InitialMaxStreamDataUni        uint64 `json:"initial_max_stream_data_uni"`
	InitialMaxStreamsBidi          uint64 `json:"initial_max_streams_bidi"`
	InitialMaxStreamsUni           uint64 `json:"initial_max_streams_uni"`

	PreferredAddress PreferredAddress `json:"preferred_address"`
}

type PreferredAddress struct {
	IPV4 IPAddress `json:"ip_v4"`
	IPV6 IPAddress `json:"ip_v6"`

	PortV4 uint16 `json:"port_v4"`
	PortV6 uint16 `json:"port_v6"`

	ConnectionID        ByteString `json:"connection_id"`
	StatelessResetToken Token      `json:"stateless_reset_token"`
}

type ParametersRestored struct {
	DisableActiveMigration bool `json:"disable_active_migration"`

	MaxIdleTimeout          uint64 `json:"max_idle_timeout"`
	MaxUdpPayloadSize       uint32 `json:"max_udp_payload_size"`
	ActiveConnectionIdLimit uint32 `json:"active_connection_id_limit"`

	InitialMaxData                 uint64 `json:"initial_max_data"`
	InitialMaxStreamDataBidiLocal  uint64 `json:"initial_max_stream_data_bidi_local"`
	InitialMaxStreamDataBidiRemote uint64 `json:"initial_max_stream_data_bidi_remote"`
	InitialMaxStreamDataUni        uint64 `json:"initial_max_stream_data_uni"`
	InitialMaxStreamsBidi          uint64 `json:"initial_max_streams_bidi"`
	InitialMaxStreamsUni           uint64 `json:"initial_max_streams_uni"`
}

type PacketSent struct {
	Header              PacketHeader `json:"header"`
	Frames              []QUICFrame  `json:"frames"`
	IsCoalesced         bool         `json:"is_coalesced"`
	RetryToken          Token        `json:"retry_token"`
	StatelessResetToken ByteString   `json:"stateless_reset_token"`
	SupportedVersions   ByteString   `json:"supported_versions"`
	Raw                 RawInfo      `json:"raw"`
	DatagramId          uint32       `json:"datagram_id"`
}

type PacketReceived struct {
	Header              PacketHeader `json:"header"`
	Frames              []QUICFrame  `json:"frames"`
	IsCoalesced         bool         `json:"is_coalesced"`
	RetryToken          Token        `json:"retry_token"`
	StatelessResetToken ByteString   `json:"stateless_reset_token"`
	SupportedVersions   ByteString   `json:"supported_versions"`
	Raw                 RawInfo      `json:"raw"`
	DatagramID          uint32       `json:"datagram_id"`
}

type PacketDropped struct {
	Header     PacketHeader `json:"header"`
	Raw        RawInfo      `json:"raw"`
	DatagramID uint32       `json:"datagram_id"`
}

type PacketBuffered struct {
	Header     PacketHeader `json:"header"`
	Raw        RawInfo      `json:"raw"`
	DatagramID uint32       `json:"datagram_id"`
}

type PacketsACKed struct {
	PacketNumberSpace PacketNumberSpace `json:"packet_number_space"`
	PacketNumbers     []uint64          `json:"packet_numbers"`
}

type DatagramsSent struct {
	Count       uint16    `json:"count"`
	Raw         []RawInfo `json:"raw"`
	DatagramIDs []uint32  `json:"datagram_i_ds"`
}

type DatagramsReceived struct {
	Count       uint16    `json:"count"`
	Raw         []RawInfo `json:"raw"`
	DatagramIDs []uint32  `json:"datagram_i_ds"`
}

type DatagramDropped struct {
	Raw RawInfo `json:"raw"`
}

type StreamStateUpdated struct {
	StreamID uint64 `json:"stream_id"`

	// TODO
	//stream_type?:"unidirectional"|"bidirectional", // mainly useful when opening the stream

	Old StreamState `json:"old"`
	New StreamState `json:"new"`

	// TODO
	//stream_side?:"sending"|"receiving"
}

// TODO
type StreamState int

type FramesProcessed struct {
	Frames       []QUICFrame `json:"frames"`
	PacketNumber uint64      `json:"packet_number"`
}

type DataMoved struct {
	StreamID uint64 `json:"stream_id"`
	Offset   uint64 `json:"offset"`
	Length   uint64 `json:"length"`

	From string `json:"from"`
	To   string `json:"to"`

	Data ByteString `json:"data"`
}
