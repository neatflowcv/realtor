package domain

type SourceKind string

const (
	SourceKindZigbang SourceKind = "zigbang"
)

type RealtySource struct {
	kind SourceKind
	id   string
}

func NewRealtySource(kind SourceKind, id string) *RealtySource {
	return &RealtySource{
		kind: kind,
		id:   id,
	}
}
