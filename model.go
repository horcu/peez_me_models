package models

import "encoding/json"

type Game struct {
	Players            []User          `json:"players"`
	Word               string          `json:"word"`
	Definition         string          `json:"Definition"`
	MissingLetterIndex int             `json:"missingLetterIndex"`
	GameId             string          `json:"gameId"`
	RoundTime          int             `json:"roundTime"`
	PlayIndex          int             `json:"playIndex"`
	PlayerTurnId       string          `json:"playerTurnId"`
	Plays              map[string]Play `json:"plays"`
	LeaderId           string          `json:"leaderId"`
	PlayDirection      string          `json:"playDirection"`
	Barriers           []string        `json:"barriers"`
	Obstacles          []string        `json:"obstacles"`
	Rewards            []string        `json:"rewards"`
}

type Play struct {
	Word          string         `json:"word"`
	GameId        string         `json:"gameId"`
	UserId        string         `json:"userId"`
	TileLocations []TileLocation `json:"tileLocations"`
	Definition    string         `json:"definition"`
	PlayDirection string         `json:"playDirection"`
	PlayIndex     int            `json:"playIndex"`
}

type WordDefinition struct {
	Word       string `json:"word"`
	Valid      bool   `json:"valid"`
	Definition string `json:"Definition"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type TileLocation struct {
	Index      int    `json:"index"`
	Letter     string `json:"letter"`
	AreaName   string `json:"areaName"`
	UserId     string `json:"userId"`
	IsSelected bool   `json:"isSelected"`
}

type Word struct {
	Word string `json:"word"`
}

type PlayArea struct {
	Lobby Lobby `json:"lobby"`
}

type User struct {
	Available bool   `json:"available"`
	FcmToken  string `json:"fcmToken"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	PhotoURL  string `json:"photoUrl"`
}

type Room struct {
	Privacy   string `json:"privacy"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	CreatedBy string `json:"createdBy"`
	Users     []User `json:"users"`
}

type Lobby struct {
	Public  []Room `json:"public"`
	Private []Room `json:"private"`
}

type Ticket struct {
	Id               string `json:"id"`
	CreatedBy        string `json:"createdBy"`
	GameType         string `json:"gameType"`
	RoomType         string `json:"roomType"`
	IsActive         bool   `json:"isActive"`
	Room             string `json:"room"`
	IsMatchTicket    bool   `json:"isMatchTicket"`
	Created          string `json:"created"`
	Expires          string `json:"expires"`
	Invitees         []User `json:"invitees"`
	AcceptedBy       []User `json:"acceptedBy"`
	IsBeingProcessed bool   `json:"isBeingProcessed"`
	InvitationSent   bool   `json:"invitationSent"`
}

func (t *Ticket) String() (string, error) {

	var users []User
	for _, invitee := range t.Invitees {
		u, errInner := json.Marshal(&invitee)
		if errInner != nil {
			continue
		}
		var usr User
		errInner2 := json.Unmarshal(u, &usr)
		if errInner2 != nil {
			continue
		}
		users = append(users, usr)
	}
	js, err := json.Marshal(&t)
	if err != nil {
		return "", err
	}

	return string(js), nil
}

type Match struct {
	Players  []User `json:"players"`
	Room     string `json:"room"`
	Created  string `json:"created"`
	GameType string `json:"gameType"`
	RoomType string `json:"roomType"`
}

type InviteResponse struct {
	Message string `json:"message"`
	Ticket  Ticket `json:"ticket"`
}

type Pool struct {
	Tickets []Ticket `json:"tickets"`
	Name    string   `json:"name"`
}

type TemplateData struct {
	Service  string
	Revision string
}

type TicketResponse struct {
	ResponseType string `json:"ResponseType"`
	Accepted     bool   `json:"Accepted"`
	TicketId     string `json:"TicketId"`
	UserId       string `json:"UserId"`
}

type WordSubmittedResponse struct {
	Score              int            `json:"score"`
	Word               string         `json:"word"`
	GameId             string         `json:"gameId"`
	MissingLetterIndex int            `json:"missingLetterIndex"`
	PlayIndex          int            `json:"playIndex"`
	PlayerTurnId       string         `json:"playerTurnId"`
	WordIsGood         bool           `json:"wordIsGood"`
	PlayDirection      string         `json:"playDirection"`
	LeaderId           string         `json:"leaderId"`
	TileLocations      []TileLocation `json:"tileLocations"`
	Definition         string         `json:"Definition"`
}

type DeleteRequest struct {
	UserId string `json:"userId"`
	GameId string `json:"gameId"`
}

type PeezMeEvent struct {
	Type  string      `json:"@type,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Delta Ticket      `json:"delta,omitempty"`
}

type LobbyRoomRequest struct {
	RoomId string `json:"id"`
	UserId string `json:"userId"`
}
