# [WIP] httua

```golang
// memoryconn - Memory set/get example:
import "github.com/deeper-x/httua/lib/memoryconn"

conn := memoryconn.Connector{Client: memoryconn.NewClient()}
conn.SetValue("key", "value")
fmt.Println(conn.GetValue("key"))
```
