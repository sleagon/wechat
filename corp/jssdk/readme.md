### 获取 jsapi_ticket 示例
```Go
package main

import (
	"fmt"

	"github.com/sleagon/wechat/corp"
	"github.com/sleagon/wechat/corp/jssdk"
)

var TokenServer = corp.NewDefaultTokenServer("corpId", "corpSecret", nil)
var TicketServer = jssdk.NewDefaultTicketServer(TokenServer, nil)

func main() {
	fmt.Println(TicketServer.Ticket())
}
```