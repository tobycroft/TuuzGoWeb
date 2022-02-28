package db_conf

func Dsn() string {
	dbname := "1"
	dbuser := "1"
	dbpass := "1"
	dbhost := "10.0.0.170"
	conntype := "tcp"
	dbport := "3306"
	charset := "utf8mb4"
	return dbuser + ":" + dbpass + "@" + conntype + "(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
}
