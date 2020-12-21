package qlog

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type QLOGFile struct {
	QLOGVersion string  `json:"qlog_version"`
	QLOGFormat  string  `json:"qlog_format"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Summary     Summary `json:"summary"`
	Traces      Traces  `json:"traces"`
}

type QLOGFileNDJSON struct {
	QLOGFormat  string  `json:"qlog_format"`
	QLOGVersion string  `json:"qlog_version"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Summary     Summary `json:"summary"`
	Trace       Trace   `json:"trace"`
}

func (q *QLOGFileNDJSON) UnmarshalNDJSON(bs []byte) error {
	reader := bytes.NewReader(bs)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	qlogFileNDJSON := scanner.Bytes()
	err := json.Unmarshal(qlogFileNDJSON, q)
	if err != nil {
		return err
	}
	for scanner.Scan() {
		bs := scanner.Bytes()
		event := EventWrapper{
			CommonFields: q.Trace.CommonFields,
		}
		err := json.Unmarshal(bs, &event)
		if err != nil {
			return err
		}
		q.Trace.Events.Events = append(q.Trace.Events.Events, event)
	}
	return nil
}

type Summary struct {
	TraceCount          uint32  `json:"trace_count"`
	MaxDuration         uint64  `json:"max_duration"`
	MaxOutgoingLossRate float64 `json:"max_outgoing_loss_rate"`
	TotalEventCount     uint64  `json:"total_event_count"`
	ErrorCount          uint64  `json:"error_count"`
}

type Traces []Trace

type Trace struct {
	*TraceError
	*TraceObject
}

func (t *Trace) UnmarshalJSON(bytes []byte) error {
	tmp := struct {
		ErrorDescription string       `json:"error_description"`
		EventFields      EventFields  `json:"event_fields"`
		CommonFields     CommonFields `json:"common_fields"`
	}{}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	if len(tmp.ErrorDescription) > 0 {
		var e TraceError
		if err := json.Unmarshal(bytes, &e); err != nil {
			return err
		}
		t.TraceError = &e
		t.TraceObject = nil
	} else {
		var o TraceObject
		o.EventFields = tmp.EventFields
		o.Events = Events{Fields: tmp.EventFields, CommonFields: tmp.CommonFields}
		if err := json.Unmarshal(bytes, &o); err != nil {
			return err
		}
		t.TraceError = nil
		t.TraceObject = &o
	}
	return nil
}

type TraceError struct {
	ErrorDescription string       `json:"error_description"`
	URI              string       `json:"uri"`
	VantagePoint     VantagePoint `json:"vantage_point"`
}

type TraceObject struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Configuration Configuration `json:"configuration"`
	CommonFields  CommonFields  `json:"common_fields"`
	VantagePoint  VantagePoint  `json:"vantage_point"`
	EventFields   EventFields   `json:"event_fields"`
	Events        Events        `json:"events"`
}

type Configuration struct {
	TimeOffset   float64  `json:"time_offset"`
	OriginalURIs []string `json:"original_uris"`

	// TODO: Implement custom fields?
	// https://tools.ietf.org/html/draft-marx-qlog-main-schema-02#section-3.3.1.3
}

type CommonFields map[string]interface{}

type VantagePoint struct {
	Name string           `json:"Name"`
	Type VantagePointType `json:"type"`
	Flow VantagePointType `json:"flow"`
}

type VantagePointType string

func (t *VantagePointType) UnmarshalJSON(bytes []byte) error {
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}
	if !VantagePointType(str).isValid() {
		return errors.New("invalid vantage point")
	}
	*t = VantagePointType(str)
	return nil
}

const (
	Server              VantagePointType = "server"
	Client              VantagePointType = "client"
	Network             VantagePointType = "network"
	UnknownVantagePoint VantagePointType = "unknown"
)

func (t VantagePointType) isValid() bool {
	return t == Server || t == Client || t == Network || t == UnknownVantagePoint
}

type EventFields []string

type Events struct {
	Fields       EventFields
	CommonFields CommonFields

	Events []EventWrapper `json:"events"`
}

var eventFields EventFields
var commonFields CommonFields

func (e *Events) UnmarshalJSON(bs []byte) error {
	eventFields = e.Fields
	commonFields = e.CommonFields
	err := json.Unmarshal(bs, &e.Events)
	if err != nil {
		return err
	}
	return nil
}

type EventWrapper struct {
	Fields       EventFields
	CommonFields CommonFields
	*Event
}

func (e *EventWrapper) UnmarshalJSON(bs []byte) error {
	if e.Fields == nil {
		e.Fields = eventFields
	}
	if e.CommonFields == nil {
		e.CommonFields = commonFields
	}
	var x interface{}
	if err := json.Unmarshal(bs, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case []interface{}:
		js := make(map[string]interface{})
		for i, x := range v {
			js[e.Fields[i]] = x
		}
		object, err := json.Marshal(js)
		if err != nil {
			return err
		}
		var event Event
		name, err := getEventNameArray(e.Fields, v)
		if err != nil {
			return err
		}
		event.Name = name
		event.fillCommonFields(e.CommonFields)
		err = json.Unmarshal(object, &event)
		if err != nil {
			return err
		}
		e.Event = &event
	case map[string]interface{}:
		var event Event
		name, err := getEventName(v)
		if err != nil {
			return err
		}
		event.Name = name
		event.Data = Data{Name: name}
		event.fillCommonFields(e.CommonFields)
		err = json.Unmarshal(bs, &event)
		if err != nil {
			return err
		}
		e.Event = &event
	default:
		fmt.Printf("%T: %v\n", v, v)
	}
	return nil
}

func indexOf(x string, ls EventFields) int {
	for i, e := range ls {
		if e == x {
			return i
		}
	}
	return -1
}

func getEventNameArray(fields EventFields, values []interface{}) (string, error) {
	m := make(map[string]interface{})
	nameFields := []string{"name", "category", "event", "type"}
	for _, n := range nameFields {
		if i := indexOf(n, fields); i >= 0 {
			m[n] = values[i]
		}
	}
	return getEventName(m)
}

func getEventName(data map[string]interface{}) (string, error) {
	var name string
	if n, ok := data["name"]; ok {
		if s, ok := n.(string); ok {
			name = s
		}
	}
	if strings.Contains(name, ":") {
		return name, nil
	}

	var category string
	if c, ok := data["category"]; ok {
		if s, ok := c.(string); ok {
			category = s
		}
	}

	var event string
	if e, ok := data["event"]; ok {
		if s, ok := e.(string); ok {
			event = s
		}
	}

	var eventType string
	if t, ok := data["type"]; ok {
		if s, ok := t.(string); ok {
			eventType = s
		}
	}

	if len(event) <= 0 && len(eventType) <= 0 {
		return "", fmt.Errorf("invalid event Name: Name=%v, category=%v, event=%v, type=%v", name, category, event, eventType)
	}

	if len(event) > 0 {
		return fmt.Sprintf("%v:%v", category, event), nil
	}

	return fmt.Sprintf("%v:%v", category, eventType), nil
}

type Event struct {
	Fields EventFields

	Time float64 `json:"time"`
	Name string  `json:"Name"`
	Data Data    `json:"data"`

	ProtocolType string `json:"protocol_type"`
	GroupID      string `json:"group_id"`
	TimeFormat   string `json:"time_format"`

	ReferenceTime float64 `json:"reference_time"`
	RelativeTime  float64 `json:"relative_time"`
	ODCID         string  `json:"odcid"`
}

func (e *Event) fillCommonFields(common map[string]interface{}) {
	for k, v := range common {
		switch k {
		case "ODCID":
			e.ODCID = v.(string)
		case "reference_time":
			e.ReferenceTime = v.(float64)
		case "group_id":
			e.GroupID = v.(string)
		}
	}
}

func (e *Event) UnmarshalJSON(bs []byte) error {
	eventType := struct {
		Category  string `json:"category"`
		Event     string `json:"event"`
		Name      string `json:"Name"`
		EventType string `json:"type"`

		RelativeTime float64 `json:"relative_time"`
	}{}

	err := json.Unmarshal(bs, &eventType)
	if err != nil {
		return err
	}
	name, err := getEventName(map[string]interface{}{
		"name":     eventType.Name,
		"category": eventType.Category,
		"event":    eventType.Event,
		"type":     eventType.EventType,
	})
	if err != nil {
		return err
	}

	tmp := struct {
		Time float64 `json:"time"`
		Name string  `json:"name"`
		Data Data    `json:"data"`

		ProtocolType string `json:"protocol_type"`
		GroupID      string `json:"group_id"`
		TimeFormat   string `json:"time_format"`

		RelativeTime  float64 `json:"relative_time"`
		ReferenceTime float64 `json:"reference_time"`
		ODCID         string  `json:"odcid"`
	}{
		Name: name,
		Data: Data{Name: name},
	}

	err = json.Unmarshal(bs, &tmp)
	if err != nil {
		return err
	}
	newEvent := Event{
		Data:          tmp.Data,
		Time:          tmp.Time,
		ProtocolType:  tmp.ProtocolType,
		GroupID:       tmp.GroupID,
		TimeFormat:    tmp.TimeFormat,
		ReferenceTime: tmp.ReferenceTime,
		RelativeTime:  tmp.RelativeTime,
		ODCID:         tmp.ODCID,
	}
	if len(newEvent.Name) <= 0 {
		newEvent.Name = e.Name
	}
	if newEvent.Time <= 0 {
		newEvent.Time = e.ReferenceTime + tmp.RelativeTime
	}
	if newEvent.ReferenceTime <= 0 {
		newEvent.ReferenceTime = e.ReferenceTime
	}
	if newEvent.RelativeTime <= 0 {
		// TODO: This obviously only works if e.Time is relative time
		newEvent.RelativeTime = newEvent.Time
	}
	if len(newEvent.GroupID) <= 0 {
		newEvent.GroupID = e.GroupID
	}
	if len(newEvent.ODCID) <= 0 {
		newEvent.ODCID = e.ODCID
	}
	*e = newEvent
	return nil
}

type Data struct {
	Name string

	// Connectivity
	*ServerListening
	*ConnectionStarted
	*ConnectionClosed
	*ConnectionIDUpdated
	*SpinBitUpdated

	// Security
	*KeyUpdated
	*KeyRetired

	// Transport
	*VersionInformation
	*ALPNInformation
	*ParametersSet
	*ParametersRestored
	*PacketSent
	*PacketReceived
	*PacketDropped
	*PacketBuffered
	*PacketsACKed
	*DatagramsSent
	*DatagramsReceived
	*DatagramDropped
	*StreamStateUpdated
	*FramesProcessed
	*DataMoved

	// Recovery
	*ParametersSetRecovery
	*MetricsUpdated
	*CongestionStateUpdated
	*LossTimerUpdated
	*PacketLost
	*MarkedForRetransmit

	// HTTP3
}

func (d *Data) UnmarshalJSON(bs []byte) error {
	switch d.Name {
	case "security:key_updated":
		var x KeyUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.KeyUpdated = &x
		return nil
	case "security:key_retired":
		var x KeyRetired
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.KeyRetired = &x
		return nil
	case "transport:version_information":
		var x VersionInformation
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.VersionInformation = &x
		return nil
	case "transport:alpn_information":
		var x ALPNInformation
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ALPNInformation = &x
		return nil
	case "transport:parameters_set":
		var x ParametersSet
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ParametersSet = &x
		return nil
	case "transport:parameters_restored":
		var x ParametersRestored
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ParametersRestored = &x
		return nil
	case "transport:packet_sent":
		var x PacketSent
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.PacketSent = &x
		return nil
	case "transport:packet_received":
		var x PacketReceived
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.PacketReceived = &x
		return nil
	case "transport:packet_dropped":
		var x PacketDropped
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.PacketDropped = &x
		return nil
	case "transport:packet_buffered":
		var x PacketBuffered
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.PacketBuffered = &x
		return nil
	case "transport:packets_acked":
		var x PacketsACKed
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.PacketsACKed = &x
		return nil
	case "transport:datagrams_sent":
		var x DatagramsSent
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.DatagramsSent = &x
		return nil
	case "transport:datagrams_received":
		var x DatagramsReceived
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.DatagramsReceived = &x
		return nil
	case "transport:datagram_dropped":
		var x DatagramDropped
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.DatagramDropped = &x
		return nil
	case "transport:stream_state_updated":
		var x StreamStateUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.StreamStateUpdated = &x
		return nil
	case "transport:frames_processed":
		var x FramesProcessed
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.FramesProcessed = &x
		return nil
	case "transport:data_moved":
		var x DataMoved
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.DataMoved = &x
		return nil
	case "recovery:parameters_set_recovery":
		var x ParametersSetRecovery
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ParametersSetRecovery = &x
		return nil
	case "recovery:metrics_updated":
		var x MetricsUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.MetricsUpdated = &x
		return nil
	case "recovery:congestion_state_updated":
		var x CongestionStateUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.CongestionStateUpdated = &x
		return nil
	case "recovery:loss_timer_updated":
		var x LossTimerUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.LossTimerUpdated = &x
		return nil
	case "recovery:packet_lost":
		var x PacketLost
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.PacketLost = &x
		return nil
	case "recovery:marked_for_retransmit":
		var x MarkedForRetransmit
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.MarkedForRetransmit = &x
		return nil
	case "connectivity:server_listening":
		var x ServerListening
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ServerListening = &x
		return nil
	case "transport:connection_started":
		fallthrough
	case "connectivity:connection_started":
		var x ConnectionStarted
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ConnectionStarted = &x
		return nil
	case "transport:connection_closed":
		fallthrough
	case "connectivity:connection_closed":
		var x ConnectionClosed
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ConnectionClosed = &x
		return nil
	case "connectivity:connection_id_updated":
		var x ConnectionIDUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.ConnectionIDUpdated = &x
		return nil
	case "connectivity:spin_bit_updated":
		var x SpinBitUpdated
		err := json.Unmarshal(bs, &x)
		if err != nil {
			return err
		}
		d.SpinBitUpdated = &x
		return nil
	default:
		return fmt.Errorf("unknown event: %v", d.Name)
	}
}
