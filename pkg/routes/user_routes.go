package routes

import (
	"encoding/json"
	"my_fin/backend/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ar *AppRouter) UserCategories(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	uCat, uInCat := ar.categoryRepo.LoadCategories(userID)
	c.JSON(http.StatusOK, gin.H{
		"ok":                  true,
		"categories":          uCat,
		"categories_incoming": uInCat,
		"assets":              ar.assetsRepo.GetPossibleAssets(),
	})
}

func (ar *AppRouter) UpdateUserCategories(c *gin.Context) {
	var p struct {
		Cat         []repository.Category `json:"cat"`
		CatIncoming []repository.Category `json:"cat_in"`
	}
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	userID := ar.getUserIDFromRequest(c)
	tableKey := "categories"
	cat := &p.Cat
	if p.Cat == nil {
		tableKey = "categories_incoming"
		cat = &p.CatIncoming
	}
	if ar.categoryRepo.UpdateCategories(userID, cat, tableKey) {
		uCat, uInCat := ar.categoryRepo.LoadCategories(userID)
		c.JSON(http.StatusOK, gin.H{"ok": true, "categories": uCat, "categories_incoming": uInCat})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
	}
}
