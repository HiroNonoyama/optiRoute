package main

import (
  "github.com/gin-gonic/gin"
  "controllers"
  "net/http"
  "strconv"
)

func main() {
  r := gin.Default()

  r.GET("/place/:place", func(c *gin.Context) {
    place := c.Param("place")
    result := controllers.QueryAutoComplete(place)
    if result == nil {
      c.JSON(404, gin.H{})
      return
    }
    c.JSON(http.StatusOK, gin.H{"predictions": result})
  })

// パラメータの形=> origin=横浜市青葉区新石川4-3&destination=東京&waypoints=青葉台|東神奈川|埼玉
  r.GET("/route/:route", func(c *gin.Context) {
    route := c.Param("route")
    result := controllers.SearchOptiRoute(route)
    if result == nil {
      c.JSON(404, gin.H{})
      return
    }
    c.JSON(http.StatusOK, gin.H{"routes": result})
  })

  r.POST("/join", func(c *gin.Context) {
    ctrl := controllers.NewUser()
    age, ageErr := strconv.Atoi(c.PostForm("age"))
    sex, sexErr := strconv.Atoi(c.PostForm("sex"))
    if ageErr != nil || sexErr != nil {
      return
    }
    result := ctrl.SignUp(age, sex)
    if result == nil {
      c.JSON(404, gin.H{})
      return
    }
    c.JSON(http.StatusOK, gin.H{"user": result})
  })
  
  r.Run(":8080")
}
