package seed

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/bygui86/go-complete-project/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "John Doe",
		Email:    "john@doe.com",
		Password: "johnssecret",
	},
	models.User{
		Nickname: "Jane Doe",
		Email:    "jane@doe.com",
		Password: "janessecret",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "John's ideas",
		Content: "This is John's very first post!",
	},
	models.Post{
		Title:   "Jane's ideas",
		Content: "This is Jane's very first post!",
	},
}

func Load(db *gorm.DB) {
	// drop tables
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	// TODO to be clarified
	// migrate tables
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	// add foreign key to Posts table
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	// import users together with posts
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
