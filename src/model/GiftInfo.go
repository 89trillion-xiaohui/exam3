package model


// GiftInfo 礼包内容
type GiftInfo struct {
	GoldCoin string	`json:"gold_coin"`			//金币
	Diamond string	`json:"diamond"`			//钻石
	Props string	`json:"props"`				//道具
	Legend string	`json:"legend"`				//英雄
	Pawn string		`json:"pawn"`				//小兵
}