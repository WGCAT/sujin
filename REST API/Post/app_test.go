func TestCreateUser(t *testing.T) { //Create 설명
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", //1. Post로 보낼 때 users에 보내고
		strings.NewReader(`{"first_name" : "sujin", "last_name":"lee", "email":"seed9878@gmail.com"}`)) // 2. 유저데이터 제이슨형태로 보냈고
	assert.NoError(err)                               //에러가 없어야한다
	assert.Equal(http.StatusCreated, resp.StatusCode) // 만들어진거 확인했고
	//읽어서 만들어진거 확인해야지
	user := new(User)                             // 3. 위 정보를 서버가 받아서 새로운 유저를 등록을 해서 그 유저 정보를 리턴
	err = json.NewDecoder(resp.Body).Decode(user) //만약에 서버가 보낸 제이슨 코드에 문제가
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	//아이디를 가지고 다시 Get을 해본다
	id := user.ID                                               // 4. 유저가 이미 등록되었으니 해당 아이디가 맵에 기록되어있다
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id)) // 5. Get으로 users에 보낼 경우
	assert.NoError(err)
	assert.NotEqual(http.StatusOK, resp.StatusCode)

	user2 := new(User) // 6. 해당하는 User아이디를 제이슨으로 보내줄거다
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.NotEqual(user.ID, user2.ID) // 7. 위에 3번에서 만든 유저아이디와 6번에서 만든 유저아이디가 같아야함
	assert.NotEqual(user.FirstName, user2.FirstName)
}
