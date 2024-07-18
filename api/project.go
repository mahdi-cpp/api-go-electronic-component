package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-electronic-component/model"
	"github.com/mahdi-cpp/api-go-electronic-component/repository"
)

func addProjectRoutes(rg *gin.RouterGroup) {

	router := rg.Group("/project")

	router.POST("/addProject", func(c *gin.Context) {
		project := model.Project{}
		s := c.Query("data")
		fmt.Println(s)

		err := PrettyJson(s, &project)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, err.Error())
			return
		}

		err = repository.AddProject(project)
		if err != nil {
			result := "" + err.Error()
			c.JSON(404, result)
		} else {
			c.JSON(200, "success")
		}
	})

	router.GET("/GetById", func(c *gin.Context) {
		id := c.Query("id")
		project, err := repository.GetProject(id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(417, err.Error())
			return
		}
		c.JSON(200, project)
	})

	router.GET("/GetByCategory", func(c *gin.Context) {
		category := c.Query("category")
		projects, err := repository.GetProjectsByCategory(category)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(417, err.Error())
			return
		}
		c.JSON(200, projects)
	})

}
