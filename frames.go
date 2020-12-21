package qlog

import (
	"encoding/json"
)

type QUICFrame struct {
	*PaddingFrame
	*PingFrame
	*ACKFrame
	*ResetStreamFrame
	*StopSendingFrame
	*CryptoFrame
	*NewTokenFrame
	*StreamFrame
	*MaxDataFrame
	*MaxStreamDataFrame
	*MaxStreamsFrame
	*DataBlockedFrame
	*StreamDataBlockedFrame
	*StreamsBlockedFrame
	*NewConnectionIDFrame
	*RetireConnectionIDFrame
	*PathChallengeFrame
	*PathResponseFrame
	*ConnectionCloseFrame
	*HandshakeDoneFrame
	*UnknownFrame
}

func (q *QUICFrame) UnmarshalJSON(bs []byte) error {
	tmp := struct {
		FrameType string `json:"frame_type"`
	}{}
	err := json.Unmarshal(bs, &tmp)
	if err != nil {
		return err
	}
	switch tmp.FrameType {
	case "Padding":
		var x PaddingFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.PaddingFrame = &x
		return nil
	case "ping":
		var x PingFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.PingFrame = &x
		return nil
	case "ack":
		var x ACKFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.ACKFrame = &x
		return nil
	case "reset_stream":
		var x ResetStreamFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.ResetStreamFrame = &x
		return nil
	case "stop_sending":
		var x StopSendingFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.StopSendingFrame = &x
		return nil
	case "crypto":
		var x CryptoFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.CryptoFrame = &x
		return nil
	case "new_token":
		var x NewTokenFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.NewTokenFrame = &x
		return nil
	case "stream":
		var x StreamFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.StreamFrame = &x
		return nil
	case "max_data":
		var x MaxDataFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.MaxDataFrame = &x
		return nil
	case "max_stream_data":
		var x MaxStreamDataFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.MaxStreamDataFrame = &x
		return nil
	case "max_streams":
		var x MaxStreamsFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.MaxStreamsFrame = &x
		return nil
	case "data_blocked":
		var x DataBlockedFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.DataBlockedFrame = &x
		return nil
	case "stream_data_blocked":
		var x StreamDataBlockedFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.StreamDataBlockedFrame = &x
		return nil
	case "streams_blocked":
		var x StreamsBlockedFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.StreamsBlockedFrame = &x
		return nil
	case "new_connection_id":
		var x NewConnectionIDFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.NewConnectionIDFrame = &x
		return nil
	case "retire_connection_id":
		var x RetireConnectionIDFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.RetireConnectionIDFrame = &x
		return nil
	case "path_challenge":
		var x PathChallengeFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.PathChallengeFrame = &x
		return nil
	case "path_response":
		var x PathResponseFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.PathResponseFrame = &x
		return nil
	case "connection_close":
		var x ConnectionCloseFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.ConnectionCloseFrame = &x
		return nil
	case "handshake_done":
		var x HandshakeDoneFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.HandshakeDoneFrame = &x
		return nil
	case "unknown":
		fallthrough
	default:
		var x UnknownFrame
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		q.UnknownFrame = &x
		return nil
	}
}

type PaddingFrame struct {
	FrameType     string `json:"frame_type"`
	Length        uint32 `json:"length"`
	PayloadLength uint32 `json:"payload_length"`
}

type PingFrame struct {
	FrameType     string `json:"frame_type"`
	Length        uint32 `json:"length"`
	PayloadLength uint32 `json:"payload_length"`
}

type ACKFrame struct {
	FrameType string  `json:"frame_type"`
	AckDelay  float64 `json:"ack_delay"`
	// TODO
	//acked_ranges []uint64, uint64]|[uint64]>;
	ECT1          uint64 `json:"ect1"`
	ECT0          uint64 `json:"ect0"`
	CE            uint64 `json:"ce"`
	Length        uint32 `json:"length"`
	PayloadLength uint32 `json:"payload_length"`
}

type ApplicationError int

const (
	HTTPNoError ApplicationError = iota
	HTTPGeneralProtocolError
	HTTPInternalError
	HTTPStreamCreationError
	HTTPClosedCriticalStream
	HTTPFrameUnexpected
	HTTPFrameError
	HTTPExcessiveLoad
	HTTPIdError
	HTTPSettingsError
	HTTPMissingSettings
	HTTPRequestRejected
	HTTPRequestCancelled
	HTTPRequestIncomplete
	HTTPEarlyResponse
	HTTPConnectError
	HTTPVersionFallback
)

type ErrorCode struct {
	*ApplicationError
	*uint32
}

type ResetStreamFrame struct {
	FrameType     string    `json:"frame_type"`
	StreamID      uint64    `json:"stream_id"`
	ErrorCode     ErrorCode `json:"error_code"`
	FinalSize     uint64    `json:"final_size"`
	Length        uint32    `json:"length"`
	PayloadLength uint32    `json:"payload_length"`
}

type StopSendingFrame struct {
	FrameType string `json:"frame_type"`
	StreamID  uint64 `json:"stream_id"`
	// TODO
	//error_code ApplicationError | uint32
	Length        uint32 `json:"length"`
	PayloadLength uint32 `json:"payload_length"`
}

type CryptoFrame struct {
	FrameType     string `json:"frame_type"`
	Offset        uint64 `json:"offset"`
	Length        uint64 `json:"length"`
	PayloadLength uint32 `json:"payload_length"`
}

type NewTokenFrame struct {
	FrameType string `json:"frame_type"`
	Token     Token  `json:"token"`
}

type StreamFrame struct {
	FrameType string `json:"frame_type"`
	Offset    uint64 `json:"offset"`
	Length    uint64 `json:"length"`
	Fin       bool   `json:"fin"`
	Raw       []byte `json:"raw"`
}

type MaxDataFrame struct {
	FrameType string `json:"frame_type"`
	Maximum   uint64 `json:"maximum"`
}

type MaxStreamDataFrame struct {
	FrameType string `json:"frame_type"`
	StreamID  uint64 `json:"stream_id"`
	Maximum   uint64 `json:"maximum"`
}

type MaxStreamsFrame struct {
	FrameType  string `json:"frame_type"`
	StreamType string `json:"stream_type"`
	Maximum    uint64 `json:"maximum"`
}

type DataBlockedFrame struct {
	FrameType string `json:"frame_type"`
	Limit     uint64 `json:"limit"`
}

type StreamDataBlockedFrame struct {
	FrameType string `json:"frame_type"`
	StreamID  uint64 `json:"stream_id"`
	Limit     uint64 `json:"limit"`
}

type StreamsBlockedFrame struct {
	FrameType string `json:"frame_type"`
	// TODO
	//streamType string = "bidirectional" | "unidirectional";
	Limit uint64 `json:"limit"`
}

type NewConnectionIDFrame struct {
	FrameType           string     `json:"frame_type"`
	SequenceNumber      uint32     `json:"sequence_number"`
	RetirePriorTo       uint32     `json:"retire_prior_to"`
	ConnectionIdLength  uint8      `json:"connection_id_length"`
	ConnectionId        ByteString `json:"connection_id"`
	StatelessResetToken ByteString `json:"stateless_reset_token"`
}

type RetireConnectionIDFrame struct {
	FrameType      string `json:"frame_type"`
	SequenceNumber uint32 `json:"sequence_number"`
}

type PathChallengeFrame struct {
	FrameType string     `json:"frame_type"`
	Data      ByteString `json:"data"`
}

type PathResponseFrame struct {
	FrameType string     `json:"frame_type"`
	Data      ByteString `json:"data"`
}

// TODO
//type ErrorSpace = "transport" | "application";

type ConnectionCloseFrame struct {
	FrameType string `json:"frame_type"`
	// TODO
	//ErrorSpace ErrorSpace
	// TODO
	//error_code?:TransportError | ApplicationError | uint32;
	RawErrorCode uint32 `json:"raw_error_code"`
	Reason       string `json:"reason"`
	// TODO
	//trigger_frame_type?:uint64 | string
}

type HandshakeDoneFrame struct {
	FrameType string `json:"frame_type"`
}

type UnknownFrame struct {
	FrameType    string     `json:"frame_type"`
	RawFrameType uint64     `json:"raw_frame_type"`
	RawLength    uint32     `json:"raw_length"`
	Raw          ByteString `json:"raw"`
}

//enum TransportError {
//no_error,
//internal_error,
//connection_refused,
//flow_control_error,
//stream_limit_error,
//stream_state_error,
//final_size_error,
//frame_encoding_error,
//transport_parameter_error,
//connection_id_limit_error,
//protocol_violation,
//invalid_token,
//application_error,
//crypto_buffer_exceeded
//}
