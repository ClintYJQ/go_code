package utils

import (
	"github.com/bwmarrin/snowflake"
)

// GenUniqueID 产生唯一ID
func GenUniqueID() string {
	node, _ := snowflake.NewNode(1)
	id := node.Generate()
	return id.Base36()
}
