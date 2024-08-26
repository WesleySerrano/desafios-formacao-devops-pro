package main

type Match struct {
	Id      uint
	Date    string
	Home    Team
	Visitor Team
	Scores  []Score
}
