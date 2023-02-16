package serializer

type Serializable interface {
	Serialize() interface{}
}
