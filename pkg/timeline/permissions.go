package timeline

import "math"

type Permission uint

const (
	NONE            Permission = 0
	READ_ONLY       Permission = 1
	READ_AND_MODIFY Permission = 2
	FULL_CONTROL    Permission = 3 // Includes modifying permissions
)

type PermissionEntry struct {
	User       string     `json:"User"`
	Any        bool       `json:"Any"` // keyword for Any User
	Permission Permission `json:"Permission"`
}

func getPermissions(user string, permissions []PermissionEntry) (p Permission) {
	p = NONE
	for _, permissionEntry := range permissions {
		if permissionEntry.Any || permissionEntry.User == user {
			p = Permission(math.Max(float64(permissionEntry.Permission), float64(p)))
		}
	}
	return
}
