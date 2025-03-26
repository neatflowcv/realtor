package zigbang

// generated by https://mholt.github.io/json-to-go/

type CatalogList struct {
	LocalCode string  `json:"localCode,omitempty"`
	Count     int     `json:"count,omitempty"`
	List      []List  `json:"list,omitempty"`
	Local1    string  `json:"local1,omitempty"`
	Local2    string  `json:"local2,omitempty"`
	Local3    string  `json:"local3,omitempty"`
	Step      string  `json:"step,omitempty"`
	Lat       float64 `json:"lat,omitempty"`
	Lng       float64 `json:"lng,omitempty"`
}

type RoomTypeTitle struct {
	M2 string `json:"m2,omitempty"`
	P  string `json:"p,omitempty"`
}

type ItemIDList struct {
	ItemSource string `json:"itemSource,omitempty"`
	ItemID     int    `json:"itemId,omitempty"`
}

type List struct {
	AreaHoID             int           `json:"areaHoId,omitempty"`
	TranType             string        `json:"tranType,omitempty"`
	AreaDanjiID          int           `json:"areaDanjiId,omitempty"`
	AreaDanjiName        string        `json:"areaDanjiName,omitempty"`
	DanjiRoomTypeID      int           `json:"danjiRoomTypeId,omitempty"`
	Local2               string        `json:"local2,omitempty"`
	Local3               string        `json:"local3,omitempty"`
	IsPriceRange         bool          `json:"isPriceRange,omitempty"`
	DepositMin           int           `json:"depositMin,omitempty"`
	RentMin              int           `json:"rentMin,omitempty"`
	RoomTypeTitle        RoomTypeTitle `json:"roomTypeTitle,omitempty"`
	SizeContractM2       float64       `json:"sizeContractM2,omitempty"`
	SizeM2               float64       `json:"sizeM2,omitempty"`
	Dong                 string        `json:"dong,omitempty"`
	Floor                string        `json:"floor,omitempty"`
	ItemTitle            string        `json:"itemTitle,omitempty"`
	ZzimCount            int           `json:"zzimCount,omitempty"`
	IsZzim               bool          `json:"isZzim,omitempty"`
	IsActualItemChecked  bool          `json:"isActualItemChecked,omitempty"`
	ThumbnailURL         string        `json:"thumbnailUrl,omitempty"`
	ItemCount            int           `json:"itemCount,omitempty"`
	ItemIDList           []ItemIDList  `json:"itemIdList,omitempty"`
	AgentThumbnailUrls   []string      `json:"agentThumbnailUrls,omitempty"`
	IsFloorPlanThumbnail bool          `json:"isFloorPlanThumbnail,omitempty"`
}
