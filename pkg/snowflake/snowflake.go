package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) {
	var st time.Time
	st, _ = time.Parse("2006-01-02", startTime)
	sf.Epoch = st.UnixNano() / 1000000
	node, _ = sf.NewNode(machineID)
}

func GenID() int64 {
	return node.Generate().Int64()
}
