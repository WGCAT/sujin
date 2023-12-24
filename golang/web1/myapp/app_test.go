package myapp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) { //앞에가 test로 시작하고 testing패키지 t인자를 포인트로 받는 함수
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/", nil)

	indexHandler(res, req)
	assert.Equal(http.StatusOK, res.Code) // 만약에 res.Code가 StatusOK가 아닐 경우에 자동으로 testFail이 됨 => 펑션을 통해 쉽게 가능하게

	// if res.Code != http.StatusBadRequest { // 만약에 res.Code가 StatusOK가 아닐 경우에 자동으로 testFail이 됨 => if문으로
	// 	t.Fatal("Failed!!", res.Code)
	// }
}
