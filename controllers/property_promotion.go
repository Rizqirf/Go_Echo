package controller

import (
	"fmt"
	"net/http"

	model "github.com/Rizqirf/go_echo/models"
	"github.com/Rizqirf/go_echo/storage"
	"github.com/labstack/echo/v4"
)


func GetAll(c echo.Context) error {
	db := storage.GetDBInstance()
	promotions := []model.PropertyPromotion{}

	if err := db.Find(&promotions).Error; err != nil {
		return err
	}

	
	return c.JSON(http.StatusOK, promotions)
}


func GetById(c echo.Context) error {
	db := storage.GetDBInstance()
	promotion := model.PropertyPromotion{}
	id := c.Param("id")

	fmt.Println(id)

	if err := db.Where("Id = ?",id).First(&promotion).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusOK, promotion)
}

func Create(c echo.Context) error {
	db := storage.GetDBInstance()
	var input model.PropertyPromotionInput 
	
	if err := c.Bind(&input); err!= nil{
		return c.JSON(http.StatusBadRequest,"Bad Request")
	}

	newPropertyPromotion := model.PropertyPromotion{Title: input.Title, Description: input.Description, Image_Banner: input.Image_Banner, Start_Time: input.Start_Time, End_Time: input.End_Time}
	db.Create(&newPropertyPromotion)

	return c.JSON(http.StatusCreated, newPropertyPromotion)
}

func Update(c echo.Context) error {
	db := storage.GetDBInstance()
	id := c.Param("id")
	var input model.PropertyPromotionInput
	var promotion model.PropertyPromotion
	

	if err := db.Where("id= ?", id).First(&promotion).Error; err != nil{
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	err := c.Bind(&input); if err!= nil{
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	db.Model(&promotion).Updates(model.PropertyPromotion{Title: input.Title, Description: input.Description, Image_Banner: input.Image_Banner, Start_Time: input.Start_Time, End_Time: input.End_Time})
	return c.JSON(http.StatusOK,input)
}

func Delete(c echo.Context) error {
	db := storage.GetDBInstance()
	var promotion model.PropertyPromotion
	id := c.Param("id")

	if err := db.Where("id= ?", id).First(&promotion).Error; err != nil{
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	db.Delete(&promotion)

	return c.JSON(http.StatusOK,"Promotion Deleted")
}