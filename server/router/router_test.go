package router

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func init() {
	r = CreateRouter()
}

func TestLoginFail(t *testing.T) {
	w := httptest.NewRecorder()

	data := url.Values{}
	data.Set("Email", "fake@email.com")
	data.Set("Password", "password123")

	body := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", "/login", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
