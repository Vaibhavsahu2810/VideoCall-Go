package handlers

import (
	"fmt"
	"os"
	"time"
	"v/pkg/chat"
	w "v/pkg/webrtc"

	"crypto/sha256"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	guuid "github.com/google/uuid"
	"github.com/pion/webrtc/v3"
	"golang.org/x/net/websocket"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s",gguid.New().String()))
}

func RoomWebsocket(c *websocket.Conn){
	uuid := c.Params("uuid")

	if uuid == ""{
		return
	}

	createOrGetRoom(uuid)
}

func Room(c *fiber.Ctx){
	uuid := c.Params("uuid")

	if uuid == ""{
		return
	}

	createOrGetRoom(uuid)

}