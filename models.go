package achieve_db_go

type User struct {
	_msgpack       struct{} `msgpack:",asArray"`
	Email          string
	Username       string
	PasswordHash  string
	Groups         []string
	EmailVerified bool
}

type Connection struct {
	_msgpack       struct{} `msgpack:",asArray"`
	Email          string
	Username       string
	PasswordHash  string
	Groups         []string
	EmailVerified bool
}
