package teams

import (
	"github.com/secnex/ms-toolbox/webhook-finder/api"
)

type TeamsAPIClient struct {
	GraphClient *api.MsGraph
	ApiBuilder  *TeamsAPIBuilder
}

func NewTeamsAPIClient(g *api.MsGraph, b *TeamsAPIBuilder) *TeamsAPIClient {
	return &TeamsAPIClient{
		GraphClient: g,
		ApiBuilder:  b,
	}
}

func (t *TeamsAPIClient) ListTeams() (*ListTeamsResponse, error) {
	var teams []Team
	pagination := true
	url := t.ApiBuilder.ListTeamsWithLimit(2)
	nextLink := url
	for pagination {
		resp, err := t.GraphClient.GET(nextLink)
		if err != nil {
			return nil, err
		}
		listTeamsResponse, err := t.ApiBuilder.GetListTeamsResponse(resp)
		if err != nil {
			return nil, err
		}
		teams = append(teams, listTeamsResponse.Value...)
		if listTeamsResponse.ODataNextLink == "" {
			pagination = false
		} else {
			nextLink = listTeamsResponse.ODataNextLink
		}
	}
	return &ListTeamsResponse{
		Value: teams,
	}, nil
}
