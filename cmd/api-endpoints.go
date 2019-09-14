package cmd

import (
	"fmt"
	"log"
	"net/http"

	"crypto/rand"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const maxUploadSize = 10 * 1024 * 1024 // 2 mb
const uploadPath = "C:/Temp"

// HomeMessage ...
func HomeMessage() string {
	return "Tech 4 Good !!!"
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HomeMessage())
}

// StartServer ...
func StartServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/message", homeLink)
	router.HandleFunc("/upload", uploadFileHandler())

	fmt.Println("Exposing API endpoints on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		fileType := r.PostFormValue("type")
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
		fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte("SUCCESS")); err != nil {
			renderError(w, "CANT_WRITE_FILE_BYTES", http.StatusInternalServerError)
			return
		}
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte(message)); err != nil {
		//err
	}
}

func randToken(len int) string {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		//err
	}
	return fmt.Sprintf("%x", b)
}
