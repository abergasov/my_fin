package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_fin/src/repository"
	"net/http"
)

func (ar *AppRouter) UserCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true, "categories": ar.categoryRepository.LoadCategories(1)})
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

	if ar.categoryRepository.UpdateCategories(1, &p.Cat) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "categories": ar.categoryRepository.LoadCategories(1)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
	}

}
