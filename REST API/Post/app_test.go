type User struct { //제이슨이 읽을 수 있는 유저 스트럭트를 만듦
	ID        int       `json:"ID"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) { //실제 유저를 생성하는 코드를 만들어야하는데 클라이언트가 유저정보를 제이슨으로 보냈음
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	userMap = make(map[int]*User) //맵에 유저를 언제 등록할거냐 바로위 크리에이트 할때
	lastID = 0
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET") //GET메소드일때 이 usersHandler가 불려라 정함
	mux.HandleFunc("/users", CreateUserHandler).Methods("POST")
	// mux.HandleFunc("/users/89", getUserInfo89Handler) 89가 아니라 아이디를 나타내는 {id:[0-9]+}문법으로 (고릴라)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux
}




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

