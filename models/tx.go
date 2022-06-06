package models

import (
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameTx = "sync_tx"
)

type (
	Tx struct {
		Time          int64         `bson:"time"`
		Height        int64         `bson:"height"`
		TxHash        string        `bson:"tx_hash"`
		Type          string        `bson:"type"` // parse from first msg
		Memo          string        `bson:"memo"`
		Status        uint32        `bson:"status"`
		Log           string        `bson:"log"`
		Fee           *types.Fee    `bson:"fee"`
		GasUsed       int64         `bson:"gas_used"`
		Types         []string      `bson:"types"`
		EventsNew     []EventNew    `bson:"events_new"`
		Signers       []string      `bson:"signers"`
		DocTxMsgs     []types.TxMsg `bson:"msgs"`
		Addrs         []string      `bson:"addrs"`
		ContractAddrs []string      `bson:"contract_addrs"`
		TxIndex       uint32        `bson:"tx_index"`
		Ext           interface{}   `bson:"ext"`
	}

	Event struct {
		Type       string   `bson:"type"`
		Attributes []KvPair `bson:"attributes"`
	}

	KvPair struct {
		Key   string `bson:"key"`
		Value string `bson:"value"`
	}

	EventNew struct {
		MsgIndex uint32  `bson:"msg_index" json:"msg_index"`
		Events   []Event `bson:"events"`
	}
)

func (d Tx) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameTx
	}
	return fmt.Sprintf("sync_%v_tx", GetSrvConf().ChainId)
}

func (d Tx) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-tx_hash"},
		Background: true,
	})

	indexes = append(indexes, mgo.Index{
		Key:        []string{"-height", "-tx_hash"},
		Unique:     true,
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d Tx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}
