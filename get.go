package redis

import (
	"fmt"
	"github.com/redis-go/redcon"
	"github.com/redis-go/redis/types"
)

// TODO
func Get(c *Client, cmd redcon.Command) {
	key := string(cmd.Args[1])

	db := c.Redis().RedisDb(c.Db())
	if db.Expires(&key) {

	}

	i := db.GetOrExpired(&key, true)
	if i == nil {
		c.Conn().WriteNull()
		return
	}

	if i.ValueType() != types.StringType {
		c.Conn().WriteError(fmt.Sprintf("%s: key is a %s not a %s", WrongTypeErr, i.ValueTypeFancy(), types.StringTypeFancy))
		return
	}

	v := *i.Value().(*string)
	c.Conn().WriteBulkString(v)
}
