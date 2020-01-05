package NeteaseMusicPlaying


type Song struct {
	Tid          int64  `json:"tid"`
	Data         string `json:"data"`
	Href         string `json:"href"`
	Text         string `json:"text"`
	NickName     string `json:"nickName"`
	StartLogTime int64  `json:"startlogtime"`
	Qid          string `json:"qid"`
	Time         int64  `json:"time"`

	Track struct {
		Album struct {
			Id     int64    `json:"id"`
			Name   string   `json:"name"`
			PicId  int64    `json:"picId"`
			PicUrl string   `json:"picUrl"`
			Alias  []string `json:"alias"`
		} `json:"album"`
		Artists []struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"artists"`
		Alias           []string `json:"alias"`
		CommentThreadId string   `json:"commentThreadId"`
		CopyrightId     int64    `json:"copyrightId"`
		Duration        int64    `json:"duration"`
		Id              int64    `json:"id"`
		Mvid            int64    `json:"mvid"`
		Name            string   `json:"name"`
		Cd              string   `json:"cd"`
		Position        int      `json:"position"`
		Fee             int      `json:"fee"`
		Popularity      int      `json:"popularity"`
		Privilege       struct {
			Id        int64 `json:"id"`
			Payed     int   `json:"payed"`
			Fee       int   `json:"fee"`
			MaxPlayBr int   `json:"maxPlayBr"`
			MaxDownBr int   `json:"maxDownBr"`
			MaxSongBr int   `json:"maxSongBr"`
			MaxFreeBr int   `json:"maxFreeBr"`
		} `json:"privilege"`
	} `json:"track"`
}


