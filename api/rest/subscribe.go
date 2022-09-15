package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/model"
	"github.com/kerrrusha/btc-api/api/internal/model/dataStorage/fileStorage"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

func CreateEmptyEmailsJSON(filepath string) {
	emails := model.Emails{Emails: []string{}}
	emailsJSON, err := json.Marshal(emails)
	utils.CheckForError(err)

	writer := fileStorage.CreateFileWriter(filepath)
	writer.Write(string(emailsJSON), false)
}

func ReadEmails(filepath string) model.Emails {
	var emails model.Emails

	if utils.FileNotExist(filepath) || utils.FileIsEmpty(filepath) {
		CreateEmptyEmailsJSON(filepath)
	}

	reader := fileStorage.CreateFileReader(filepath)
	fileBytes := reader.Read()
	err := json.Unmarshal(fileBytes, &emails)
	utils.CheckForError(err)

	return emails
}

func WriteNewEmailToFile(filepath string, email string) {
	var emails model.Emails
	storage := fileStorage.CreateFileStorage(filepath)

	fileBytes := storage.Read()

	err := json.Unmarshal(fileBytes, &emails)
	utils.CheckForError(err)

	emails.Emails = append(emails.Emails, email)

	emailsJSON, err := json.Marshal(emails)
	utils.CheckForError(err)

	storage.Write(string(emailsJSON), true)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log.Println("subscribe endpoint")

	decoder := json.NewDecoder(r.Body)
	cfg := config.GetConfig()

	var newEmail model.Email
	err := decoder.Decode(&newEmail)
	utils.CheckForError(err)

	if utils.StringArraySearch(ReadEmails(cfg.GetEmailsFilepath()).Emails, newEmail.Email) != -1 {
		utils.SendResponse(
			w,
			model.ErrorResponse{Error: "Email was not subscribed: it already exists"},
			http.StatusConflict,
		)
		return
	}
	if !newEmail.IsValid() {
		utils.SendResponse(
			w,
			model.ErrorResponse{Error: "Email is not correct. Please, enter valid email"},
			http.StatusConflict,
		)
		return
	}

	WriteNewEmailToFile(cfg.GetEmailsFilepath(), newEmail.Email)

	utils.SendResponse(
		w,
		model.SuccessResponse{Success: "Email was subscribed successfully"},
		http.StatusOK,
	)
}
