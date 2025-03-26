package naver

type ArticleList struct {
	Code        string `json:"code"`
	PaidPreSale struct {
		PreSaleComplexNumber    int     `json:"preSaleComplexNumber"`
		PreSaleComplexName      string  `json:"preSaleComplexName"`
		PreSaleAddress          string  `json:"preSaleAddress"`
		PreSaleAddressSector    string  `json:"preSaleAddressSector"`
		PreSaleDetailAddress    string  `json:"preSaleDetailAddress"`
		PreSaleStageCode        string  `json:"preSaleStageCode"`
		PreSaleStageDetails     string  `json:"preSaleStageDetails"`
		PreSaleTypeCode         string  `json:"preSaleTypeCode"`
		PreSaleFormCode         string  `json:"preSaleFormCode"`
		OccupancyYearMonth      string  `json:"occupancyYearMonth"`
		Thumbnail               string  `json:"thumbnail"`
		FeatureMarkTypeCode     string  `json:"featureMarkTypeCode"`
		MinPreSalePrice         int     `json:"minPreSalePrice"`
		MaxPreSalePrice         int64   `json:"maxPreSalePrice"`
		MinPreSaleArea          float64 `json:"minPreSaleArea"`
		MaxPreSaleArea          float64 `json:"maxPreSaleArea"`
		TotalHouseholdsNumber   int     `json:"totalHouseholdsNumber"`
		PreSaleHouseholdsNumber int     `json:"preSaleHouseholdsNumber"`
		Xcoordinate             float64 `json:"xcoordinate"`
		Ycoordinate             float64 `json:"ycoordinate"`
		PreSaleDetailsPageURL   string  `json:"preSaleDetailsPageURL"`
	} `json:"paidPreSale"`
	HasPaidPreSale bool           `json:"hasPaidPreSale"`
	More           bool           `json:"more"`
	TIME           bool           `json:"TIME"`
	Z              int            `json:"z"`
	Page           int            `json:"page"`
	Bodies         []*ArticleBody `json:"body"`
}

type ArticleBody struct {
	AtclNo              string   `json:"atclNo"`
	CortarNo            string   `json:"cortarNo"`
	AtclNm              string   `json:"atclNm"`
	AtclStatCd          string   `json:"atclStatCd"`
	RletTpCd            string   `json:"rletTpCd"`
	UprRletTpCd         string   `json:"uprRletTpCd"`
	RletTpNm            string   `json:"rletTpNm"`
	TradTpCd            string   `json:"tradTpCd"`
	TradTpNm            string   `json:"tradTpNm"`
	VrfcTpCd            string   `json:"vrfcTpCd"`
	FlrInfo             string   `json:"flrInfo"`
	Prc                 int      `json:"prc"`
	RentPrc             int      `json:"rentPrc"`
	HanPrc              string   `json:"hanPrc"`
	Spc1                string   `json:"spc1"`
	Spc2                string   `json:"spc2"`
	Direction           string   `json:"direction"`
	AtclCfmYmd          string   `json:"atclCfmYmd"`
	RepImgURL           string   `json:"repImgUrl"`
	RepImgTpCd          string   `json:"repImgTpCd"`
	RepImgThumb         string   `json:"repImgThumb"`
	Lat                 float64  `json:"lat"`
	Lng                 float64  `json:"lng"`
	AtclFetrDesc        string   `json:"atclFetrDesc,omitempty"`
	TagList             []string `json:"tagList"`
	BildNm              string   `json:"bildNm"`
	Minute              int      `json:"minute"`
	SameAddrCnt         int      `json:"sameAddrCnt"`
	SameAddrDirectCnt   int      `json:"sameAddrDirectCnt"`
	SameAddrHash        string   `json:"sameAddrHash,omitempty"`
	SameAddrMaxPrc      string   `json:"sameAddrMaxPrc,omitempty"`
	SameAddrMaxPrc2     string   `json:"sameAddrMaxPrc2,omitempty"`
	SameAddrMinPrc      string   `json:"sameAddrMinPrc,omitempty"`
	SameAddrMinPrc2     string   `json:"sameAddrMinPrc2,omitempty"`
	Cpid                string   `json:"cpid"`
	CpNm                string   `json:"cpNm"`
	CpCnt               int      `json:"cpCnt"`
	RltrNm              string   `json:"rltrNm"`
	DirectTradYn        string   `json:"directTradYn"`
	MinMviFee           int      `json:"minMviFee"`
	MaxMviFee           int      `json:"maxMviFee"`
	EtRoomCnt           int      `json:"etRoomCnt"`
	TradePriceHan       string   `json:"tradePriceHan"`
	TradeRentPrice      int      `json:"tradeRentPrice"`
	TradeCheckedByOwner bool     `json:"tradeCheckedByOwner"`
	CpLinkVO            struct {
		CpID                               string `json:"cpId"`
		MobileArticleLinkTypeCode          string `json:"mobileArticleLinkTypeCode"`
		MobileBmsInspectPassYn             string `json:"mobileBmsInspectPassYn"`
		PcArticleLinkUseAtArticleTitle     bool   `json:"pcArticleLinkUseAtArticleTitle"`
		PcArticleLinkUseAtCpName           bool   `json:"pcArticleLinkUseAtCpName"`
		MobileArticleLinkUseAtArticleTitle bool   `json:"mobileArticleLinkUseAtArticleTitle"`
		MobileArticleLinkUseAtCpName       bool   `json:"mobileArticleLinkUseAtCpName"`
	} `json:"cpLinkVO"`
	DtlAddrYn   string `json:"dtlAddrYn"`
	DtlAddr     string `json:"dtlAddr"`
	IsVrExposed bool   `json:"isVrExposed"`
	VrURL       string `json:"vrUrl"`
}
