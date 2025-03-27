package domain

type Area struct {
	total float64 // 계약 면적, ㎡2 단위
	net   float64 // 전용 면적, ㎡2 단위
}

func NewArea(total float64, net float64) *Area {
	return &Area{
		total: total,
		net:   net,
	}
}
