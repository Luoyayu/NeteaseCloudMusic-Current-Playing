package NeteaseMusicPlaying

type HistoryStruct struct {
	Track struct {
		Source struct {
			SourceId   int64  `json:"SourceId"`
			SourceName string `json:"SourceName"`
		} `json:"Source"`
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		Artists []struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"ar"`
		Album struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
			Pic  int64  `json:"pic"`
		} `json:"al"`
		Cd        string  `json:"cd"`
		No        int     `json:"no"`
		Pop       float64 `json:"pop"`
		Fee       int     `json:"fee"`
		Privilege struct {
			Id    int64 `json:"id"`
			Fee   int   `json:"fee"`
			Payed int   `json:"payed"`
			Pl    int   `json:"pl"`
			Dl    int   `json:"dl"`
			Maxbr int   `json:"maxbr"`
		} `json:"privilege"`
		Dt int64 `json:"dt"`
	} `json:"track"`
}
