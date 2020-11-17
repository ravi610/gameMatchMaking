package main

import (
	"fmt"
	"math"
)

var playerMap map[string]Player
var teamMap map[string]Team
var playerLookUp map[string]Player
var exists = struct{}{}
var lock int

func startMatchMaking() {
	playerMap = make(map[string]Player)
	teamMap = make(map[string]Team)
	playerLookUp = make(map[string]Player)
	lock = 1
}

func newPlayer(p Player) *Player {
	return &p
}

func newTeam(t Team) *Team {
	return &t
}

func (player *Player) addMatchMakingRequest() {
	for lock != 1 {
	}
	lock = 0
	if _, exist := playerMap[player.playerID]; exist {
		return
	}

	playerMap[player.playerID] = *player
	lock = 1
}

func (team *Team) addMatchMakingRequest() {
	for lock != 1 {
	}
	lock = 0
	if _, exist := teamMap[team.teamID]; exist {
		return
	}

	teamMap[team.teamID] = *team
	lock = 1
}

func (player *Player) addTeamBuildingRequest() {
	for lock != 1 {
	}
	lock = 0
	if _, exist := playerLookUp[player.playerID]; exist {
		return
	}

	playerLookUp[player.playerID] = *player
	lock = 1
}

//assuming team size will be a constant = 3
func buildTeam() {
	go func() {
		for {
			for lock != 1 {
			}
			lock = 0

			for key, player := range playerLookUp {
				size := 1
				list := make([]string, 3)
				avgRating := player.rating
				list[0] = player.playerID

				for key1, player1 := range playerLookUp {
					if key == key1 {
						continue
					}

					ratingDiff := player.rating - player1.rating

					if math.Abs(float64(ratingDiff)) <= 100.00 {
						list[size] = player1.playerID
						avgRating = (avgRating*size + player1.rating) / (size + 1)
						size++
					}

					if size == 3 {
						break
					}
				}

				if size == 3 {
					team := Team{
						teamID:    String(6),
						playerIDs: list,
						gameID:    player.gameID,
						gameType:  player.gameType,
						rating:    avgRating,
					}
					teamMap[team.teamID] = team

					for _, value := range list {
						delete(playerLookUp, value)
					}
					break
				}
			}
			lock = 1
		}
	}()
}

func computeMatchesForPlayers() {
	go func() {
		for {
			for lock != 1 {
			}
			lock = 0

			for key, player := range playerMap {
				matched := false
				for key1, player1 := range playerMap {
					if key == key1 {
						continue
					}

					ratingDiff := player.rating - player1.rating

					if math.Abs(float64(ratingDiff)) <= 100.00 {
						fmt.Printf("player 1 : %v and player 2 : %v got matched \n", player.playerID, player1.playerID)
						delete(playerMap, player1.playerID)
						matched = true
						break
					}
				}

				if matched == true {
					delete(playerMap, player.playerID)
				}
			}
			lock = 1
		}
	}()
}

func computeMatchesForTeams() {
	go func() {
		for {
			for lock != 1 {
			}
			lock = 0
			for key, team := range teamMap {
				matched := false
				for key1, team1 := range teamMap {
					if key == key1 {
						continue
					}

					ratingDiff := team.rating - team1.rating

					if math.Abs(float64(ratingDiff)) <= 100.00 {
						fmt.Printf("team 1 : %v and team 2 : %v got matched \n", team.teamID, team1.teamID)
						delete(teamMap, team1.teamID)
						matched = true
						break
					}
				}

				if matched == true {
					delete(teamMap, team.teamID)
				}
			}
			lock = 1
		}
	}()
}

func (player *Player) removeMatchMakingRequest() {
	delete(playerMap, player.playerID)
}

func (team *Team) removeMatchMakingRequest() {
	delete(teamMap, team.teamID)
}
