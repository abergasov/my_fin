package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_fin/src/repository"
	"net/http"
)

func (ar *AppRouter) UserCategories(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": true, "categories": ar.categoryRepository.LoadCategories(userId)})
}

func (ar *AppRouter) UpdateUserCategories(c *gin.Context) {
	var p struct {
		Cat []repository.Category `json:"cat"`
	}
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	userId := ar.getUserIdFromRequest(c)
	if ar.categoryRepository.UpdateCategories(userId, &p.Cat) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "categories": ar.categoryRepository.LoadCategories(userId)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
	}

}
