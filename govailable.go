package govailable

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Appointment struct {
	ID               int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	start_time       time.Time
	end_time         time.Time
	service_provided []ServiceProvided
	client_id        int  `sql:`
	nurse_id         int  `sql:`
	deleted          bool `sql:"DEFAULT:false"`
}

type ServiceProvided struct {
	ID    int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Price int
}

// globals vars
var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:bjorn@/availability?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Appointment{}, &ServiceProvided{})

	r := gin.Default()
	r.GET("/nurse-appointments/:id", GetAppointmentsNurse)
	r.GET("/client-appointments/:id", GetAppointmentsClient)
	r.PUT("/cancel/:id", DeleteAppointment)
	r.GET("/appointments/:nurse")
	r.Run(":9000")
}

func DeleteAppointment(c *gin.Context) {
	// put delete to false
}

func GetAppointmentsClient(c *gin.Context) {
	id := c.Params.ByName("id")
	var appointments []Appointment
	if err := db.Where("client_id = ? AND deleted=false", id).Find(&appointments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, appointments)
	}
}

func GetAppointmentsNurse(c *gin.Context) {
	id := c.Params.ByName("id")
	var appointments []Appointment
	if err := db.Where("nurse_id =? AND deleted=false", id).Find(&appointments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, appointments)
	}
}

