package moqt

type Owner string

const (
	OwnerLocal  Owner = "local"
	OwnerRemote Owner = "remote"
)

func GetOwner(o Owner) *Owner {
	return &o
}
