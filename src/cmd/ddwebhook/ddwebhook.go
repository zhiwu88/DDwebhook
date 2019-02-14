package main

import (
    "flag"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    model "model/ddmodel"
    "notifier/ddnotifier"
)

var (
    h            bool
    defaultRobot string
)

func init() {
    flag.BoolVar(&h, "h", false, "help")
    flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook")
}


func main() {

    flag.Parse()

    if h {
        flag.Usage()
        return
    }

    router := gin.Default()
    router.POST("/webhook", func(c *gin.Context) {
        var notification model.Notification
        err := c.BindJSON(&notification)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        fmt.Printf("%+v\n", notification)

        err = notifier.Send(notification, defaultRobot)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }

        c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})
    })
    router.Run()
}
