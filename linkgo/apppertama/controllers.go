package main

import (
	"bytes"
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
	fmt.Println("Data diterima dari HTTP Req:")
	cetakDataDiTerminal(newPesan)

	responAppSatu := newPesan.Pesan[0]
	switch responAppSatu {
	case '0':
		newPesan.PesanAppPertama = "Accepted"
		cetakDataDiTerminal(newPesan)

	case '1':
		newPesan.PesanAppPertama = "Pending"
		cetakDataDiTerminal(newPesan)

	case '2':
		newPesan.PesanAppPertama = "Rejected"
		cetakDataDiTerminal(newPesan)

	default:
		fmt.Println("Kode tidak tersedia")
		cetakDataDiTerminal(newPesan)
	}

	//-----Kirim/POST data ke app dua-----//
	pesanInJSON, err := json.MarshalIndent(newPesan, "", "  ")
	if err != nil {
		fmt.Println("Error mencetak JSON:", err)
		return
	}
	URLAppKedua := "http://localhost:8081/pesan"
	respon, err := http.Post(URLAppKedua, "application/json", bytes.NewBuffer(pesanInJSON))
	if err != nil {
		fmt.Println("Error saat POST ke app kedua", err)
		return
	}

	//-----decode reponse dari app kedua-----//
	json.NewDecoder(respon.Body).Decode(&newPesan)
	fmt.Println("Response dari app kedua :")
	cetakDataDiTerminal(newPesan)
	defer respon.Body.Close()

	//-----Response ke web channel-----//
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPesan)

}
