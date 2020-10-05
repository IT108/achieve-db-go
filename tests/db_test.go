package achieve_db_go_test

import (
	db "../../achieve-db-go"
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	db.ConfigureFromEnv()
	log.Println(db.Select("auth", "primary", "a@a.com").Data)
	log.Println(db.Select("auth", "secondary", "test1"))
	log.Println(db.SelectUser("auth", "secondary", "ttest"))
	log.Println(db.Insert("auth", []interface{}{"test1@test1.ru", "test1", "passhash", []string{}, false}))
	log.Println(db.Delete("auth", "primary", "test1@test1.ru"))
}
