package main

import (
	"net/http"
	"log"
	"encoding/json"
	"os/exec"
	"strings"
)

// json -> struct
type fileNames struct {
	Names []string `json:"names"`
}

// curl -X POST -d "{\"names\": [\"doge1.jpg\", \"doge2.jpg\"]}" http://localhost:8080/gifize
//TODO: make only posts
func gifize(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var fn fileNames   
	err := decoder.Decode(&fn)
	if err != nil {
		log.Println("error decoding JSON")
	} else {
		log.Println("received JSON")
	}

	gifCommand := constructGifCommand(fn.Names)	
	log.Println(gifCommand)

	makeGif(gifCommand)


}

func constructGifCommand(names []string) string {
	gifCommand := "gifsicle --delay=75 --loop "
	filePrefix := "./daily_pictures/"
	gifStorageFile := "./pong.gif"

	for i := range names {
		fileAddress := strings.Split(names[i], "/")
		fileName := fileAddress[len(fileAddress)-1]
		gifCommand += filePrefix + fileName + " "
	}

	gifCommand += "> " + gifStorageFile

	return gifCommand
}

func makeGif(gifCommand string) {
	cmd := exec.Command("sh", "-c", gifCommand)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("making gif")
	err = cmd.Wait()
	if err != nil {
		log.Printf("gif creation has error: %v", err)
	} else {
		log.Println("gif created")
	}
}

func main() {
	// serve most current photo
	http.HandleFunc("/latest.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./daily_pictures/doge1.jpg")
	})

	// make gifs given pics
	http.HandleFunc("/gifize", gifize)

	// serve all photos
	http.Handle("/pong-pics/", http.StripPrefix("/pong-pics/", http.FileServer(http.Dir("./daily_pictures"))))

	// home site
	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

