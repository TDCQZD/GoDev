package routes

import (
	"GoDev/projects/chat_gin/routes/chatroom"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LongPolling(server *gin.Engine) {
	r := server.Group("/longpolling")

	r.GET("/room", func(c *gin.Context) {
		user := c.Query("user")
		if len(user) == 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}
		chatroom.Join(user)
		c.HTML(http.StatusOK, "longpolling.html", struct {
			User string
		}{user})
	})

	r.POST("/room/messages", func(c *gin.Context) {
		user := c.PostForm("user")
		message := c.PostForm("message")
		chatroom.Send(user, message)
		c.Redirect(http.StatusMovedPermanently, "/longpolling/room")
	})

	r.GET("/room/leave", func(c *gin.Context) {
		user := c.Query("user")
		chatroom.Leave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/room/messages", func(c *gin.Context) {

		lastReceived, _ := strconv.ParseInt(c.Query("lastReceived"), 10, 64)
		subscription := chatroom.Subscribe()

		defer subscription.Cancel()

		// See if anything is new in the archive.
		var events []chatroom.Event
		for _, event := range subscription.Archive {
			e := event
			if e.Timestamp > lastReceived {
				events = append(events, e)
			}
		}

		fmt.Println(len(events))
		// If we found one, grand.
		if len(events) > 0 {
			c.JSON(http.StatusOK, events)
			return
		}

		// Else, wait for something new.
		event := <-subscription.Message

		c.JSON(http.StatusOK, []chatroom.Event{event})

		return
	})

}
