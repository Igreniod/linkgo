package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pesan struct {
	IdNumber        string `json:"id_number"`
	Nama            string `json:"nama"`
	Pesan           string `json:"pesan"`
	PesanAppPertama string `json:"pesan_app_pertama,omitempty"`
	PesanAppKedua   string `json:"pesan_app_kedua,omitempty"`
}

func olahPesan(w http.ResponseWriter, r *http.Request) {
	var newPesan Pesan
	err := json.NewDecoder(r.Body).Decode(&newPesan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//-----Cetak data di terminal-----//
	fmt.Println("Data diterima dari App 1:")
	cetakDataDiTerminal(newPesan)

	responAppSatu := newPesan.Pesan[1]
	switch responAppSatu {
	case '0':
		newPesan.PesanAppKedua = "Request Balance"
		cetakDataDiTerminal(newPesan)

	case '1':
		newPesan.PesanAppKedua = "Request Information"
		cetakDataDiTerminal(newPesan)

	case '2':
		newPesan.PesanAppKedua = "Transfer Credit"
		cetakDataDiTerminal(newPesan)

	case '3':
		newPesan.PesanAppKedua = "Deposit"
		cetakDataDiTerminal(newPesan)

	default:
		fmt.Println("Kode tidak tersedia")
		cetakDataDiTerminal(newPesan)
	}

	//-----Response ke port 8081-----//
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPesan)
}
