package qlog

type KeyUpdated struct {
	KeyType    KeyType
	Old        []byte
	New        []byte
	Generation uint32

	Trigger string
}

type KeyRetired struct {
	KeyType    KeyType
	Key        []byte
	Generation uint32

	Trigger string
}
