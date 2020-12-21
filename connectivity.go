package qlog

type ServerListening struct {
	IPV4          IPAddress `json:"ip_v4"`
	IPV6          IPAddress `json:"ip_v6"`
	PortV4        uint32    `json:"port_v4"`
	PortV6        uint32    `json:"port_v6"`
	RetryRequired bool      `json:"retry_required"`
}

type ConnectionStarted struct {
	IPVersion   string     `json:"ip_version"`
	SrcIP       IPAddress  `json:"src_ip"`
	DstIP       IPAddress  `json:"dst_ip"`
	Protocol    string     `json:"protocol"`
	QUICVersion string     `json:"quic_version"`
	SrcPort     uint32     `json:"src_port"`
	DstPort     uint32     `json:"dst_port"`
	SrcCid      ByteString `json:"src_cid"`
	DstCid      ByteString `json:"dst_cid"`
}

type ConnectionClosed struct {
	Owner string `json:"owner"`
	//Connection_code?:TransportError | CryptoError | uint32
	//Application_code?:ApplicationError | uint32
	InternalCode uint32 `json:"internal_code"`
	Reason       string `json:"reason"`
}

type ConnectionIDUpdated struct {
	Owner string     `json:"owner"`
	Old   ByteString `json:"old"`
	New   ByteString `json:"new"`
}

type SpinBitUpdated struct {
	State bool `json:"state"`
}
