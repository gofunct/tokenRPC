/*
type TokenServer interface {
	Allowance(context.Context, *contracts.AllowanceReq) (*contracts.AllowanceResp, error)
	Approve(context.Context, *contracts.ApproveReq) (*contracts.TransactionResp, error)
	ApproveAndCall(context.Context, *contracts.ApproveAndCallReq) (*contracts.TransactionResp, error)
	BalanceOf(context.Context, *contracts.BalanceOfReq) (*contracts.BalanceOfResp, error)
	Burn(context.Context, *contracts.BurnReq) (*contracts.TransactionResp, error)
	BurnFrom(context.Context, *contracts.BurnFromReq) (*contracts.TransactionResp, error)
	Decimals(context.Context, *contracts.Empty) (*contracts.DecimalsResp, error)
	Name(context.Context, *contracts.Empty) (*contracts.NameResp, error)
	Symbol(context.Context, *contracts.Empty) (*contracts.SymbolResp, error)
	TotalSupply(context.Context, *contracts.Empty) (*contracts.TotalSupplyResp, error)
	Transfer(context.Context, *contracts.TransferReq) (*contracts.TransactionResp, error)
	TransferFrom(context.Context, *contracts.TransferFromReq) (*contracts.TransactionResp, error)
	OnApproval(context.Context, *contracts.OnApprovalReq) (*contracts.TransactionResp, error)
	OnBurn(context.Context, *contracts.OnBurnReq) (*contracts.TransactionResp, error)
	OnTransfer(context.Context, *contracts.OnTransferReq) (*contracts.TransactionResp, error)
}
 */

package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type account struct {
	conn *ethclient.Client

	privateKey *ecdsa.PrivateKey
	auth       *bind.TransactOpts
	nonce      uint64
}

func New(conn *ethclient.Client, key string) account {
	acc := account{
		conn: conn,
	}

	// load private key
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(key))
	if err != nil {
		fmt.Printf("Failed to get private key: %v\n", err)
		os.Exit(-1)
	}
	acc.privateKey = privateKey
	acc.auth = bind.NewKeyedTransactor(privateKey)
	acc.auth.GasLimit = uint64(4712388)

	// get nonce
	nonce, err := conn.NonceAt(context.Background(), acc.auth.From, nil)
	if err != nil {
		fmt.Printf("Failed to get nonce: %v\n", err)
		os.Exit(-1)
	}
	acc.nonce = nonce
	acc.auth.GasPrice = big.NewInt(20000000000)
	fmt.Printf("New account (address, nonce) = (%v, %v)\n", acc.auth.From.Hex(), nonce)
	return acc
}

// transactOpts retrun TransactOpts with latest nonce and increase nonce
func (acc *account) TransactOpts() *bind.TransactOpts {
	acc.auth.Nonce = big.NewInt(int64(acc.nonce))
	acc.increaseNonce()
	return acc.auth
}

func (acc *account) increaseNonce() {
	acc.nonce++
}
