package repository

import "my_fin/backend/pkg/database"

type AssetsRepository struct {
	db *database.DBAdapter
}

type AssetType struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

var bankDeposit = AssetType{ID: 1, Title: "bank_deposit"}
var cryptoAsset = AssetType{ID: 2, Title: "crypto_asset"}
var actions = AssetType{ID: 3, Title: "asset_actions"}

var availableAssets []AssetType

func InitAssetsRepository(db *database.DBAdapter) *AssetsRepository {
	availableAssets = []AssetType{
		bankDeposit, cryptoAsset, actions,
	}
	return &AssetsRepository{db: db}
}

func (ar *AssetsRepository) GetPossibleAssets() (res []Category) {
	for _, v := range availableAssets {
		res = append(res, Category{Title: v.Title, ID: v.ID})
	}
	return res
}
