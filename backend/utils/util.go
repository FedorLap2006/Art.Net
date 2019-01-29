package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	sio "github.com/graarh/golang-socketio"
	// _ "github.com/lib/pq"
)

var Database *sql.DB
var SIOServer *sio.Server
