package timeline

import "time"

type Operation int

const (
	ADD Operation = 0
	MODIFY Operation = 1
)

type Node struct {
	author string
	timestamp time.Time
	operation Operation
	prevNodeHash string
	hash string
	signature string
	permissionList []PermissionEntry
	txId string
	fileId string
}


func addFile(path string) {
	// Add file to file server
	// Create new node
	// Default permissions
	// Save the node as JSON
	// Calculate Auth
	// Send Transaction
}


func modifyFile() {
	// Check if file exists in the File System
	// Check if user allowed to modify file
	// Calculate the diff
	// Create new node
	// Send transaction
	// Calculate Auth
	// Send Transaction
}
