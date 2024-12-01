package events

type Deserializable interface {
	Deserialize() ([]byte, error)
}
