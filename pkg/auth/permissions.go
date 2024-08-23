package auth

type PermissionLevel int

const (
	Public PermissionLevel = iota
	Authenticated
	Admin
	// 可以根据需要添加更多级别
)
