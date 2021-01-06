package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type UploadController struct {
}

func NewUploadController() UploadController {
	return UploadController{}
}

type Response struct {
	Url string `json:"url"`
}

func (u *UploadController) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("avatar")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Parsing File"))
		return
	}
	defer file.Close()

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	tempFile, err := ioutil.TempFile("upload", "avatar-*.png")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Uploading File"))
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Uploading File"))
		return
	}
	tempFile.Write(fileBytes)

	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(&Response{Url: tempFile.Name()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Marshalling Response"))
		return
	}
	w.Write(response)
}
