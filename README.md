# [WIP] HTTUA - HTtp AUth

```golang
// memoryconn - Memory set/get example:
import "github.com/deeper-x/httua/lib/memoryconn"

conn := memoryconn.Connector{Client: memoryconn.NewClient()}
conn.SetValue("key", "value")
fmt.Println(conn.GetValue("key"))
```

```golang
// httpmanager - Http methods to GET, PUT, DELETE, POST
// TODO
import "github.com/deeper-x/httua/lib/httpmanager"
```
