//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type RolesPermissions struct {
	RoleID       int64 `sql:"primary_key"`
	PermissionID int64 `sql:"primary_key"`
	Level        int16
}
