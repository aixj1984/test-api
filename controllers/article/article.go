package article

import (
	"log"
	"strconv"

	"test-api/models"

	"test-api/providers"

	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var articles []*models.Article

	_, err := providers.Article.GetMore(&articles, "", "1", (start-1)*length, length)

	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"code":  101,
			"msg":   err.Error(),
			"count": 0,
			"data":  nil,
		})
		return
	}

	count, err := providers.Article.Count("", "1")

	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"code":  101,
			"msg":   err.Error(),
			"count": 0,
			"data":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "",
		"count": count,
		"data":  articles,
	})

}

func GetArticleDetail(c *gin.Context) {

	articleid, _ := strconv.Atoi(c.DefaultQuery("articleid", "1"))

	var article models.Article
	article.Id = articleid

	err := providers.Article.GetOne(&article)

	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"code":  101,
			"msg":   err.Error(),
			"count": 0,
			"data":  nil,
		})
		return
	}

	providers.Article.UpdateReader(articleid)

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"data": article,
	})

}
