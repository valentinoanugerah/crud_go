package controller

import(
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/valentinoanugerah/crud_go/database"
	"github.com/valentinoanugerah/crud_go/models"

)

//definisi produk input

type ProductInput struct{
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock int `json:"stock" binding:"required"`
}

func GetProduct(c *gin.Context){
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if result := database.DB.First(&product, id); result.RowsAffected == 0{
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)

}

func CreateProduct(c *gin.Context){
	var input ProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		Stock: input.Stock,
	}
	
	result := database.DB.Create(&product)
	if result.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return

	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context){
	var product models.Product

	id := c.Param("id")
	result := database.DB.First(&product, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return

	}

	var input ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock

	if result := database.DB.Save(&product); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }


	c.JSON(http.StatusOK, product)
}


func DeleteProduct(c *gin.Context){
	var product models.Product
	id := c.Param("id")
	result := database.DB.First(&product, id)
	if result.RowsAffected == 0{
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	database.DB.Delete(&product)
		c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}