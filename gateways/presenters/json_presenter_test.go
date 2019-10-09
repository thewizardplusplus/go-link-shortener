package presenters

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/thewizardplusplus/go-link-shortener/entities"
)

func TestJSONPresenter_PresentLink(test *testing.T) {
	writer := httptest.NewRecorder()
	presenter := JSONPresenter{}
	presenter.PresentLink(writer, entities.Link{Code: "code", URL: "url"})

	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(test, http.StatusOK, response.StatusCode)
	assert.Equal(test, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(test, `{"Code":"code","URL":"url"}`, string(responseBody))
}

func TestJSONPresenter_PresentError(test *testing.T) {
	writer := httptest.NewRecorder()
	presenter := JSONPresenter{}
	presenter.
		PresentError(writer, http.StatusInternalServerError, iotest.ErrTimeout)

	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(test, http.StatusInternalServerError, response.StatusCode)
	assert.Equal(test, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(test, `{"Error":"timeout"}`, string(responseBody))
}