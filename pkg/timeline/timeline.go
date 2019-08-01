package timeline

import (
	"errors"
	"fmt"
	"time"
)
import "github.com/google/uuid"

type Operation int

const (
	ADD    Operation = 0
	MODIFY Operation = 1
)

type Node struct {
	Author         string            `json:"Author"`
	Timestamp      string            `json:"time"`
	Op             Operation         `json:"op"`
	PrevTxId       string            `json:"prev"`
	Signature      string            `json:"sig"`
	PermissionList []PermissionEntry `json:"permissions"`
	TxId           string            `json:"TxId"`
	// Add a struct for file details (should be defined in filesystem pkg)
}

func AddFile(author string, path string) (n Node, err error) {

	// TODO: Add file to file server

	// Create new node
	n = Node{}
	n.Author = author

	n.Timestamp = time.Now().String()
	n.Op = ADD

	// Default permissions - Full control to the Author
	p1 := PermissionEntry{author, false, FULL_CONTROL}
	n.PermissionList = append(n.PermissionList, p1)

	// TODO: Compute Signature
	// TODO: Send Transaction with Node
	uuid, err := uuid.NewUUID()
	if err != nil {
		return
	}
	n.TxId = uuid.String()

	return
}

func ModifyFile(author string, prevNode Node) (newNode Node, err error) {
	// TODO: Check if file exists in the File System

	// Check if User allowed to modify file
	perm := getPermissions(author, prevNode.PermissionList)
	fmt.Println(prevNode)
	fmt.Println(perm)
	if uint(perm) < uint(READ_AND_MODIFY) {
		err = errors.New("Not enough permissions to modify")
		return
	}

	// Create new node
	newNode = Node{}
	newNode.Author = author
	newNode.Timestamp = time.Now().String()
	newNode.Op = MODIFY
	newNode.PrevTxId = prevNode.TxId

	// TODO: Compute Signature
	// TODO: Send Transaction
	uuid, err := uuid.NewUUID()
	if err != nil {
		return
	}
	newNode.TxId = uuid.String()

	return
}
