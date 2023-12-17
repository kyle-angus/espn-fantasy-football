package espnclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kyle-angus/espn-fantasy-football-client/pkg/config"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/response"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/team"
)

const (
	BASE_URL = "https://fantasy.espn.com/apis/v3/games/ffl/seasons/"
)

type EspnClient struct {
	config *config.EspnConfig
}

func New(config *config.EspnConfig) *EspnClient {
	if config == nil {
		fmt.Println("config is nil")
		return nil
	}

	if config.Season == 0 {
		fmt.Println("season is empty")
		return nil
	}

	if config.LeagueId == 0 {
		fmt.Println("leagueId is empty")
		return nil
	}

	if config.EspnS2 == "" {
		fmt.Println("espns2 is empty")
		return nil
	}

	if config.SWID == "" {
		fmt.Println("swid is empty")
		return nil
	}

	return &EspnClient{
		config,
	}
}

func (c *EspnClient) GetLeagueData() response.EspnResponse {
	api := BASE_URL + fmt.Sprintf("%d/segments/0/leagues/%d?view=modular&view=mNav&view=mMatchupScore&view=mScoreboard&view=mStatus&view=mSettings&view=mTeam&view=mPendingTransactions", c.config.Season, c.config.LeagueId)

	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Fatalf("Error: GetLeagueInfo: %s\n", err)
	}
	request.Header.Add("Cookie", fmt.Sprintf("espn_s2=%s; SWID=%s;", c.config.EspnS2, c.config.SWID))

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error: GetLeagueInfo: %s\n", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error: GetLeagueInfo: %s\n", err)
	}

	var data response.EspnResponse
	json.Unmarshal(body, &data)

	return data
}

func (c *EspnClient) GetTeams() []team.Team {
	response := c.GetLeagueData()
	return response.Teams
}
