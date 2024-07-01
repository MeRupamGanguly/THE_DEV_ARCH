package domain

type UserAsset struct {
	ID                 string  `json:"id" bson:"_id"`
	AssetId            string  `json:"asset_id" bson:"asset_id"`
	BuyAveragePrice    float32 `json:"buy_average_price" bson:"buy_average_price"`
	BuyTotalQuanitity  float32 `json:"buy_total_quantity" bson:"buy_total_quantity"`
	AssetPNL           float32 `json:"asset_pnl" bson:"asset_pnl"`
	IsActive           bool    `json:"is_active" bson:"is_active"`
	SellAveragePrice   float32 `json:"sell_average_price" bson:"sell_average_price"`
	SellTotalQuanitity float32 `json:"sell_total_quantity" bson:"sell_total_quantity"`
}
type Portfolio struct {
	ID       string      `json:"id" bson:"_id"`
	UserId   string      `json:"user_id" bson:"user_id"`
	PNL      float32     `json:"pnl" bson:"pnl"`
	Holdings []UserAsset `json:"holdings" bson:"holdings"`
}
