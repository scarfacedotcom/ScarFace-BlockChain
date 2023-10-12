package main

import (
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/scarface-blockchain/utils"
	"github.com/scarface-blockchain/wallet"
)

const tempDir = "/home/ubuntu/Desktop/GO/scarface-blockchain/wallet_server/templates/"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodGet:
// 		tmplPath := "/home/ubuntu/Desktop/GO/scarface-blockchain/wallet_server/templates/index.html"
// 		t, err := template.ParseFiles(tmplPath)

//			if err != nil {
//				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//				log.Printf("ERROR: Failed to parse template: %v", err)
//				return
//			}
//			err = t.Execute(w, nil)
//			if err != nil {
//				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//				log.Printf("ERROR: Failed to execute template: %v", err)
//			}
//		default:
//			log.Printf("ERROR: Invalid HTTP Method")
//		}
//	}

func (ws *WalletServer) Wallet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		io.WriteString(w, string(utils.JsonStatus("success")))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
