package qlog

import (
	"fmt"

	"github.com/francoispqt/gojay"
)

type QLOGFile struct {
	QLOGVersion string
	QLOGFormat  string
	Title       string
	Description string
	Summary     Summary
	Traces      Traces
}

func (f *QLOGFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "qlog_version":
		return dec.String(&f.QLOGVersion)
	case "qlog_format":
		return dec.String(&f.QLOGFormat)
	case "title":
		return dec.String(&f.Title)
	case "description":
		return dec.String(&f.Description)
	case "summary":
		return dec.Object(&f.Summary)
	case "traces":
		return dec.Array(&f.Traces)
	}
	return nil
}

func (f *QLOGFile) NKeys() int {
	return 6
}

type Summary struct {
	TraceCount          uint32
	MaxDuration         uint64
	MaxOutgoingLossRate float64
	TotalEventCount     uint64
	ErrorCount          uint64
}

func (s *Summary) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "trace_count":
		return dec.Uint32(&s.TraceCount)
	case "max_duration":
		return dec.Uint64(&s.MaxDuration)
	case "max_outgoing_loss_rate":
		return dec.Float64(&s.MaxOutgoingLossRate)
	case "total_event_count":
		return dec.Uint64(&s.TotalEventCount)
	case "error_count":
		return dec.Uint64(&s.ErrorCount)
	}
	return nil
}

func (s *Summary) NKeys() int {
	return 5
}

type Traces []Trace

func (t *Traces) UnmarshalJSONArray(dec *gojay.Decoder) error {
	trace := Trace{}
	if err := dec.Object(&trace); err != nil {
		return err
	}
	*t = append(*t, trace)
	return nil
}

type Trace struct {
	Title         string
	Description   string
	Configuration Configuration
	CommonFields  CommonFields
	VantagePoint  VantagePoint
	EventFields   EventFields
	Events        Events
}

func (t *Trace) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "title":
		return dec.String(&t.Title)
	case "description":
		return dec.String(&t.Description)
	case "configuration":
		return dec.Object(&t.Configuration)
	case "common_fields":
		t.CommonFields = CommonFields{}
		return dec.Object(&t.CommonFields)
	case "vantage_point":
		return dec.Object(&t.VantagePoint)
	case "event_fields":
		return dec.Array(&t.EventFields)
	case "events":
		t.Events = Events{}
		if t.EventFields != nil {
			t.Events.Fields = t.EventFields
		}
		if t.CommonFields != nil {
			t.Events.CommonFields = t.CommonFields
		}
		return dec.Array(&t.Events)
	}
	return nil
}

func (t Trace) NKeys() int {
	return 6
}

type OriginalURIs []string

func (u *OriginalURIs) UnmarshalJSONArray(dec *gojay.Decoder) error {
	str := ""
	if err := dec.String(&str); err != nil {
		return err
	}
	*u = append(*u, str)
	return nil
}

type Configuration struct {
	TimeOffset   float64      `json:"time_offset"`
	OriginalURIs OriginalURIs `json:"original_ur_is"`

	// TODO: Implement custom fields?
	// https://tools.ietf.org/html/draft-marx-qlog-main-schema-02#section-3.3.1.3
}

func (c *Configuration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "time_offset":
		return dec.Float64(&c.TimeOffset)
	case "original_uris":
		return dec.Array(&c.OriginalURIs)
	}
	return nil
}

func (c Configuration) NKeys() int {
	return 2
}

type VantagePoint struct {
	Name string
	Type VantagePointType
	Flow VantagePointType
}

func (v *VantagePoint) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "name":
		return dec.String(&v.Name)
	case "type":
		s := ""
		if err := dec.String(&s); err != nil {
			return err
		}
		v.Type = VantagePointType(s)
		if !v.Type.isValid() {
			return fmt.Errorf("invalid vantage point: %v", v.Type)
		}
	case "flow":
		s := ""
		if err := dec.String(&s); err != nil {
			return err
		}
		v.Flow = VantagePointType(s)
		if !v.Flow.isValid() {
			return fmt.Errorf("invalid vantage point: %v", v.Type)
		}
	}
	return nil
}

func (v *VantagePoint) NKeys() int {
	return 3
}

const (
	Server              VantagePointType = "server"
	Client              VantagePointType = "client"
	Network             VantagePointType = "network"
	UnknownVantagePoint VantagePointType = "unknown"
)

type VantagePointType string

func (t VantagePointType) isValid() bool {
	return t == Server || t == Client || t == Network || t == UnknownVantagePoint
}

type CommonFields map[string]interface{}

func (c CommonFields) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var str interface{}
	err := dec.Interface(&str)
	if err != nil {
		return err
	}
	c[k] = str
	return nil
}

// we return 0, it tells the Decoder to decode all keys
func (c CommonFields) NKeys() int {
	return 0
}

type EventFields []string

func (f *EventFields) UnmarshalJSONArray(dec *gojay.Decoder) error {
	str := ""
	if err := dec.String(&str); err != nil {
		return err
	}
	*f = append(*f, str)
	return nil
}

type Events struct {
	Fields       EventFields
	Events       []Event
	CommonFields CommonFields
}

func (e *Events) UnmarshalJSONArray(dec *gojay.Decoder) error {
	if e.Fields != nil {
		event := Event{Fields: e.Fields}
		if err := dec.Array(&event); err != nil {
			return err
		}
		for k, v := range e.CommonFields {
			switch k {
			case "ODCID":
				event.ODCID = v.(string)
			case "group_id":
				event.GroupID = v.(string)
			case "reference_time":
				event.ReferenceTime = v.(float64)
			}
		}
		(*e).Events = append((*e).Events, event)
		return nil
	}

	event := Event{}
	if err := dec.Object(&event); err != nil {
		return err
	}
	(*e).Events = append((*e).Events, event)
	return nil
}

type Event struct {
	Fields EventFields

	Time          float64
	ReferenceTime float64
	Category      string
	Name          string
	Data          Data

	ProtocolType string
	GroupID      string
	TimeFormat   string
	ODCID        string
}

func (e *Event) UnmarshalJSONArray(dec *gojay.Decoder) error {

	f := e.Fields[dec.Index()]
	switch f {
	case "relative_time":
		return dec.Float64(&e.Time)
	case "category":
		return dec.String(&e.Category)
	case "event":
		return dec.String(&e.Name)
	case "data":
		e.Data = Data{}
		return dec.Object(&e.Data)
	}
	return nil
}

func (e *Event) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "time":
		return dec.Float64(&e.Time)
	case "reference_time":
		return dec.Float64(&e.ReferenceTime)
	case "name":
		return dec.String(&e.Name)
	case "protocol_type":
		return dec.String(&e.ProtocolType)
	case "group_id":
		return dec.String(&e.GroupID)
	case "time_format":
		return dec.String(&e.TimeFormat)
	}
	return nil
}

func (e *Event) NKeys() int {
	return 0
}

type Data map[string]interface{}

func (d Data) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var str interface{}
	err := dec.Interface(&str)
	if err != nil {
		return err
	}
	d[k] = str
	return nil
}

// we return 0, it tells the Decoder to decode all keys
func (d Data) NKeys() int {
	return 0
}
