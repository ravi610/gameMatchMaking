package main

import (
	"time"
)

func main() {
	startMatchMaking()
	computeMatchesForPlayers()
	computeMatchesForTeams()
	buildTeam()

	player1 := newPlayer(Player{
		playerID:    "1",
		gameID:      "chess",
		gameType:    "singlePlayer",
		rating:      1200,
		requestTime: 1605607168,
	})

	player2 := newPlayer(Player{
		playerID:    "2",
		gameID:      "chess",
		gameType:    "singlePlayer",
		rating:      1250,
		requestTime: 1605607168,
	})

	player3 := newPlayer(Player{
		playerID:    "3",
		gameID:      "chess",
		gameType:    "singlePlayer",
		rating:      1250,
		requestTime: 1605607168,
	})

	player4 := newPlayer(Player{
		playerID:    "4",
		gameID:      "chess",
		gameType:    "singlePlayer",
		rating:      1250,
		requestTime: 1605607168,
	})

	player1.addMatchMakingRequest()
	player3.addMatchMakingRequest()
	player2.addMatchMakingRequest()
	player4.addMatchMakingRequest()

	team1 := newTeam(Team{
		teamID:      "xyz123",
		playerIDs:   []string{"1", "2", "3"},
		gameID:      "chess",
		gameType:    "multiPlayer",
		rating:      1000,
		requestTime: 1605607168,
	})

	team2 := newTeam(Team{
		teamID:      "abc123",
		playerIDs:   []string{"4", "5", "6"},
		gameID:      "chess",
		gameType:    "multiPlayer",
		rating:      1100,
		requestTime: 1605607168,
	})

	team1.addMatchMakingRequest()
	team2.addMatchMakingRequest()

	player11 := newPlayer(Player{
		playerID:    "1",
		gameID:      "chess",
		gameType:    "multiPlayer",
		rating:      1200,
		requestTime: 1605607168,
	})

	player12 := newPlayer(Player{
		playerID:    "2",
		gameID:      "chess",
		gameType:    "multiPlayer",
		rating:      1250,
		requestTime: 1605607168,
	})

	player13 := newPlayer(Player{
		playerID:    "3",
		gameID:      "chess",
		gameType:    "multiPlayer",
		rating:      1250,
		requestTime: 1605607168,
	})

	team11 := newTeam(Team{
		teamID:      "pqr123",
		playerIDs:   []string{"1", "2", "3"},
		gameID:      "chess",
		gameType:    "multiPlayer",
		rating:      1300,
		requestTime: 1605607168,
	})

	team11.addMatchMakingRequest()
	player11.addTeamBuildingRequest()
	player12.addTeamBuildingRequest()
	player13.addTeamBuildingRequest()

	time.Sleep(10 * time.Second)
}
