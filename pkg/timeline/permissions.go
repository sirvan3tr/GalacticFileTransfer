package timeline

type Permission uint

const (
	READ_FILE               Permission = 0
	MODIFY_FILE             Permission = 1
	MODIFY_FILE_PERMISSIONS Permission = 2
)

type PermissionEntry struct {
	user       string
	any        bool // keyword for any user
	permission Permission
}
