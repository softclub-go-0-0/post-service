package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/post-service/pkg/handlers"
	"github.com/softclub-go-0-0/post-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	DBName := flag.String("dbname", "crm_service", "Enter the name of DB")
	DBUser := flag.String("dbuser", "crm_service", "Enter the name of a DB user")
	DBPassword := flag.String("dbpassword", "crm_service", "Enter the password of user")
	DBPort := flag.String("dbport", "5432", "Enter the port of DB")
	flag.Parse()

	db, err := DBInit(*DBUser, *DBPassword, *DBName, *DBPort)
	if err != nil {
		log.Fatal("db connection error:", err)
	}

	log.Println("successfully connected to DB")

	h := handlers.NewHandler(db)
	r := gin.Default()

	r.GET("/posts", h.GetAllPosts)
	r.GET("/posts/:postID", h.GetOnePost)
	r.GET("/posts/search/:title", h.SearchPosts)

	r.Run(":4000")
}

func DBInit(user, password, dbname, port string) (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Post{},
		&models.ProcessingResult{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
