package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

const (
	machineID = 1
)

func Init() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", "2020-12-19")
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	if err != nil {
		return
	}
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
