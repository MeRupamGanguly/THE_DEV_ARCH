package domain

type AssetType string

const (
	Crypto           AssetType = "Crypto"
	Stock            AssetType = "Stock"
	Index            AssetType = "Index"
	Forex            AssetType = "Forex"
	Commodity        AssetType = "Commodity"
	IndexOptions     AssetType = "IndexOptions"
	IndexFutures     AssetType = "IndexFutures"
	StockOptions     AssetType = "StockOptions"
	StockFutures     AssetType = "StockFutures"
	CryptoOptions    AssetType = "CryptoOptions"
	CryptoFutures    AssetType = "CryptoFutures"
	CommodityOptions AssetType = "CommodityOptions"
	CommodityFutures AssetType = "CommodityFutures"
)

func (at AssetType) IsValid() bool {
	switch at {
	case Crypto, Stock, Index, Forex, Commodity, IndexOptions, IndexFutures, StockOptions, StockFutures, CryptoOptions, CryptoFutures, CommodityOptions, CommodityFutures:
		return true
	default:
		return false
	}
}

type Asset struct {
	ID     string    `json:"id" bson:"_id"`
	Symbol string    `json:"symbol" bson:"symbol"`
	Type   AssetType `json:"type" bson:"type"`
}
