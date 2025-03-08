package entity

type RoleAction = int

const (
	RoleActionRead   RoleAction = iota // 0
	RoleActionUpdate RoleAction = iota
	RoleActionCreate RoleAction = iota
	RoleActionDelete RoleAction = iota
)
