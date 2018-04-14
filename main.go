package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/quran", getFile).Methods("GET","OPTIONS")
	log.Fatal(http.ListenAndServe(":8089",router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Oid struct {
	Oid string `json:"$oid"`
}

type Ayat struct {
	ID Oid `json:"_id"`
	Index int8 `json:"index"`
	Sura int8 `json:"sura"'`
	Aya int8 `json:"aya"`
	Text string `json:"text"`
}

func getAyat() [] Ayat {
	ayat := make([]Ayat, 12)
	raw, err := ioutil.ReadFile("file/quran.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &ayat)
	return ayat
}

func getFile(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	json.NewEncoder(res).Encode(getAyat())
}