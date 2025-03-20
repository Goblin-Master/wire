package global

import "wire/internal/utils/snowflake"

const DEFAULT_NODE_ID = 1

var Node, _ = snowflake.NewNode(DEFAULT_NODE_ID)
