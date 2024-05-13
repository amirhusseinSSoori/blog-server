package migration

import (
	articleModle "blog/internal/modules/article/models"
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"
)

func Migrate() {

	db := database.Connection()

	err := db.AutoMigrate(&userModel.User{}, &articleModle.Article{})

	if err != nil {
		log.Fatal("cant Migration")
		return
	}

	fmt.Println("Migration done ...")
}
