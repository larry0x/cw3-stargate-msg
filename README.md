# cw3-stargate-msg

In this example we show how a cw3 multisig can give an EoA the grant to send MsgStoreCode.

First, edit the following values in the script:

```go
var (
	// the chain's bech32 address prefix
	addressPrefix = "osmo"

	// the account that will give the grants
	// in our case this is the cw3 contract address
	granter = "osmo1..."

	// the account that will be granted the ability to store code
	// in our case this should be an EOA
	grantee = "osmo1..."

	// when the grants expire
	expiration = time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
)
```

Then, run the script:

```bash
go run main.go
```

It should output the following:

```json
{
  "stargate": {
    "type_url": "/cosmos.authz.v1beta1.MsgGrant",
    "value": "..."
  }
}
```

Where `"..."` is a base64-encoded message data.

Execute this with the cw3 (e.g. using the Apollo Safe webapp).
