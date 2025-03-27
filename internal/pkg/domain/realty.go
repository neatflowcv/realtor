package domain

type TransactionKind string

const (
	TransactionKindTrade  TransactionKind = "trade"  // 매매
	TransactionKindJeonse TransactionKind = "jeonse" // 전세
	TransactionKindRent   TransactionKind = "rent"   // 임대
)

type Realty struct {
	source         *RealtySource
	transaction    TransactionKind
	deposit        int          // 만원 단위
	rent           int          // 만원 단위, 비어있을 수 있음
	maintenanceFee int          // 만원 단위, 비어있을 수 있음
	area           *Area        // 비어있을 수 있음
	address        string       // 비어있을 수 있음
	coordinates    *Coordinates // 비어있을 수 있음
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

func (r *Realty) Deposit() int {
	return r.deposit
}

func (r *Realty) Rent() int {
	return r.rent
}

func (r *Realty) TotalArea() float64 {
	return r.area.total
}

func (r *Realty) NetArea() float64 {
	return r.area.net
}
