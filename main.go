package main

import (
	"log"

	"github.com/secnex/ms-toolbox/webhook-finder/api"
	"github.com/secnex/ms-toolbox/webhook-finder/api/teams"
	"github.com/secnex/ms-toolbox/webhook-finder/config"
)

func main() {
	c := config.NewCustomConfig("config.json")
	file, _ := c.Load()

	graph := api.NewMsGraph(file.Client.ClientID, file.Client.ClientSecret, file.Client.TenantID)
	graph.NewAPIClient()
	graph.Authenticate(api.SCOPE_MS_GRAPH_DEFAULT)
	teamsAPI := teams.NewTeamsAPIBuilder(api.MS_GRAPH_API)
	teamsClient := teams.NewTeamsAPIClient(graph, teamsAPI)
	teams, err := teamsClient.ListTeams()
	if err != nil {
		log.Printf("ERROR: We could not list the teams in the tenant: %s", file.Client.TenantID)
	}
	if len(teams.Value) == 0 {
		log.Printf("INFO: There are no teams in the tenant: %s", file.Client.TenantID)
	}
}
