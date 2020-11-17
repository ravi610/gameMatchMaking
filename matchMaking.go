package main

type matchMaking interface {
	addMatchMakingRequest()
	addTeamBuildingRequest()
	buildTeam()
	computeMatchesForPlayers()
	computeMatchesForTeams()
	removeMatchMakingRequest()
}
