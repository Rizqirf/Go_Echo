package seeds

import (
	model "github.com/Rizqirf/go_echo/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)


func PopulateTable(db *gorm.DB, Title string, Description string,Image_Banner string,Start_Time datatypes.Date,End_Time datatypes.Date) error {
	return db.Create(&model.PropertyPromotion{Title:Title, Description:Description , Image_Banner: Image_Banner, Start_Time: Start_Time, End_Time: End_Time}).Error
}

