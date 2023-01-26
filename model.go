package models

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4/db"
	_ "firebase.google.com/go/v4/db"
	"net/http"
)

type Game struct {
	Players            []User                 `json:"players"`
	Word               string                 `json:"word"`
	Definition         string                 `json:"Definition"`
	MissingLetterIndex int                    `json:"missingLetterIndex"`
	GameId             string                 `json:"gameId"`
	RoundTime          int                    `json:"roundTime"`
	PlayIndex          int                    `json:"playIndex"`
	PlayerTurn         User                   `json:"playerTurn"`
	PlayOrder          map[int]User           `json:"playOrder"`
	Plays              map[string]PlayerMoves `json:"plays"`
	Player             User                   `json:"player"`
	PlayDirection      string                 `json:"playDirection"`
	Barriers           []TileLocation         `json:"barriers"`
	Obstacles          []TileLocation         `json:"obstacles"`
	Rewards            []TileLocation         `json:"rewards"`
	Time               string                 `json:"time"`
}

type PlayerMoves struct {
	Id            string         `json:"id"`
	GameId        string         `json:"gameId"`
	Player        User           `json:"player"`
	TileLocations []TileLocation `json:"tileLocations"`
	PlayIndex     int            `json:"playIndex"`
	Color         int            `json:"color"`
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
	PLayer     User   `json:"player"`
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
	Id        string `json:"id"`
	Privacy   string `json:"privacy"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	CreatedBy User   `json:"createdBy"`
	Users     []User `json:"users"`
}

type Lobby struct {
	Public  []Room `json:"public"`
	Private []Room `json:"private"`
}

type Ticket struct {
	Id               string       `json:"id"`
	GameType         string       `json:"gameType"`
	RoomType         string       `json:"roomType"`
	IsActive         bool         `json:"isActive"`
	RoomId           string       `json:"roomId"`
	IsMatchTicket    bool         `json:"isMatchTicket"`
	Created          string       `json:"created"`
	Expires          string       `json:"expires"`
	CreatedBy        User         `json:"createdBy"`
	Invitees         []User       `json:"invitees"`
	AcceptedBy       []User       `json:"acceptedBy,omitempty"`
	RejectedBy       []User       `json:"rejectedBy,omitempty"`
	IsBeingProcessed bool         `json:"isBeingProcessed"`
	InvitationSent   bool         `json:"invitationSent"`
	Status           TicketStatus `json:"status"`
	GameId           string       `json:"gameId,omitempty"`
	Capacity         int32        `json:"capacity,omitempty"`
}

func (t *Ticket) String() (string, error) {

	var users *[]User
	for _, invitee := range t.Invitees {
		u, errInner := json.Marshal(&invitee)
		if errInner != nil {
			continue
		}
		var usr *User
		errInner2 := json.Unmarshal(u, &usr)
		if errInner2 != nil {
			continue
		}
		*users = append(*users, *usr)
	}
	js, err := json.Marshal(&t)
	if err != nil {
		return "", err
	}

	return string(js), nil
}

func (t *Ticket) FromFbRef(ref *db.Ref) error {
	err := ref.Get(context.Background(), &t)
	if err != nil {
		//could not find invitation
		return err
	}
	return nil
}

func (t *Ticket) FromBody(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return err
	}
	return nil
}

type Match struct {
	Id       string `json:"id"`
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
	id           string `json:"id"`
	ResponseType string `json:"ResponseType"`
	Accepted     bool   `json:"Accepted"`
	TicketId     string `json:"TicketId"`
	Player       User   `json:"player"`
}

type WordSubmittedResponse struct {
	id                 string         `json:"id"`
	Score              int            `json:"score"`
	Word               string         `json:"word"`
	GameId             string         `json:"gameId"`
	MissingLetterIndex int            `json:"missingLetterIndex"`
	PlayIndex          int            `json:"playIndex"`
	PlayerTurn         User           `json:"playerTurn"`
	WordIsGood         bool           `json:"wordIsGood"`
	PlayDirection      string         `json:"playDirection"`
	Leader             User           `json:"leader"`
	TileLocations      []TileLocation `json:"tileLocations"`
	Definition         string         `json:"Definition"`
	GameOver           bool           `json:"gameOver"`
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

type LobbyRoomAccessRequest struct {
	RoomId string `json:"roomId"`
	UserId string `json:"userId"`
}

type LobbyRoomCrudRequest struct {
	RoomName string `json:"roomName"`
	RoomId   string `json:"roomId"`
	User     User   `json:"user"`
}

type TicketStatus string

const (
	New    TicketStatus = "New"
	Staged TicketStatus = "Staged"
	Ready  TicketStatus = "Ready"
	Active TicketStatus = "Active"
	Ended  TicketStatus = "Ended"
)
