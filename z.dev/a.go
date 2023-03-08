package coinprocessor

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"github.com/hyperledger/fabric/core/coinprocessor/rwsetparser"
	"github.com/hyperledger/fabric/core/common/cointransaction"
	"github.com/hyperledger/fabric/core/common/cointransaction/cointransactionfactoryimpl"
	"github.com/hyperledger/fabric/core/common/cointransaction/cointransactionimpl"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/bookkeeping"
	"github.com/hyperledger/fabric/core/ledger/kvledger/coin"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	mdl "github.com/hyperledger/fabric/core/peer/mdldata"
	"github.com/hyperledger/fabric/core/peer/mdldata"
	"github.com/hyperledger/fabric/mdl"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
const (
	EMPTY_ADMIN_ADDRESS = ""
	ADMIN_ADDRESS       = "0xadmin"
	FEE_ADMIN_ADDRESS   = "0xfeeadmin"
	USER1_ADDRESS       = "0xuser1"
	USER2_ADDRESS       = "0xuser2"
	USER3_ADDRESS       = "0xuser3"
)

func TestInitNotAthorizedProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	init := coinTxProcessor.rwsetParser.Clone("init")
	appendKvWrite(nsRwSet, init, []byte(fmt.Sprintf("init,,5000,%s,5,10", FEE_ADMIN_ADDRESS)))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_NOT_ATHORIZED)
}

func TestInitValidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	init := coinTxProcessor.rwsetParser.Clone("init")
	appendKvWrite(nsRwSet, init, []byte(fmt.Sprintf("init,,5000,%s,5,10", FEE_ADMIN_ADDRESS)))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, EMPTY_ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_VALID)
}

func TestMintNotAthorizedProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	mint := coinTxProcessor.rwsetParser.Clone("mint")
	appendKvWrite(nsRwSet, mint, []byte("mint,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, USER1_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_NOT_ATHORIZED)
}

func TestMintValidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	mint := coinTxProcessor.rwsetParser.Clone("mint")
	appendKvWrite(nsRwSet, mint, []byte("mint,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_VALID)
}

func TestBurnNotAthorizedProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	burn := coinTxProcessor.rwsetParser.Clone("burn")
	appendKvWrite(nsRwSet, burn, []byte("burn,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, USER1_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_NOT_ATHORIZED)
}

func TestBurnInvalidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	burn := coinTxProcessor.rwsetParser.Clone("burn")
	appendKvWrite(nsRwSet, burn, []byte("burn,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_INVALID_BALANCE)
}

func TestBurnValidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	mint := coinTxProcessor.rwsetParser.Clone("mint")
	appendKvWrite(nsRwSet, mint, []byte("mint,50"))
	appendNsRWSet(txAction, nsRwSet)

	burn := coinTxProcessor.rwsetParser.Clone("burn")
	appendKvWrite(nsRwSet, burn, []byte("burn,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_VALID)
}

func TestFeeConfigNotAthorizedProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	burn := coinTxProcessor.rwsetParser.Clone("feeconfig")
	appendKvWrite(nsRwSet, burn, []byte("feeconfig,coin,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, USER1_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_NOT_ATHORIZED)
}

func TestFeeConfigValidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	mint := coinTxProcessor.rwsetParser.Clone("feeconfig")
	appendKvWrite(nsRwSet, mint, []byte("feeconfig,coin,50"))
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_VALID)
}

func TestTransferInvalidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	mint := coinTxProcessor.rwsetParser.Clone("mint")
	appendKvWrite(nsRwSet, mint, []byte("mint,50"))
	appendNsRWSet(txAction, nsRwSet)

	fee := []byte("5")
	transfer := coinTxProcessor.rwsetParser.Clone("transfer")
	appendKvWriteWithFeeMetadata(nsRwSet, transfer, []byte(fmt.Sprintf("transfer,%s,50", USER2_ADDRESS)), fee)
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_INVALID_BALANCE)
}

func TestTransferValidProcess(t *testing.T) {
	coinTxProcessor := getCoinTxProcessor()
	txAction := makeUtxAction()
	nsRwSet := makeNsRwSet()

	mint := coinTxProcessor.rwsetParser.Clone("mint")
	appendKvWrite(nsRwSet, mint, []byte("mint,55"))
	appendNsRWSet(txAction, nsRwSet)

	fee := []byte("5")
	transfer := coinTxProcessor.rwsetParser.Clone("transfer")
	appendKvWriteWithFeeMetadata(nsRwSet, transfer, []byte(fmt.Sprintf("%s,50", USER1_ADDRESS)), fee)
	appendNsRWSet(txAction, nsRwSet)

	utx := createUtx(txAction, ADMIN_ADDRESS, ADMIN_ADDRESS)

	resultCode, _ := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_VALID)
}

func createUtx(txAction *mdldata.UnmarshaledTransactionAction, user, admin string) *mdldata.UnmarshaledTransaction {
	utx := makeUtx(user, admin)
	utx.TxAction = append(utx.TxAction, txAction)
	return utx
}

func appendNsRWSet(txAction *mdldata.UnmarshaledTransactionAction, nsRwSet *rwsetutil.NsRwSet) {
	txAction.TxRwSet.(*rwsetutil.TxRwSet).NsRwSets = append(txAction.TxRwSet.(*rwsetutil.TxRwSet).NsRwSets, nsRwSet)
}

func appendKvWrite(nsRwSet *rwsetutil.NsRwSet, tx cointransaction.CoinTransaction, args []byte) {
	kvWrite := &kvrwset.KVWrite{
		Key:   tx.Key(),
		Value: args,
	}
	nsRwSet.KvRwSet.Writes = append(nsRwSet.KvRwSet.Writes, kvWrite)
}

func appendKvWriteWithFeeMetadata(nsRwSet *rwsetutil.NsRwSet, tx cointransaction.CoinTransaction, args, fee []byte) {
	kvWrite := &kvrwset.KVWrite{
		Key:   tx.Key(),
		Value: args,
	}
	nsRwSet.KvRwSet.Writes = append(nsRwSet.KvRwSet.Writes, kvWrite)
	kvMetadataWrite := &kvrwset.KVMetadataWrite{
		Key:     mdl.CoinChaincodeName,
		Entries: make([]*kvrwset.KVMetadataEntry, 0, 0),
	}
	kvMetadataEntry := &kvrwset.KVMetadataEntry{
		Name:  "fee",
		Value: fee,
	}
	kvMetadataWrite.Entries = append(kvMetadataWrite.Entries, kvMetadataEntry)
	nsRwSet.KvRwSet.MetadataWrites = append(nsRwSet.KvRwSet.MetadataWrites, kvMetadataWrite)
}

func makeNsRwSet() *rwsetutil.NsRwSet {
	return &rwsetutil.NsRwSet{
		NameSpace: mdl.CoinChaincodeName,
		KvRwSet: &kvrwset.KVRWSet{
			Writes:         make([]*kvrwset.KVWrite, 0, 0),
			MetadataWrites: make([]*kvrwset.KVMetadataWrite, 0, 0),
		},
	}
}
func makeUtxAction() *mdldata.UnmarshaledTransactionAction {
	return &mdldata.UnmarshaledTransactionAction{
		TxRwSet: &rwsetutil.TxRwSet{
			NsRwSets: make([]*rwsetutil.NsRwSet, 0, 0),
		},
	}
}

func makeUtx(from, admin string) *mdldata.UnmarshaledTransaction {
	return &mdldata.UnmarshaledTransaction{
		TxAction: make([]*mdldata.UnmarshaledTransactionAction, 0, 0),
		CoinTx: &cointransaction.CoinTransactionInfo{
			Admin:    admin,
			FeeAdmin: FEE_ADMIN_ADDRESS,
			From:     from,
		},
	}
}

func getCoinTxProcessor() *CoinTxProcessorImpl2 {
	levelDBPath := filepath.Join(os.TempDir(), "test")
	bookkeepingPath := filepath.Join(os.TempDir(), "bookkeeping")

	defer os.RemoveAll(levelDBPath)
	defer os.RemoveAll(bookkeepingPath)
	bookkeepingProvider, _ := bookkeeping.NewProvider(
@@ -48,55 +299,5 @@ func TestProcess(t *testing.T) {
	coinBatch := coin.NewCoinBatchProcessor(coinLedger.GetQueryExcutor())
	coinTxSim := NewCoinTxExcutor(coinBatch)
	rwsetParser := rwsetparser.NewRWSetParser(cointransactionfactoryimpl.CoinTransactionFactory())
	coinTxProcessor := NewCoinTxProcessor2(rwsetParser, coinTxSim, NewCoinTxValidator(coinBatch))
	utx := createRwset()
	resultCode, result := coinTxProcessor.Process(utx, uint64(1))
	assert.Equal(t, resultCode, cointransaction.CoinTxResultCode_NOT_ATHORIZED)
	t.Log(result)
	t.Log(resultCode)

}

func createRwset() *mdl.UnmarshaledTransaction {

	txAction := mdl.UnmarshaledTransactionAction{}
	txRwSet := &rwsetutil.TxRwSet{
		NsRwSets: make([]*rwsetutil.NsRwSet, 0, 0),
	}
	txAction.TxRwSet = txRwSet

	nsRwSet := &rwsetutil.NsRwSet{NameSpace: "mdl"}
	nsRwSet.KvRwSet = &kvrwset.KVRWSet{}
	nsRwSet.KvRwSet.Writes = make([]*kvrwset.KVWrite, 0, 0)
	kvwrite := &kvrwset.KVWrite{}
	nsRwSet.KvRwSet.Writes = append(nsRwSet.KvRwSet.Writes, kvwrite)
	nsRwSet.KvRwSet.MetadataWrites = make([]*kvrwset.KVMetadataWrite, 0, 0)
	kvmetawrite := &kvrwset.KVMetadataWrite{}
	nsRwSet.KvRwSet.MetadataWrites = append(nsRwSet.KvRwSet.MetadataWrites, kvmetawrite)
	txAction.TxRwSet.(*rwsetutil.TxRwSet).NsRwSets = append(txAction.TxRwSet.(*rwsetutil.TxRwSet).NsRwSets, nsRwSet)

	mint := cointransactionimpl.MintTransaction{}
	kvwrite.Key = mint.Key()
	kvwrite.Value = []byte("mint,50")

	// kvmetawrite.Key = "mdl"
	// kvmetawrite.Entries = make([]*kvrwset.KVMetadataEntry, 0, 0)
	// kvmetaEndtry := &kvrwset.KVMetadataEntry{}

	// fee := cointransaction.FeeTransaction{from: "", to: "", amount: 123}
	// kvmetaEndtry.Name = fee.Key()
	// kvmetaEndtry.Value = fee.Serialize()
	// kvmetawrite.Entries = append(kvmetawrite.Entries, kvmetaEndtry)

	utx := &mdl.UnmarshaledTransaction{
		TxAction: make([]*mdl.UnmarshaledTransactionAction, 0, 0),
		CoinTx: &cointransaction.CoinTransactionInfo{
			Admin:    "admin",
			FeeAdmin: "feeadmin",
			From:     "user1",
		},
	}
	utx.TxAction = append(utx.TxAction, &txAction)

	return utx
	return NewCoinTxProcessor2(rwsetParser, coinTxSim, NewCoinTxValidator(coinBatch))
}

