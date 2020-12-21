package qlog

type ParametersSetRecovery struct {
	ReorderingThreshold uint16  `json:"reordering_threshold"`
	TimeThreshold       float64 `json:"time_threshold"`
	TimerGranularity    uint16  `json:"timer_granularity"`
	InitialRTT          float64 `json:"initial_rtt"`

	MaxDatagramSize               uint32  `json:"max_datagram_size"`
	InitialCongestionWindow       uint64  `json:"initial_congestion_window"`
	MinimumCongestionWindow       uint32  `json:"minimum_congestion_window"`
	LossReductionFactor           float64 `json:"loss_reduction_factor"`
	PersistentCongestionThreshold uint16  `json:"persistent_congestion_threshold"`
}

type MetricsUpdated struct {
	MinRTT      float64 `json:"min_rtt"`
	SmoothedRTT float64 `json:"smoothed_rtt"`
	LatestRTT   float64 `json:"latest_rtt"`
	RTTVariance float64 `json:"rtt_variance"`

	PTOCount uint16 `json:"pto_count"`

	CongestionWindow uint64 `json:"congestion_window"`
	BytesInFlight    uint64 `json:"bytes_in_flight"`

	SSThresh uint64 `json:"ss_thresh"`

	PacketsInFlight uint64 `json:"packets_in_flight"`

	PacingRate uint64 `json:"pacing_rate"`
}
type CongestionStateUpdated struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type LossTimerUpdated struct {
	// TODO
	//TimerType "ack"|"pto"

	PacketNumberSpace PacketNumberSpace `json:"packet_number_space"`

	// TODO
	//event_type:"set"|"expired"|"cancelled",

	Delta float64 `json:"delta"`
}

type PacketLost struct {
	Header PacketHeader `json:"header"`
	Frames []QUICFrame  `json:"frames"`
}

type MarkedForRetransmit struct {
	frames []QUICFrame
}
