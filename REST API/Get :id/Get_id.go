func getUserInfoHandler(w http.ResponseWriter, r *http.Request) { //고정 89가 아니라 아이디를 나타내줘야하므로 mux.Vars사용한다
	vars := mux.Vars(r) //알아서 테스트 패싱해줌
	fmt.Fprint(w, "User Id:", vars["id"])
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	// mux.HandleFunc("/users/89", getUserInfo89Handler) 89가 아니라 아이디를 나타내는 {id:[0-9]+}문법으로 (고릴라)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux
}

func TestGetUserInfo(t *testing.T) { //Get 설명
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89") // 1. 특정ID에 해당하는 유저(89번유저)의 정보를 가져오고싶다를 http.Get으로 보냄
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:89") // 2. 만약 그 유저가 없으면 No User Id가 리턴
	//사실 정해진 아이디가 아닌 유저아이디가 뒤에 붙었을때 맞게 와야함 URL.Path를 가지고

	resp, err = http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:56")
}
