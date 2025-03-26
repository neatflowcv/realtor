package domain

type TransactionKind string

const (
	TransactionKindTrade  TransactionKind = "trade"  // 매매
	TransactionKindJeonse TransactionKind = "jeonse" // 전세
	TransactionKindRent   TransactionKind = "rent"   // 임대
)

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

type Realty struct {
	source      *RealtySource
	transaction TransactionKind
}

func NewRealty(source *RealtySource, transaction TransactionKind) *Realty {
	return &Realty{
		source:      source,
		transaction: transaction,
	}
}

func (r *Realty) SourceID() string {
	return r.source.id
}

func (r *Realty) TransactionKind() TransactionKind {
	return r.transaction
}

func (r *Realty) SourceKind() SourceKind {
	return r.source.kind
}
