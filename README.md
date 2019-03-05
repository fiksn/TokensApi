# TokensApi
A library in Go(lang) to use the API from the cryptocurrency exchange [Tokens](https://www.tokens.net)

Warning this library is provided as-is, authors are not liable for any kind of damages (like financial loss but not limited to)

Example

```
package test

import (
    api "github.com/fiksn/TokensApi"
    entities "github.com/fiksn/TokensApi/entities"
)

api.Init("./credentials)

resp, err := GetTicker("btcusdt, DAY)
```

## Rate limits

Unofficial information: public requests are not intentionally rate-limited, private ones are limited to 300 requests/minute and then
(at least) 1 min firewall is applied
