package main

type Result struct {
	GameId int
	UserId int
	Score  int
}

func main() {
	var resultA Result
	resultA.GameId = 1
	resultA.UserId = 1
	resultA.Score = 100
}
