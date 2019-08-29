package SafexRPC

import (
	"net/http"

	"github.com/safex/gosafex/pkg/chain"
	log "github.com/sirupsen/logrus"
)

type WalletDummy struct {
}

type WalletRPC struct {
	logger  *log.Logger
	wallet  *chain.Wallet
	mainnet bool // false for testnet
}

func (w *WalletRPC) OpenCheck(rw *http.ResponseWriter) bool {
	if w.wallet == nil || !w.wallet.IsOpen() {
		FormJSONResponse(nil, WalletIsNotOpened, rw)
		return false
	}
	return true
}

// Getting status of current wallet. If its open, syncing etc.
func (w *WalletRPC) GetStatus(rw http.ResponseWriter, r *http.Request) {
	var data JSONElement
	w.logger.Infof("[RPC] Getting wallet status")
	data = make(JSONElement)
	data["msg"] = "Hello Load"

	FormJSONResponse(data, EverythingOK, &rw)

}

// Getting status of current wallet. If its open, syncing etc.
func (w *WalletRPC) Close(rw http.ResponseWriter, r *http.Request) {
	w.logger.Infof("[RPC] Closing wallet")
	w.wallet.Close()
	w.wallet = nil
	data := make(JSONElement)
	data["msg"] = "Its closed!"

	FormJSONResponse(data, EverythingOK, &rw)

}

func (w *WalletRPC) SetLogger(prevLog *log.Logger) {
	w.logger = prevLog
}
