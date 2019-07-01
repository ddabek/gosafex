package filewallet

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/safex/gosafex/pkg/safex"
)

const filename = "test.db"
const blockFile = "blocks.test"
const outputFile = "outputs.test"
const walletName = "wallet1"
const masterPass = "masterpass"
const foldername = "test"

func prepareFolder() {

	fullpath := strings.Join([]string{foldername, filename}, "/")

	if _, err := os.Stat(fullpath); os.IsExist(err) {
		os.Remove(fullpath)
	}
	os.Mkdir(foldername, os.FileMode(int(0770)))
}

func CleanAfterTests(w *FileWallet, fullpath string) {

	w.Close()

	err := os.Remove(fullpath)
	if err != nil {
		fmt.Println(err)
	}
}

/*
func prepareWallet(w *FileWallet) {

	blockpath := strings.Join([]string{foldername, blockFile}, "/")
	outputpath := strings.Join([]string{foldername, outputFile}, "/")
	blockF, _ := os.Open(blockpath)
	outputF, _ := os.Open(outputpath)
	rblock := bufio.NewReader(blockF)
	routput := bufio.NewReader(outputF)

	arr := []string{"a", ""}
	for el, err := rblock.ReadString('\n'); err != io.EOF; el, err = rblock.ReadString('\n') {
		prevHash := arr[1]
		arr := strings.Split(el, ";")
		val, _ := strconv.Atoi(arr[0])
		header := &safex.BlockHeader{Depth: uint64(val), Hash: arr[1], PrevHash: prevHash}
		w.PutBlockHeader(header)
	}
	for el, err := routput.ReadString('\n'); err != io.EOF; el, err = routput.ReadString('\n') {
		arr := strings.Split(el, ";")
		val, _ := strconv.Atoi(arr[0])
		val1, _ := strconv.Atoi(arr[1])
		out := &safex.Txout{Amount: uint64(val)}
		w.AddOutput(out, uint64(val1), arr[2], arr[3], "normal", "")
	}
}
*/
func TestGenericDataRW(t *testing.T) {

	prepareFolder()
	fullpath := strings.Join([]string{foldername, filename}, "/")
	w, err := New(fullpath, walletName, masterPass, true)
	defer CleanAfterTests(w, fullpath)

	if err != nil {
		t.Fatalf("%s", err)
	}

	head1 := &safex.BlockHeader{Depth: 10, Hash: "aaaab", PrevHash: ""}
	head2 := &safex.BlockHeader{Depth: 11, Hash: "aaaac", PrevHash: "aaaab"}
	tx1 := &TransactionInfo{version: 1, unlockTime: 10, extra: []byte("asdasd"), blockHeight: head2.GetDepth(), blockTimestamp: 5, doubleSpendSeen: false, inPool: false, txHash: "tx01"}
	out1 := &safex.Txout{Amount: 20}

	err = w.PutBlockHeader(head1)
	if err != nil {
		t.Fatalf("%s", err)
	}
	err = w.PutBlockHeader(head2)
	if err != nil {
		t.Fatalf("%s", err)
	}

	err = w.PutTransactionInfo(tx1, head2.GetHash())
	if err != nil {
		t.Fatalf("%s", err)
	}
	outputInfo := &OutputInfo{"Cash", "aaaac", "tx01", "U", "normal"}
	outID, err := w.AddOutput(out1, 1, outputInfo, "")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if err = w.PutData("Test", []byte("asd")); err != nil {
		t.Fatalf("%s", err)
	}

	out, err := w.GetAllOutputs()

	if err != nil {
		t.Fatalf("%s", err)
	}

	found := false
	for _, el := range out {
		if el == outID {
			found = true
		}
	}
	if !found {
		fmt.Println(outID)
		for _, el := range out {
			fmt.Println(el)
		}
		t.Fatalf("Output not read")
	}

	if data, err := w.GetData("Test"); err != nil {
		t.Fatalf("%s", err)
	} else {
		if string(data) != "asd" {
			t.Fatalf("Failing reading generic data")
		}
	}

}

func TestBlockRW(t *testing.T) {
	prepareFolder()
	fullpath := strings.Join([]string{foldername, filename}, "/")
	w, err := New(fullpath, walletName, masterPass, true)
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer CleanAfterTests(w, fullpath)
	head1 := &safex.BlockHeader{Depth: 10, Hash: "aaaab", PrevHash: ""}
	head2 := &safex.BlockHeader{Depth: 11, Hash: "aaaac", PrevHash: "aaaab"}
	head3 := &safex.BlockHeader{Depth: 12, Hash: "aaaad", PrevHash: "aaaac"}
	head4 := &safex.BlockHeader{Depth: 13, Hash: "aaaae", PrevHash: "aaaad"}
	arr := []*safex.BlockHeader{head1, head2, head3, head4}
	if err = w.PutBlockHeader(head1); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutBlockHeader(head2); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutBlockHeader(head3); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutBlockHeader(head4); err != nil {
		t.Fatalf("%s", err)
	}
	if data, err := w.GetAllBlocks(); err != nil {
		t.Fatalf("%s", err)
	} else if len(data) != 4 {
		t.Fatalf("Read BlockHeader total length mismatch %d", len(data))
	} else {
		for i, el := range data {
			if head, err := w.GetBlockHeader(el); err != nil {
				t.Fatalf("%s", err)
			} else if head.GetHash() != arr[i].GetHash() {
				t.Fatalf("Read BlockHeader data mismatch\nhash: %s\nread:%s", arr[i].GetHash(), head.GetHash())
			}
		}
	}
	if err := w.RewindBlockHeader("aaaab"); err != nil {
		t.Fatalf("%s", err)
	}
	if data, err := w.GetAllBlocks(); err != nil {
		t.Fatalf("%s", err)
	} else if len(data) != 1 {
		t.Fatalf("Read BlockHeader total length mismatch %d", len(data))
	} else {
		head, err := w.GetBlockHeader(head1.GetHash())
		if err != nil {
			t.Fatalf("%s", err)
		}
		if head.GetHash() != head1.GetHash() {
			t.Fatalf("Error in block rewinding")
		}
		head, err = w.GetBlockHeader(head2.GetHash())
		if err == nil {
			t.Fatalf("Block not rewinded")
		}
	}
}

func TestTransactionRW(t *testing.T) {
	prepareFolder()
	fullpath := strings.Join([]string{foldername, filename}, "/")
	w, err := New(fullpath, walletName, masterPass, true)
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer CleanAfterTests(w, fullpath)
	head1 := &safex.BlockHeader{Depth: 10, Hash: "aaaab", PrevHash: ""}
	head2 := &safex.BlockHeader{Depth: 11, Hash: "aaaac", PrevHash: "aaaab"}
	tx1 := &TransactionInfo{version: 1, unlockTime: 10, extra: []byte("asdasd"), blockHeight: head2.GetDepth(), blockTimestamp: 5, doubleSpendSeen: false, inPool: false, txHash: "tx01"}
	tx2 := &TransactionInfo{version: 1, unlockTime: 10, extra: []byte("asdasd"), blockHeight: head2.GetDepth(), blockTimestamp: 5, doubleSpendSeen: false, inPool: false, txHash: "tx02"}
	tx3 := &TransactionInfo{version: 1, unlockTime: 10, extra: []byte("asdasd"), blockHeight: head2.GetDepth(), blockTimestamp: 5, doubleSpendSeen: false, inPool: false, txHash: "tx03"}
	tx4 := &TransactionInfo{version: 1, unlockTime: 10, extra: []byte("asdasd"), blockHeight: head2.GetDepth(), blockTimestamp: 5, doubleSpendSeen: false, inPool: false, txHash: "tx04"}

	if err = w.PutBlockHeader(head1); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutBlockHeader(head2); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutTransactionInfo(tx1, head2.GetHash()); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutTransactionInfo(tx2, head2.GetHash()); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutTransactionInfo(tx3, head2.GetHash()); err != nil {
		t.Fatalf("%s", err)
	}
	if err = w.PutTransactionInfo(tx4, head2.GetHash()); err != nil {
		t.Fatalf("%s", err)
	}
	transactionInfoArray, err := w.GetAllTransactionInfos()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(transactionInfoArray) != 4 {
		t.Fatalf("Read transaction  data mismatch %d", len(transactionInfoArray))
	}
	if err := w.RewindBlockHeader("aaaab"); err != nil {
		t.Fatalf("%s", err)
	}
	transactionInfoArray, err = w.GetAllTransactionInfos()
	if len(transactionInfoArray) != 0 {
		t.Fatalf("Error removing data")
	}
}

func TestOutputRW(t *testing.T) {
	prepareFolder()
	fullpath := strings.Join([]string{foldername, filename}, "/")
	w, err := New(fullpath, walletName, masterPass, true)
	if err != nil {
		t.Fatalf("%s", err)
	}
	head1 := &safex.BlockHeader{Depth: 10, Hash: "aaaab", PrevHash: ""}
	head2 := &safex.BlockHeader{Depth: 11, Hash: "aaaac", PrevHash: "aaaab"}
	tx1 := &TransactionInfo{version: 1, unlockTime: 10, extra: []byte("asdasd"), blockHeight: head2.GetDepth(), blockTimestamp: 5, doubleSpendSeen: false, inPool: false, txHash: "tx01"}
	out1 := &safex.Txout{Amount: 20}

	err = w.PutBlockHeader(head1)
	if err != nil {
		t.Fatalf("%s", err)
	}
	err = w.PutBlockHeader(head2)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if err != nil {
		t.Fatalf("%s", err)
	}

	err = w.PutTransactionInfo(tx1, head2.GetHash())
	if err != nil {
		t.Fatalf("%s", err)
	}

	outputInfo := &OutputInfo{"Cash", head2.GetHash(), "tx01", "U", "normal"}
	outID, err := w.AddOutput(out1, 1, outputInfo, "")
	if err != nil {
		t.Fatalf("%s", err)
	}
	outID, err = w.AddOutput(out1, 2, outputInfo, "")
	if err != nil {
		t.Fatalf("%s", err)
	}
	out, err := w.GetAllOutputs()

	if err != nil {
		t.Fatalf("%s", err)
	}
	found := false
	for _, el := range out {
		if el == outID {
			found = true
		}
	}
	if !found {
		t.Fatalf("Output not read")
	}

	w.Close()

	//Re-open just to read

	w, err = New(fullpath, walletName, masterPass, true)
	defer CleanAfterTests(w, fullpath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	out, err = w.GetAllOutputs()
	if err != nil {
		t.Fatalf("%s", err)
	}
	found = false
	for _, el := range out {
		if el == outID {
			found = true
		}
	}
	if !found {
		t.Fatalf("Output not read")
	}

	lock, err := w.GetOutputLock(outID)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if lock != "U" {
		t.Fatalf("Error in lock status consistency")
	}
	if err := w.LockOutput(outID); err != nil {
		t.Fatalf("%s", err)
	}

	lock, err = w.GetOutputLock(outID)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if lock != "L" {
		t.Fatalf("Error in changing lock status consistency")
	}

	if err := w.RewindBlockHeader("aaaab"); err != nil {
		t.Fatalf("%s", err)
	}
	out, err = w.GetAllOutputs()
	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(out) != 0 {
		t.Fatalf("Error removing data")
	}
}