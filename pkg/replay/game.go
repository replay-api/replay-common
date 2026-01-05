package replay

import (
	"github.com/replay-api/replay-common/pkg/replay/events"
)

// Gaming-specific types
type GameIDKey string

const (
	CS2_GAME_ID   GameIDKey = "cs2"
	CSGO_GAME_ID  GameIDKey = "csgo"
	VLRNT_GAME_ID GameIDKey = "vlrnt"
)

type TickIDType float64

type Game struct {
	ID     GameIDKey        `json:"id"`             // ID is the unique identifier of the game.
	Name   string           `json:"name"`           // Name is the name of the game.
	Events []events.EventIDKey `json:"in_game_events"` // Events is a map of SUPPORTED/IMPLEMENTED in-game events to their corresponding event names.
}

func mapCSEvents() []events.EventIDKey {
	return []events.EventIDKey{
		events.Event_MatchStartID,
		events.Event_RoundMVPAnnouncementID,
		events.Event_RoundEndID,
		events.Event_GenericGameEventID,
		events.Event_ClutchStartID,
		events.Event_ClutchProgressID,
		events.Event_ClutchEndID,
		events.Event_Economy,
	}
}

// Game instances
var (
	CS2 = &Game{
		ID:     CS2_GAME_ID,
		Name:   "Counter-Strike: 2",
		Events: mapCSEvents(),
	}

	CSGO = &Game{
		ID:     CSGO_GAME_ID,
		Name:   "Counter-Strike: Global Offensive",
		Events: mapCSEvents(),
	}
)