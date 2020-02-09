package main

import(
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)
func main() {
	db, err := sql.Open("postgres", "postgres");
	fmt.Println("数据库加载成功。")
	fmt.Println(db,err)
}
