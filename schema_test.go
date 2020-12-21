package qlog

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func TestTraces(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Traces
		wantErr bool
	}{
		{
			name: "TraceError",
			args: args{
				bs: []byte(`[{"error_description": "error", "uri": "err.or", "vantage_point": { "type": "server"}}]`),
			},
			want: Traces{Trace{
				TraceError: &TraceError{
					ErrorDescription: "error",
					URI:              "err.or",
					VantagePoint: VantagePoint{
						Type: "server",
					},
				},
			}},
			wantErr: false,
		},
		{
			name: "TraceObject",
			args: args{
				bs: []byte(`[{"title": "Trace", "description": "Description"}]`),
			},
			want: Traces{Trace{
				TraceObject: &TraceObject{
					Title:       "Trace",
					Description: "Description",
				},
			}},
			wantErr: false,
		},
		{
			name: "Mixed",
			args: args{
				bs: []byte(`[{"title": "Trace", "description": "Description"}, {"error_description": "error", "uri": "err.or", "vantage_point": { "type": "server"}}]`),
			},
			want: Traces{
				Trace{
					TraceObject: &TraceObject{
						Title:       "Trace",
						Description: "Description",
					},
				},
				Trace{
					TraceError: &TraceError{
						ErrorDescription: "error",
						URI:              "err.or",
						VantagePoint: VantagePoint{
							Type: "server",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Traces
			err := json.Unmarshal(tt.args.bs, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("json.Unmarshal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestVantagePointType(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    VantagePointType
		wantErr bool
	}{
		{
			name: "server",
			args: args{
				bs: []byte(`"server"`),
			},
			want:    Server,
			wantErr: false,
		},
		{
			name: "client",
			args: args{
				bs: []byte(`"client"`),
			},
			want:    Client,
			wantErr: false,
		},
		{
			name: "network",
			args: args{
				bs: []byte(`"network"`),
			},
			want:    Network,
			wantErr: false,
		},
		{
			name: "unknown",
			args: args{
				bs: []byte(`"unknown"`),
			},
			want:    UnknownVantagePoint,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				bs: []byte(`"invalid"`),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got VantagePointType
			err := json.Unmarshal(tt.args.bs, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != got {
				t.Errorf("json.Unmarshal() mismatch want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestTestJSON(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    QLOGFile
		wantErr bool
	}{
		{
			name: "server.json",
			args: args{
				reader: mustOpen("server.json"),
			},
			want: QLOGFile{
				QLOGVersion: "draft-02-wip",
				Title:       "quic-go qlog",
				Traces: Traces{
					Trace{
						TraceObject: &TraceObject{
							VantagePoint: VantagePoint{
								Type: "server",
							},
							CommonFields: CommonFields{
								"ODCID":          "4375c1a55669d5dec87d34611c2a3d",
								"group_id":       "4375c1a55669d5dec87d34611c2a3d",
								"reference_time": 1604843883465.9907,
							},
							EventFields: EventFields{
								"relative_time",
								"category",
								"event",
								"data",
							},
							Events: Events{
								CommonFields: CommonFields{
									"ODCID":          "4375c1a55669d5dec87d34611c2a3d",
									"group_id":       "4375c1a55669d5dec87d34611c2a3d",
									"reference_time": 1604843883465.9907,
								},
								Fields: EventFields{
									"relative_time",
									"category",
									"event",
									"data",
								},
								Events: []EventWrapper{
									{
										CommonFields: CommonFields{
											"ODCID":          "4375c1a55669d5dec87d34611c2a3d",
											"group_id":       "4375c1a55669d5dec87d34611c2a3d",
											"reference_time": 1604843883465.9907,
										},
										Fields: EventFields{
											"relative_time",
											"category",
											"event",
											"data",
										},
										Event: &Event{
											Time: 1604843883465.9907 + 0.875758,
											Name: "security:key_updated",
											Data: Data{
												Name: "security:key_updated",
												KeyUpdated: &KeyUpdated{
													Trigger: "tls",
													KeyType: ClientHandshakeSecret,
												},
											},
											ODCID:         "4375c1a55669d5dec87d34611c2a3d",
											GroupID:       "4375c1a55669d5dec87d34611c2a3d",
											ReferenceTime: 1604843883465.9907,
										},
									},
									{
										CommonFields: CommonFields{
											"ODCID":          "4375c1a55669d5dec87d34611c2a3d",
											"group_id":       "4375c1a55669d5dec87d34611c2a3d",
											"reference_time": 1604843883465.9907,
										},
										Fields: EventFields{
											"relative_time",
											"category",
											"event",
											"data",
										},
										Event: &Event{
											Time: 1604843883465.9907 + 0.875758,
											Name: "security:key_updated",
											Data: Data{
												Name: "security:key_updated",
												KeyUpdated: &KeyUpdated{
													Trigger: "tls",
													KeyType: ClientHandshakeSecret,
												},
											},
											ODCID:         "4375c1a55669d5dec87d34611c2a3d",
											GroupID:       "4375c1a55669d5dec87d34611c2a3d",
											ReferenceTime: 1604843883465.9907,
										},
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got QLOGFile
			bs, err := ioutil.ReadAll(tt.args.reader)
			if err != nil {
				t.Errorf("ioutil.ReadAll() error = %v", err)
			}
			err = json.Unmarshal(bs, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("json.Unmarshal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func mustOpen(file string) io.Reader {
	r, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	return r
}
