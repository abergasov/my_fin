package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_fin/src/repository"
	"net/http"
)

func (ar *AppRouter) UserCategories(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	uCat, uInCat := ar.categoryRepository.LoadCategories(userId)
	c.JSON(http.StatusOK, gin.H{"ok": true, "categories": uCat, "categories_incoming": uInCat})
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

	userId := ar.getUserIdFromRequest(c)
	tableKey := "categories"
	cat := &p.Cat
	if p.Cat == nil {
		tableKey = "categories_incoming"
		cat = &p.CatIncoming
	}
	if ar.categoryRepository.UpdateCategories(userId, cat, tableKey) {
		uCat, uInCat := ar.categoryRepository.LoadCategories(userId)
		c.JSON(http.StatusOK, gin.H{"ok": true, "categories": uCat, "categories_incoming": uInCat})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
	}

}
