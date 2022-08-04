package main

import (
	"fmt"
	"time"
	"unsafe"
)

type Post struct {
	IsDraft     bool      // 1 byte
	Title       string    // 16 bytes
	Id          int64     // 8 bytes
	Description string    // 16 bytes
	IsDeleted   bool      // 1 byte
	Author      string    // 16 bytes
	CreatedAt   time.Time // 24 bytes
}

type SortedPost struct {
	IsDraft     bool      // 1 byte
	IsDeleted   bool      // 1 byte
	Id          int64     // 8 bytes
	Title       string    // 16 bytes
	Description string    // 16 bytes
	Author      string    // 16 bytes
	CreatedAt   time.Time // 24 bytes
}

func main() {
	post := Post{}
	sortedPost := SortedPost{}

	fmt.Printf("Memory: %v\n", unsafe.Sizeof(post))
	fmt.Printf("Memory: %v\n", unsafe.Sizeof(sortedPost))
}
