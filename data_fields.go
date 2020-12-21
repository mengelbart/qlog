package qlog

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type ByteString []byte

func (b *ByteString) UnmarshalJSON(bs []byte) error {
	trim := bytes.Trim(bs, "\"")
	dst := make([]byte, hex.DecodedLen(len(trim)))
	n, err := hex.Decode(dst, trim)
	if err != nil {
		return err
	}
	*b = dst[:n]
	return nil
}

type RawInfo struct {
	Length        uint64
	PayloadLength uint64

	data []byte
}

type IPAddress string

type PacketType int

func (p *PacketType) UnmarshalJSON(bs []byte) error {
	var str string
	err := json.Unmarshal(bs, &str)
	if err != nil {
		return err
	}
	switch str {
	case "initial_packet_type":
		*p = InitialPacketType
	case "handshake":
		*p = Handshake
	case "0RTT":
		*p = ZeroRTT
	case "1RTT":
		*p = OneRTT
	case "retry":
		*p = Retry
	case "version_negotiation":
		*p = VersionNegotiation
	case "stateless_reset":
		*p = StatelessReset
	case "unknown_packet_type":
		*p = UnknownPacketType
	default:
		return fmt.Errorf("unknown packet type: %v", str)
	}
	return nil
}

const (
	InitialPacketType PacketType = iota
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
	case InitialPacketType:
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

type PacketNumberSpace int

func (p *PacketNumberSpace) UnmarshalJSON(bs []byte) error {
	var str string
	err := json.Unmarshal(bs, &str)
	if err != nil {
		return err
	}
	switch str {
	case "initial":
		*p = InitialPacketNumberSpace
	case "handshake":
		*p = HandshakePacketNumberSpace
	case "application_data":
		*p = ApplicationDataPacketNumberSpace
	default:
		return fmt.Errorf("invalid packet number space: %v", str)
	}
	return nil
}

const (
	InitialPacketNumberSpace PacketNumberSpace = iota
	HandshakePacketNumberSpace
	ApplicationDataPacketNumberSpace
)

func (p PacketNumberSpace) String() string {
	switch p {
	case InitialPacketNumberSpace:
		return "initial"
	case HandshakePacketNumberSpace:
		return "handshake"
	case ApplicationDataPacketNumberSpace:
		fallthrough
	default:
		return "application_data"
	}
}

type PacketHeader struct {
	PacketType    PacketType `json:"packet_type"`
	PacketNumber  uint64     `json:"packet_number"`
	Flags         uint8      `json:"flags"`
	Token         Token      `json:"token"`
	Length        uint16     `json:"length"`
	PayloadLength uint16     `json:"payload_length"`
	PacketSize    uint16     `json:"packet_size"`
	Version       ByteString `json:"version"`
	SCIL          uint8      `json:"scil"`
	DCIL          uint8      `json:"dcil"`
	SCID          ByteString `json:"scid"`
	DCID          ByteString `json:"dcid"`
}

type Token string

// TODO: Make token use this struct:
//type Token struct {
//	Type    string
//	Length  uint32
//	Data    ByteString
//	Details map[string]interface{}
//}

//func (t *Token) UnmarshalJSON(i []byte) error {
//
//}

type KeyType int

func (k *KeyType) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	switch s {
	case "server_initial_secret":
		*k = ServerInitialSecret
	case "client_initial_secret":
		*k = ClientInitialSecret
	case "server_handshake_secret":
		*k = ServerHandshakeSecret
	case "client_handshake_secret":
		*k = ClientHandshakeSecret
	case "server_0rtt_secret":
		*k = Server0rttSecret
	case "client_0rtt_secret":
		*k = Client0rttSecret
	case "server_1rtt_secret":
		*k = Server1rttSecret
	case "client_1rtt_secret":
		*k = Server1rttSecret
	default:
		return fmt.Errorf("unknown key type: %v", s)
	}
	return nil
}

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
