package timer

import "time"

func get1s() time.Duration {
	return time.Duration(int64(time.Now().UnixNano() / int64(time.Second)))
}
