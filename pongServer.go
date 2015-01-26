package main

func main() {
    getGif("/pong.gif", "/root/pong_bot/pong_output_folder/output000001.gif")
    http.ListenAndServe(":80", nil)


import (
	"net/http"
	"log"
	"encoding/json"
)

// needed to loosen up access protocol?
func changeHeaderThenServe(h http.Handler) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) { 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r) 
	}
}

type fileNames struct {
    Names []string
}

// curl -X POST -d "{\"names\": [\"that\", \"and\"]}" http://localhost:8080/gifize
func gifize(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var fn fileNames   
    err := decoder.Decode(&fn)
    if err != nil {
    	log.Println("error decoding JSON")
    } else {
    	log.Printf("received file names: %s", fn.Names)
	}

	// iterate over file names to call gif script
	// copy selected files to another folder and gifify them
	for idx := range fn.Names {
		log.Println(fn.Names[idx])
	}
}

func main() {

	// TODO: make these endpoints
	// angularJS home site -- weird with how file is served :(
	// POST site to make custom gif from selection and let you download it

	// serve most current photo
	// http.HandleFunc("/latest.gif", "/root/angular_server/pong_snapshots/latest.gif")
	http.HandleFunc("/gifize", gifize)

	// serve all photos
	http.Handle("/pong-pics/", http.StripPrefix("/pong-pics/", changeHeaderThenServe(http.FileServer(http.Dir("./daily_pictures")))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
