package main

import (
	"encoding/json"
	"fmt"
)

type TennisDoubles struct {
	Players  []string
	Nickname string `json:",omitempty"`
	W        int    `json:"Won"`
	L        int    `json:"Lost"`
	ties     int
}

func main() {
	var teams = []TennisDoubles{
		{
			Players:  []string{"Chris Evert", "Pam Shriver"},
			Nickname: "Hot Shots",
			W:        7,
			L:        3,
			ties:     2},

		{
			[]string{"Venus Williams", "Maria Sharapova"},
			"",
			9,
			1,
			1}}

	jsonData, err := json.MarshalIndent(teams, "", "\t")

	if err != nil {
		fmt.Println("Marshaling failed!")
		return
	}
	fmt.Println(teams)
	fmt.Printf("%s", jsonData)
}
