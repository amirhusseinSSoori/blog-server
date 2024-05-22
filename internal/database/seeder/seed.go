package seeder

import (
	userModel "blog/internal/modules/user/models"

	aticleModel "blog/internal/modules/article/models"
	"blog/pkg/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("Mxy283rt@"), 12)
	if err != nil {
		log.Fatal("hash Password error")
		return
	}

	user := userModel.User{Name: "AmirHusseinSoori", Email: "amirhusseinsoori75@gmail.com", Password: string(hashPassword)}

	db.Create(&user)

	log.Printf("User Created succesfully with eamil address % \n", user.Email)

	// for aticle

	aticle := aticleModel.Article{Title: "Title 2", Content: "Content 2", UserID: 2}

	db.Create(&aticle)

	log.Printf("aticle Created succesfully with Title % \n", aticle.Title)
}
