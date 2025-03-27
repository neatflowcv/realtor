package domain

type RealtyBuilder struct {
	source         *RealtySource
	transaction    TransactionKind
	deposit        int
	rent           int
	maintenanceFee int
	area           *Area
	address        string
	coordinates    *Coordinates
}

func NewRealtyBuilder(source *RealtySource, transaction TransactionKind) *RealtyBuilder {
	return &RealtyBuilder{
		source:         source,
		transaction:    transaction,
		deposit:        0,
		rent:           0,
		maintenanceFee: 0,
		area:           nil,
		address:        "",
		coordinates:    nil,
	}
}

func (b *RealtyBuilder) Deposit(deposit int) *RealtyBuilder {
	b.deposit = deposit
	return b
}

func (b *RealtyBuilder) Rent(rent int) *RealtyBuilder {
	b.rent = rent
	return b
}

func (b *RealtyBuilder) MaintenanceFee(maintenanceFee int) *RealtyBuilder {
	b.maintenanceFee = maintenanceFee
	return b
}

func (b *RealtyBuilder) Area(area *Area) *RealtyBuilder {
	b.area = area
	return b
}

func (b *RealtyBuilder) Address(address string) *RealtyBuilder {
	b.address = address
	return b
}

func (b *RealtyBuilder) Coordinates(coordinates *Coordinates) *RealtyBuilder {
	b.coordinates = coordinates
	return b
}

func (b *RealtyBuilder) Build() *Realty {
	return &Realty{
		source:         b.source,
		transaction:    b.transaction,
		deposit:        b.deposit,
		rent:           b.rent,
		maintenanceFee: b.maintenanceFee,
		area:           b.area,
		address:        b.address,
		coordinates:    b.coordinates,
	}
}
