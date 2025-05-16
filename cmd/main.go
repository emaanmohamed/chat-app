package main

import (
	"fmt"
	"github.com/emaanmohamed/chat-app/internal/db"
	"gorm.io/gorm"
)

func main() {

	db.ConnectToDB()

}
