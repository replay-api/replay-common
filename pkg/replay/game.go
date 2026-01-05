package replay

import (
	fps_events "github.com/replay-api/replay-common/pkg/replay/events/game/fps"
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
	Events []fps_events.EventIDKey `json:"in_game_events"` // Events is a map of SUPPORTED/IMPLEMENTED in-game events to their corresponding event names.
}

func mapCSEvents() []fps_events.EventIDKey {
	return []fps_events.EventIDKey{
		fps_events.Event_MatchStartID,
		fps_events.Event_RoundMVPAnnouncementID,
		fps_events.Event_RoundEndID,
		fps_events.Event_GenericGameEventID,
		fps_events.Event_ClutchStartID,
		fps_events.Event_ClutchProgressID,
		fps_events.Event_ClutchEndID,
		fps_events.Event_Economy,
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