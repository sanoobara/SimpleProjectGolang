package SQL

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
)

func createstringAge() string {
	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(30)
	var s string
	for i := 1; i < m; i++ {
		c := 'a' + rune(rand.Intn('z'-'a'+1))
		s += string(c)
	}
	fmt.Println(s)
	return s
}

func Exect() {

	connStr := "user=postgres password=programmer2013 dbname=accounting sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(100)

	name := createstringAge()
	result, err := db.Exec("insert into users (Name, Age) values ($1, $2)",
		name, m)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected()) // количество добавленных строк

}
