# TokensApi
Tokens.NET API still in beta

Private repo, use ```git config --global url."git@github.com:".insteadOf "https://github.com/"``` to be
able to use go get.

Example

```
import (
    api "github.com/fiksn/TokensApi"
    entities "github.com/fiksn/TokensApi/entities"
)

api.Init("./credentials)

resp, err := GetTicker("btcusdt, DAY)
```

Unofficial information: public requests are not intentionally rate-limited, private ones are limited to 300 requests/minute and then
(at least) 1 min firewall is applied
