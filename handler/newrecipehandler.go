package handler

import (
	"encoding/json"
	"gin-api/model"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []model.Recipe

func init() {
	recipes = make([]model.Recipe, 0)
	// 起動時にデータ注入
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe model.Recipe
	// リクエストデータを取り出す
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// ユニークなIDを生成
	recipe.ID = xid.New().String()
	// 現在時刻を追加
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}