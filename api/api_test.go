package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestCreateModel(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := SetupRouter()

	model := Model{
		Message: "안녕하세요",
	}

	// JSON으로 직렬화
	jsonValue, _ := json.Marshal(model)

	// POST 요청 생성
	req, _ := http.NewRequest("POST", "/model", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// HTTP 응답 기록을 위한 기록기
	w := httptest.NewRecorder()

	// 요청을 라우터에 전달
	r.ServeHTTP(w, req)

	// 응답 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// 응답 내용 검증
	assert.Equal(t, "Model created", response["message"])
	assert.NotNil(t, response["data"])
	assert.Equal(t, model.Message, response["data"].(map[string]interface{})["message"])
}

func TestGetModelExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := SetupRouter()

	// 사전에 모델 생성하기
	model := Model{
		Message:   "안녕하세요",
		Timestamp: time.Now(),
	}
	c.Set("model", model, cache.DefaultExpiration)

	// GET 요청 생성
	req, _ := http.NewRequest("GET", "/model", nil)

	// HTTP 응답 기록을 위한 기록기
	w := httptest.NewRecorder()

	// 요청을 라우터에 전달
	r.ServeHTTP(w, req)

	// 응답 상태 코드 확인
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// 응답 내용 검증
	assert.NotNil(t, response["data"])
	assert.Equal(t, model.Message, response["data"].(map[string]interface{})["message"])
}

func TestGetModelNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := SetupRouter()
	// 테스트 시작 전에 캐시를 초기화
	c.Flush()

	// GET 요청 생성
	req, _ := http.NewRequest("GET", "/model", nil)

	// HTTP 응답 기록을 위한 기록기
	w := httptest.NewRecorder()

	// 요청을 라우터에 전달
	r.ServeHTTP(w, req)

	// 응답 상태 코드 확인
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// 응답 내용 검증
	assert.Equal(t, "Model not found", response["message"])
}
