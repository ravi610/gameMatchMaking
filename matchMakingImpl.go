package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var playerMap map[string]Player
var teamMap map[string]Team
var playerLookUp map[string]Player
var exists = struct{}{}
var mu sync.Mutex

func startMatchMaking() {
	playerMap = make(map[string]Player)
	teamMap = make(map[string]Team)
	playerLookUp = make(map[string]Player)
}

func newPlayer(p Player) *Player {
	return &p
}

func newTeam(t Team) *Team {
	return &t
}

func (player *Player) addMatchMakingRequest() {
	mu.Lock()
	if _, exist := playerMap[player.playerID]; exist {
		return
	}

	playerMap[player.playerID] = *player
	mu.Unlock()
}

func (team *Team) addMatchMakingRequest() {
	mu.Lock()
	if _, exist := teamMap[team.teamID]; exist {
		return
	}

	teamMap[team.teamID] = *team
	mu.Unlock()
}

func (player *Player) addTeamBuildingRequest() {
	mu.Lock()
	if _, exist := playerLookUp[player.playerID]; exist {
		return
	}

	playerLookUp[player.playerID] = *player
	mu.Unlock()
}

//assuming team size will be a constant = 3
func buildTeam() {
	go func() {
		for {
			mu.Lock()

			for key, player := range playerLookUp {
				size := 1
				list := make([]string, 3)
				avgRating := player.rating
				list[0] = player.playerID

				for key1, player1 := range playerLookUp {
					if key == key1 || player.gameID != player1.gameID || player.gameType != player1.gameType {
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
						teamID:      String(6),
						playerIDs:   list,
						gameID:      player.gameID,
						gameType:    player.gameType,
						rating:      avgRating,
						requestTime: time.Now().Unix(),
					}
					teamMap[team.teamID] = team

					for _, value := range list {
						delete(playerLookUp, value)
					}
					break
				}
			}
			mu.Unlock()
		}
	}()
}

func computeMatchesForPlayers() {
	go func() {
		for {
			mu.Lock()

			for key, player := range playerMap {
				matched := false
				for key1, player1 := range playerMap {
					if key == key1 || player.gameID != player1.gameID || player.gameType != player1.gameType {
						continue
					}

					ratingDiff := player.rating - player1.rating
					//now is current epoch in seconds
					now := time.Now().Unix()
					timeDiff := float64(now - player.requestTime)

					if math.Abs(float64(ratingDiff)) <= (100.00 + 0.1*timeDiff) {
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
			mu.Unlock()
		}
	}()
}

func computeMatchesForTeams() {
	go func() {
		for {
			mu.Lock()
			for key, team := range teamMap {
				matched := false
				for key1, team1 := range teamMap {
					if key == key1 || team.gameID != team1.gameID || team.gameType != team1.gameType {
						continue
					}

					ratingDiff := team.rating - team1.rating
					//now is current epoch in seconds
					now := time.Now().Unix()
					timeDiff := float64(now - team.requestTime)

					if math.Abs(float64(ratingDiff)) <= (100.00 + 0.1*timeDiff) {
						fmt.Printf("team 1 : %v and team 2 : %v got matched \n", team.teamID, team1.teamID)
						delete(teamMap, team1.teamID)
						matched = true
						break
					}
				}

				if matched == true {
					delete(teamMap, team.teamID)
					break
				}
			}
			mu.Unlock()
		}
	}()
}

func (player *Player) removeMatchMakingRequest() {
	mu.Lock()
	if _, exist := playerMap[player.playerID]; !exist {
		return
	}
	delete(playerMap, player.playerID)
	mu.Unlock()
}

func (team *Team) removeMatchMakingRequest() {
	mu.Lock()
	if _, exist := teamMap[team.teamID]; !exist {
		return
	}
	delete(teamMap, team.teamID)
	mu.Unlock()
}
