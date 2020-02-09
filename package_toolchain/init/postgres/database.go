package postgres

import "database/sql"

func init(){
	//创建一个数据库驱动实例
	sql.Register("postgres",new(PostgresDriver));
}
