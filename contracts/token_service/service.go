package token_service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Service struct{
	Addr 	common.Address
	EthClient 	*ethclient.Client
}

func NewService(addr common.Address, ethClient *ethclient.Client) TokenServer {
	return &Service{
		Addr: addr,
		EthClient: ethClient,
	}
}

func (svc *Service) Allowance(context.Context, *AllowanceReq) (*AllowanceResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Approve(context.Context, *ApproveReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) ApproveAndCall(context.Context, *ApproveAndCallReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) BalanceOf(context.Context, *BalanceOfReq) (*BalanceOfResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Burn(context.Context, *BurnReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) BurnFrom(context.Context, *BurnFromReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Decimals(context.Context, *Empty) (*DecimalsResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Name(context.Context, *Empty) (*NameResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Symbol(context.Context, *Empty) (*SymbolResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) TotalSupply(context.Context, *Empty) (*TotalSupplyResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Transfer(context.Context, *TransferReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) TransferFrom(context.Context, *TransferFromReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) OnApproval(context.Context, *OnApprovalReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) OnBurn(context.Context, *OnBurnReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) OnTransfer(context.Context, *OnTransferReq) (*TransactionResp, error) {
	return nil, fmt.Errorf("not implemented")
}
