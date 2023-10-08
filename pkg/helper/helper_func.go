package helper

import (
	"Nik/pkg/models"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	lifetime   = 3 * time.Hour
	signingKey = "grkjk#4#%35FSFJLja#4353KSFjH"
	KeyUserId  = "userID"
)

/*
	Takes the response text as input, adds it to the field

of the structure it, and sends it to the User.
*/
func ResponseAnswer(w http.ResponseWriter, report string) (err error) {
	answer := models.Answer{
		Date:           time.Now(),
		ResponseAnswer: report,
	}
	myAnswer, err := json.MarshalIndent(answer, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(myAnswer)
	if err != nil {
		return err
	}
	return
}

func SendToken(w http.ResponseWriter, sendToken *models.SendToken) error {
	answer, err := json.MarshalIndent(sendToken, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(answer)
	if err != nil {
		return err
	}

	return nil
}

func BadRequest(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info("the user entered incorrect data", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Error(http.StatusText(http.StatusInternalServerError), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func Forbidden(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info(http.StatusText(http.StatusForbidden), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func NotFoundErr(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info(http.StatusText(http.StatusNotFound), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func ResetContentServerError(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info(http.StatusText(http.StatusResetContent), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusResetContent), http.StatusResetContent)
}

func CreateToken(login string) (token string, err error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: login,
	})
	token, err = tok.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserIdFromContext(r *http.Request) (id int, err error) {
	id, ok := r.Context().Value(KeyUserId).(int)
	if !ok {
		err = errors.New("User is not found!")
		return 0, err
	}
	return id, nil
}

func ConvertToIntTheParams(strPage, strCount string) (page, count int, err error) {
	page, err = strconv.Atoi(strPage)
	if err != nil {
		return 0, 0, err
	}

	count, err = strconv.Atoi(strCount)
	if err != nil {
		return 0, 0, err
	}

	return page, count, nil
}

func SendContents(w http.ResponseWriter, contents []*models.SendContents) error {
	if contents == nil {
		err := ResponseAnswer(w, "You don't have any content yet!")
		if err != nil {
			return err
		}
		return err
	}

	bytesContents, err := json.MarshalIndent(contents, "", "    ")
	if err != nil {
		return err
	}

	_, err = w.Write(bytesContents)
	if err != nil {
		return err
	}

	return nil
}

func SendContent(w http.ResponseWriter, contents *models.SendConcreteContent) error {
	if contents == nil {

		err := ResponseAnswer(w, "You don't have any content yet!")
		if err != nil {
			return err
		}
		return err
	}

	bytesContents, err := json.MarshalIndent(contents, "", "    ")
	if err != nil {
		return err
	}

	_, err = w.Write(bytesContents)
	if err != nil {
		return err
	}

	return nil
}

func SendMessages(w http.ResponseWriter, messages any, logger *zap.Logger) error {
	switch messages.(type) {
	case []*models.SendNotification:
		if messages == nil {
			w.WriteHeader(http.StatusNotFound)
			err := ResponseAnswer(w, "You have no message for this ID!")
			if err != nil {
				InternalServerError(w, err, logger)
				return err
			}

			err = errors.New("You have no message for this ID!")
			return err
		}

		byteMessages, err := json.MarshalIndent(messages, "", "    ")
		if err != nil {
			return err
		}

		_, err = w.Write(byteMessages)
		if err != nil {
			return err
		}

		return nil
	case *models.SendNotification:
		if messages == nil {
			w.WriteHeader(http.StatusNotFound)
			err := ResponseAnswer(w, "You have no message for this ID!")
			if err != nil {
				InternalServerError(w, err, logger)
				return err
			}

			err = errors.New("You have no message for this ID!")
			return err
		}

		byteMessages, err := json.MarshalIndent(messages, "", "    ")
		if err != nil {
			return err
		}

		_, err = w.Write(byteMessages)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func SendLinks(w http.ResponseWriter, links *models.SendLinks) error {
	if links == nil {
		err := ResponseAnswer(w, "Content don't have any Links yet!")
		if err != nil {
			return err
		}
		return err
	}

	bytesContents, err := json.MarshalIndent(links, "", "    ")
	if err != nil {
		return err
	}

	_, err = w.Write(bytesContents)
	if err != nil {
		return err
	}

	return nil
}
