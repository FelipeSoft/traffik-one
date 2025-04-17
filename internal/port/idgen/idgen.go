package idgen

import (
    "log"
    "sync"

    "github.com/bwmarrin/snowflake"
)

var (
    once sync.Once
    node *snowflake.Node
)

func InitNode(nodeID int64) {
    once.Do(func() {
        var err error
        node, err = snowflake.NewNode(nodeID)
        if err != nil {
            log.Fatalf("Failed to initialize Snowflake node: %v", err)
        }
    })
}

func GenerateID() string {
    return node.Generate().String()
}
