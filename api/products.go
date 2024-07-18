package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-electronic-component/model"
	"github.com/mahdi-cpp/api-go-electronic-component/repository"
)

func AddProductsRoutes(rg *gin.RouterGroup) {

	router := rg.Group("/products")

	router.POST("/addProduct", func(c *gin.Context) {
		product := model.Products{}
		s := c.Query("data")
		fmt.Println(s)

		err := PrettyJson(s, &product)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, err.Error())
			return
		}

		//product.English = strings.ToLower(word.English)
		fmt.Println("addProduct")
		//fmt.Println("word.Hashtags: ", word.Hashtags)

		err = repository.AddProduct(product)
		if err != nil {
			result := "" + err.Error()
			c.JSON(404, result)
		} else {
			c.JSON(200, "success")
		}
	})

	router.GET("/GetProducts", func(c *gin.Context) {
		products, err := repository.GetProducts()
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(417, err.Error())
			return
		}
		c.JSON(200, products)
	})

}

func PrettyJson(param string, T interface{}) error {
	err := json.Unmarshal([]byte(param), T)

	if err != nil {
		fmt.Println(err)
		return err
	}
	b, err := json.MarshalIndent(T, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	return nil
}
