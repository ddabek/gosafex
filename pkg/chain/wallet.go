package chain

import (
	"github.com/safex/gosafex/pkg/balance"
)

// TODO: figure out where to place the wallet struct.

// Wallet is a structure containing an Account, its Balance and tx Outputs.
type Wallet struct {
	balance balance.Balance
	account Account
	client  *Client
	// outputs OutputMap
}

// BlockFetchCnt is the the nubmer of blocks to fetch at once.
// TODO: Move this to some config, or recalculate based on response time
//const BlockFetchCnt = 100
//
//func matchOutput(txOut *safex.Txout, index uint64, der [32]byte, outputKey *[32]byte) bool {
//	derivatedPubKey := crypto.KeyDerivation_To_PublicKey(index, crypto.Key(der), w.Address.SpendKey.Public)
//	var outKeyTemp []byte
//	if txOut.Target.TxoutToKey != nil {
//		outKeyTemp, _ = hex.DecodeString(txOut.Target.TxoutToKey.Key)
//	} else {
//		outKeyTemp, _ = hex.DecodeString(txOut.Target.TxoutTokenToKey.Key)
//	}	// Return also outputkey
//	copy(outputKey[:], outKeyTemp[:32])
//	return *outputKey == [32]byte(derivatedPubKey)
//}

// // ProcessBlockRange processes all transactions in a range of blocks.
// func (w *Wallet) ProcessBlockRange(blocks safex.Blocks) bool {
// 	// @todo Here handle block metadata.

// 	// @todo This must be refactored due new discoveries regarding get_tx_hash
// 	// Get transaction hashes
// 	var txs []string
// 	for _, blck := range blocks.Block {
// 		txs = append(txs, blck.Txs...)
// 		txs = append(txs, blck.MinerTx)
// 	}

// 	// Get transaction data and process.
// 	loadedTxs, err := w.client.GetTransactions(txs)
// 	if err != nil {
// 		return false
// 	}

// 	for _, tx := range loadedTxs.Tx {
// 		w.ProcessTransaction(tx)
// 	}

// 	return true
// }

// func (w *Wallet) GetBalance() (b Balance, err error) {
// 	w.outputs = make(map[crypto.Key]*safex.Txout)
// 	// Connect to node.
// 	w.client = safexdrpc.InitClient("127.0.0.1", 38001)

// 	info, err := w.client.GetDaemonInfo()

// 	if err != nil {
// 		return b, errors.New("Cant get daemon info!")
// 	}

// 	bcHeight := info.Height

// 	var curr uint64
// 	curr = 0

// 	var blocks safex.Blocks
// 	var end uint64

// 	// @todo Here exists some error during overlaping block ranges. Deal with it later.
// 	for curr < (bcHeight - 1) {
// 		// Calculate end of interval for loading
// 		if curr+blockInterval > bcHeight {
// 			end = bcHeight - 1
// 		} else {
// 			end = curr + blockInterval
// 		}
// 		start := time.Now()
// 		blocks, err = w.client.GetBlocks(curr, end) // Load blocks from daemon
// 		fmt.Println(time.Since(start))

// 		// If there was error during loading of blocks return err.
// 		if err != nil {
// 			return b, err
// 		}

// 		// Process block
// 		w.ProcessBlockRange(blocks)

// 		fmt.Println("---------------------------------------------------------------------------------------------")
// 		curr = end
// 	}

// 	return w.balance, nil
// }