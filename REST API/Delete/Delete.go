func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}
	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID:", id)
}

func NewHandler() http.Handler {
	userMap = make(map[int]*User) //맵에 유저를 언제 등록할거냐 바로위 크리에이트 할때
	lastID = 0
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET") //GET메소드일때 이 usersHandler가 불려라 정함
	mux.HandleFunc("/users", updateUserHandler).Methods("POST")
	// mux.HandleFunc("/users/89", getUserInfo89Handler) 89가 아니라 아이디를 나타내는 {id:[0-9]+}문법으로 (고릴라)
	mux.HandleFunc("/users", createUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

	return mux
}


func TestDeleteUser(t *testing.T) { //Delete 테스트
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	//Get이랑 Post는 http api에서 기본함수로 제공해주는데 delete는 없음
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name" : "sujin", "last_name":"lee", "email":"seed9878@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User ID:1")

}
