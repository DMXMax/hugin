package hugindb 

import (
	"fmt"
	_ "errors"
	_"log"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)
func Init() error{
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/HUGIN", "Raven", os.Getenv("NOPPA_DB")))
	return err
}
