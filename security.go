package qlog

type KeyUpdated struct {
	KeyType    KeyType    `json:"key_type"`
	Old        ByteString `json:"old"`
	New        ByteString `json:"new"`
	Generation uint32     `json:"generation"`

	Trigger string `json:"trigger"`
}

type KeyRetired struct {
	KeyType    KeyType    `json:"key_type"`
	Key        ByteString `json:"key"`
	Generation uint32     `json:"generation"`

	Trigger string `json:"trigger"`
}
