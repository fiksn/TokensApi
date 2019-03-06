# TokensApi
An (unoffical) library in Go(lang) to use the API of the cryptocurrency exchange [Tokens](https://www.tokens.net)

Warning: this library is provided as-is, contributors are not liable for any kind of damages (like but not limited to financial loss).

I tried to learn golang by doing so there might be some ugly parts inside the code. However it follows some idiomatic patterns (like returning err object).

# Example

```
package test

import (
    api "github.com/fiksn/TokensApi"
    entities "github.com/fiksn/TokensApi/entities"
)

resp, err := GetTicker("btcusdt, DAY)
```

Before you are able to call private methods, you need to invoke:
```
api.Init("./credentials)
```
and provide a JSON file like [credentials.example](./credentials.example) with your API credentials.

Then you can do for instance:

```
balance, err = GetBalance("usdt")
```

# Hints

Try the PlaceOrderTyped() method to avoid problems with type-safety. Placing an order will _not_ make it immediately available, expect to get "131 Invalid order id" back occasionally when querying for it.

## Rate limits

Unofficial information: public requests are not intentionally rate-limited (they are cached anyway so more than 1 req/sec doesn't make sense), private ones are limited to 300 requests/minute and then (at least) 1 min firewall is applied

## Tipping

If you like it you can send some ether or abandoned kitties to 0xFF0da2B849aAbd5F37265190fFe1a64D4Febb52D ;)
