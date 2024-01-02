package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bernanda0/oprec-netlab-be/database/sqlc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func NewDriveApiHandler(l *log.Logger, q *sqlc.Queries, u *AuthedUser) *DriveApiHandler {
	var c uint = 0
	return &DriveApiHandler{&Handler{l, q, &c, u}}
}

// Mewakili satu services
func (dah *DriveApiHandler) UploadFileH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodPost, dah.uploadFile}
	dah.h.handleRequest(hp, dah.h.u)
}

// Implementasi
func (dah *DriveApiHandler) uploadFile(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return err
	}

	// Retrieve form values
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	defer file.Close()

	parentID := r.FormValue("parent_id")

	fileID, err := uploadFile(handler.Filename, file, parentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	toJSON(w, fileID)
	return nil
}

func uploadFile(fileName string, file io.Reader, parentID string) (string, error) {
	// Read the service account JSON file
	credentials, err := ioutil.ReadFile("apiKey.json")
	if err != nil {
		return "", err
	}

	config, err := google.JWTConfigFromJSON(credentials, drive.DriveFileScope)
	if err != nil {
		return "", err
	}

	client := config.Client(oauth2.NoContext)

	service, err := drive.New(client)
	if err != nil {
		return "", err
	}

	fileMetadata := &drive.File{
		Name:    fileName,
		Parents: []string{parentID},
	}

	fileResource, err := service.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		return "", err
	}

	fileID := fileResource.Id
	link := fmt.Sprintf("https://drive.google.com/file/d/%s/view", fileID)

	return link, nil
}
