package executor

/*
coins 是一个货币的exec。内置货币的执行器。

主要提供两种操作：
EventTransfer -> 转移资产
*/

//package none execer for unknow execer
//all none transaction exec ok, execept nofee
//nofee transaction will not pack into block

import (
	"gitlab.33.cn/chain33/chain33/common/address"
	drivers "gitlab.33.cn/chain33/chain33/system/dapp"
	"gitlab.33.cn/chain33/chain33/types"
)

//var clog = log.New("module", "execs.coins")

func Init() {
	drivers.Register(GetName(), newCoins, 0)
}

func GetName() string {
	return newCoins().GetName()
}

type Coins struct {
	drivers.DriverBase
}

func newCoins() drivers.Driver {
	c := &Coins{}
	c.SetChild(c)
	return c
}

func (c *Coins) GetName() string {
	return "coins"
}

func (c *Coins) CheckTx(tx *types.Transaction, index int) error {
	return nil
}

func (c *Coins) Exec(tx *types.Transaction, index int) (*types.Receipt, error) {
	_, err := c.DriverBase.Exec(tx, index)
	if err != nil {
		return nil, err
	}
	var action types.CoinsAction
	err = types.Decode(tx.Payload, &action)
	if err != nil {
		return nil, err
	}
	coinsAccount := c.GetCoinsAccount()
	if action.Ty == types.CoinsActionTransfer && action.GetTransfer() != nil {
		transfer := action.GetTransfer()
		from := tx.From()
		//to 是 execs 合约地址
		if drivers.IsDriverAddress(tx.GetRealToAddr(), c.GetHeight()) {
			return coinsAccount.TransferToExec(from, tx.GetRealToAddr(), transfer.Amount)
		}

		return coinsAccount.Transfer(from, tx.GetRealToAddr(), transfer.Amount)
	} else if action.Ty == types.CoinsActionTransferToExec && action.GetTransferToExec() != nil {
		if c.GetHeight() < types.ForkV12TransferExec {
			return nil, types.ErrActionNotSupport
		}
		transfer := action.GetTransferToExec()
		from := tx.From()
		//to 是 execs 合约地址
		if !isExecAddrMatch(transfer.ExecName, tx.GetRealToAddr()) {
			return nil, types.ErrToAddrNotSameToExecAddr
		}
		return coinsAccount.TransferToExec(from, tx.GetRealToAddr(), transfer.Amount)
	} else if action.Ty == types.CoinsActionWithdraw && action.GetWithdraw() != nil {
		withdraw := action.GetWithdraw()
		if !types.IsMatchFork(c.GetHeight(), types.ForkV16Withdraw) {
			withdraw.ExecName = ""
		}
		from := tx.From()
		//to 是 execs 合约地址
		if drivers.IsDriverAddress(tx.GetRealToAddr(), c.GetHeight()) || isExecAddrMatch(withdraw.ExecName, tx.GetRealToAddr()) {
			return coinsAccount.TransferWithdraw(from, tx.GetRealToAddr(), withdraw.Amount)
		}
		return nil, types.ErrActionNotSupport
	} else if action.Ty == types.CoinsActionGenesis && action.GetGenesis() != nil {
		genesis := action.GetGenesis()
		if c.GetHeight() == 0 {
			if drivers.IsDriverAddress(tx.GetRealToAddr(), c.GetHeight()) {
				return coinsAccount.GenesisInitExec(genesis.ReturnAddress, genesis.Amount, tx.GetRealToAddr())
			}
			return coinsAccount.GenesisInit(tx.GetRealToAddr(), genesis.Amount)
		} else {
			return nil, types.ErrReRunGenesis
		}
	} else {
		return nil, types.ErrActionNotSupport
	}
}

func isExecAddrMatch(name string, to string) bool {
	toaddr := address.ExecAddress(name)
	return toaddr == to
}

//0: all tx
//1: from tx
//2: to tx

func (c *Coins) ExecLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := c.DriverBase.ExecLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}
	//执行成功
	var action types.CoinsAction
	err = types.Decode(tx.GetPayload(), &action)
	if err != nil {
		panic(err)
	}
	var kv *types.KeyValue
	if action.Ty == types.CoinsActionTransfer && action.GetTransfer() != nil {
		transfer := action.GetTransfer()
		kv, err = updateAddrReciver(c.GetLocalDB(), tx.GetRealToAddr(), transfer.Amount, true)
	} else if action.Ty == types.CoinsActionTransferToExec && action.GetTransferToExec() != nil {
		transfer := action.GetTransferToExec()
		kv, err = updateAddrReciver(c.GetLocalDB(), tx.GetRealToAddr(), transfer.Amount, true)
	} else if action.Ty == types.CoinsActionWithdraw && action.GetWithdraw() != nil {
		transfer := action.GetWithdraw()
		from := tx.From()
		kv, err = updateAddrReciver(c.GetLocalDB(), from, transfer.Amount, true)
	} else if action.Ty == types.CoinsActionGenesis && action.GetGenesis() != nil {
		gen := action.GetGenesis()
		kv, err = updateAddrReciver(c.GetLocalDB(), tx.GetRealToAddr(), gen.Amount, true)
	}
	if err != nil {
		return set, nil
	}
	if kv != nil {
		set.KV = append(set.KV, kv)
	}
	return set, nil
}

func (c *Coins) ExecDelLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := c.DriverBase.ExecDelLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}
	//执行成功
	var action types.CoinsAction
	err = types.Decode(tx.GetPayload(), &action)
	if err != nil {
		panic(err)
	}
	var kv *types.KeyValue
	if action.Ty == types.CoinsActionTransfer && action.GetTransfer() != nil {
		transfer := action.GetTransfer()
		kv, err = updateAddrReciver(c.GetLocalDB(), tx.GetRealToAddr(), transfer.Amount, false)
	} else if action.Ty == types.CoinsActionTransferToExec && action.GetTransferToExec() != nil {
		transfer := action.GetTransferToExec()
		kv, err = updateAddrReciver(c.GetLocalDB(), tx.GetRealToAddr(), transfer.Amount, false)
	} else if action.Ty == types.CoinsActionWithdraw && action.GetWithdraw() != nil {
		transfer := action.GetWithdraw()
		from := tx.From()
		kv, err = updateAddrReciver(c.GetLocalDB(), from, transfer.Amount, false)
	}
	if err != nil {
		return set, nil
	}
	if kv != nil {
		set.KV = append(set.KV, kv)
	}
	return set, nil
}

func (c *Coins) GetAddrReciver(addr *types.ReqAddr) (types.Message, error) {
	reciver := types.Int64{}
	db := c.GetLocalDB()
	addrReciver, err := db.Get(calcAddrKey(addr.Addr))
	if addrReciver == nil || err != nil {
		return &reciver, types.ErrEmpty
	}
	err = types.Decode(addrReciver, &reciver)
	if err != nil {
		return &reciver, err
	}
	return &reciver, nil
}

func (c *Coins) Query(funcName string, params []byte) (types.Message, error) {
	if funcName == "GetAddrReciver" {
		var in types.ReqAddr
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.GetAddrReciver(&in)
	} else if funcName == "GetTxsByAddr" {
		var in types.ReqAddr
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.GetTxsByAddr(&in)
	} else if funcName == "GetPrefixCount" {
		var in types.ReqKey
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.GetPrefixCount(&in)
	} else if funcName == "GetAddrTxsCount" {
		var in types.ReqKey
		err := types.Decode(params, &in)
		if err != nil {
			return nil, err
		}
		return c.GetAddrTxsCount(&in)
	}
	return nil, types.ErrActionNotSupport
}