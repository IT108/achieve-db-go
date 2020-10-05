package achieve_db_go

import "os"

var tarantoolUser = "internal"
var tarantoolPassword = "127.0.0.1:3301"
var tarantoolHost = ""

func set(variable *string, data string) {
	if data != "" {
		*variable = data
	}
}

func Setup(user string, password string, host string)  {
	set(&tarantoolUser, user)
	set(&tarantoolPassword, password)
	set(&tarantoolHost, host)
}

func ConfigureFromEnv() {
	set(&tarantoolUser, os.Getenv("tarantool_user"))
	set(&tarantoolPassword, os.Getenv("tarantool_password"))
	set(&tarantoolHost, os.Getenv("tarantool_host"))
}