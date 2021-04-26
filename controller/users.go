package controller

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users  []structs.User
		result gin.H
	)

	idb.DB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  len(users),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	name := c.PostForm("name")
	email := c.PostForm("email")
	hobby := c.PostForm("hobby")
	user.Name = name
	user.Email = email
	user.Hobby = hobby
	idb.DB.Create(&user)
	result = gin.H{
		"result": user,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateUser(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	hobby := c.PostForm("hobby")
	var (
		user    structs.User
		newUser structs.User
		result  gin.H
	)

	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newUser.Name = name
	newUser.Email = email
	newUser.Hobby = hobby

	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "succesfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&user).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
