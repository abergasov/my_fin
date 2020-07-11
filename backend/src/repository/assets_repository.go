package repository

import "my_fin/src/data_provider"

type AssetsRepository struct {
	db *data_provider.DBAdapter
}

type AssetType struct {
	Title string `json:"title"`
	Id    int64  `json:"title"`
}

var bankDeposit = AssetType{Id: 1, Title: "bank_deposit"}
var cryptoAsset = AssetType{Id: 2, Title: "crypto_asset"}
var actions = AssetType{Id: 3, Title: "asset_actions"}

var availableAssets []AssetType

const deposit = "bank_deposit"

func InitAssetsRepository(db *data_provider.DBAdapter) *AssetsRepository {
	availableAssets = []AssetType{
		bankDeposit, cryptoAsset, actions,
	}
	return &AssetsRepository{db: db}
}

func (ar *AssetsRepository) GetPossibleAssets() (res []Category) {
	for _, v := range availableAssets {
		res = append(res, Category{Title: v.Title, Id: v.Id})
	}
	return res
}
