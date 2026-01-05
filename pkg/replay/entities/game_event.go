package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	fps_events "github.com/replay-api/replay-common/pkg/replay/events/game/fps"
	"github.com/replay-api/replay-common/pkg/replay"
	shared "github.com/resource-ownership/go-common/pkg/common"
)

type GameEvent struct {
	// header/meta
	ID       uuid.UUID           `json:"id" bson:"_id"`
	Type     fps_events.EventIDKey   `json:"type" bson:"type"`
	GameID   replay.GameIDKey    `json:"game_id" bson:"game_id"`
	MatchID  uuid.UUID           `json:"match_id" bson:"match_id"`
	TickID   replay.TickIDType   `json:"tick_id" bson:"tick_id"`
	GameTime time.Duration       `json:"event_time" bson:"event_time"` // CurrentTime

	// data
	Payload  interface{}                           `json:"-" bson:"payload"`
	Entities map[shared.ResourceType][]interface{} `json:"-" bson:"-"`
	Stats    map[replay.StatType][]interface{}     `json:"stats" bson:"stats"`

	// default/trail
	ResourceOwner shared.ResourceOwner `json:"-" bson:"resource_owner"`
	CreatedAt     time.Time            `json:"-" bson:"created_at"`
}

func NewGameEvent[T any](matchID uuid.UUID, tickID replay.TickIDType, gameTime time.Duration, eventType fps_events.EventIDKey, payload T, entities map[shared.ResourceType][]interface{}, stats map[replay.StatType][]interface{}, res shared.ResourceOwner) *GameEvent {
	return &GameEvent{
		ID:            uuid.New(),
		GameID:        replay.CS2_GAME_ID, // TODO: refact => quando aplicavel para go/vlr
		MatchID:       matchID,
		TickID:        tickID,
		Type:          eventType,
		GameTime:      gameTime,
		Payload:       payload,
		Entities:      entities,
		Stats:         stats,
		ResourceOwner: res,
		CreatedAt:     time.Now(),
	}
}

func (e GameEvent) GetID() uuid.UUID {
	return e.ID
}

func (e GameEvent) GetPlayerIDs() ([]uuid.UUID, error) {
	if playerEntities, ok := e.Entities[replay.ResourceTypePlayerProfile]; ok {
		playerIDs := make([]uuid.UUID, 0, len(playerEntities))
		for _, entity := range playerEntities {
			if id, ok := entity.(uuid.UUID); ok {
				playerIDs = append(playerIDs, id)
			}
		}
		return playerIDs, nil
	}
	return nil, fmt.Errorf("no player entities found")
}