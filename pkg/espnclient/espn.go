package espnclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kyle-angus/espn-fantasy-football-client/pkg/member"
	r "github.com/kyle-angus/espn-fantasy-football-client/pkg/response"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/team"
)

const (
	BASE_URL = "https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/seasons/"
)

type EspnClient struct {
	LeagueId uint
	EspnS2   string
	SWID     string
}

func New(leagueid uint, espns2 string, swid string) *EspnClient {
	if leagueid == 0 {
		fmt.Println("leagueId is empty")
		return nil
	}

	if espns2 == "" {
		fmt.Println("espns2 is empty")
		return nil
	}

	if swid == "" {
		fmt.Println("swid is empty")
		return nil
	}

	return &EspnClient{
		leagueid,
		espns2,
		swid,
	}
}

func (c *EspnClient) getLeagueData(season uint) (*r.EspnResponse, error) {
	api := BASE_URL + fmt.Sprintf("%d/segments/0/leagues/%d?view=modular&view=mNav&view=mMatchupScore&view=mScoreboard&view=mStatus&view=mSettings&view=mTeam&view=mPendingTransactions", season, c.LeagueId)

	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Printf("Error: GetLeagueInfo: %s\n", err)
		return nil, err
	}
	request.Header.Add("Cookie", fmt.Sprintf("espn_s2=%s; SWID=%s;", c.EspnS2, c.SWID))

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Printf("Error: GetLeagueInfo: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error: GetLeagueInfo: %s\n", err)
		return nil, err
	}

	var data r.EspnResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error: GetLeagueInfo: %s\n", err)
		return nil, err
	}

	return &data, nil
}

func (c *EspnClient) GetTeams(season uint) ([]team.Team, error) {
	response, err := c.getLeagueData(season)
	if err != nil {
		return nil, err
	}
	return response.Teams, nil
}

func (c *EspnClient) GetTeam(season uint, teamId int) (*team.Team, error) {
	response, err := c.getLeagueData(season)
	if err != nil {
		return nil, err
	}

	var team team.Team
	for _, t := range response.Teams {
		if t.Id == teamId {
			team = t
			break
		}
	}

	return &team, nil
}

func (c *EspnClient) GetMembers(season uint) ([]member.Member, error) {
	response, err := c.getLeagueData(season)
	if err != nil {
		return nil, err
	}

	return response.Members, nil
}

func (c *EspnClient) GetTeamByMemberId(season uint, memberId string) (*team.Team, error) {
	response, err := c.getLeagueData(season)
	if err != nil {
		return nil, err
	}

	var team team.Team
	for _, t := range response.Teams {
		if t.PrimaryOwner == memberId {
			team = t
			break
		}
	}

	return &team, nil
}
