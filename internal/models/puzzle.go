package models

type Puzzle struct {
	Game struct {
		Id   string `json:"id"`
		Perf struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"perf"`
		Rated   bool `json:"rated"`
		Players []struct {
			UserId string `json:"userId"`
			Name   string `json:"name"`
			Color  string `json:"color"`
		} `json:"players"`
		Pgn   string `json:"pgn"`
		Clock string `json:"clock"`
	} `json:"game"`
	Puzzle struct {
		Id         string   `json:"id"`
		Rating     int      `json:"rating"`
		Plays      int      `json:"plays"`
		Solution   []string `json:"solution"`
		Themes     []string `json:"themes"`
		InitialPly int      `json:"initialPly"`
	} `json:"puzzle"`
}
