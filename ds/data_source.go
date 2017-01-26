package ds

import (
	"database/sql"
	"github.com/mskoroglu/golaxy/config"
	"fmt"
)

var properties = config.GetProperties()

func SQL() *sql.DB {
	if properties.DataSource.Sql.Driver != "" && properties.DataSource.Sql.Url != "" {
		db, err := sql.Open(properties.DataSource.Sql.Driver, properties.DataSource.Sql.Url)
		if err != nil {
			fmt.Println(err.Error())
		}
		return db
	}
	return nil
}
