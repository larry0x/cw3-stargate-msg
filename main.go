package main

import (
	"encoding/json"
	"fmt"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvm "github.com/CosmWasm/wasmvm/types"
)

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

func init() {
	// set address prefix so that sdk.AccAddressFromBech32 works
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(addressPrefix, addressPrefix+"pub")
	cfg.Seal()
}

func main() {
	granterAddr, err := sdk.AccAddressFromBech32(granter)
	if err != nil {
		panic(err)
	}

	granteeAddr, err := sdk.AccAddressFromBech32(grantee)
	if err != nil {
		panic(err)
	}

	msg, err := authz.NewMsgGrant(
		granterAddr,
		granteeAddr,
		authz.NewGenericAuthorization(sdk.MsgTypeURL(&wasm.MsgStoreCode{})),
		expiration,
	)
	if err != nil {
		panic(err)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		panic(err)
	}

	cosmosMsg := &wasmvm.CosmosMsg{
		Stargate: &wasmvm.StargateMsg{
			TypeURL: any.TypeUrl,
			Value:   any.Value,
		},
	}

	cosmosMsgStr, err := json.MarshalIndent(cosmosMsg, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(cosmosMsgStr))
}
