package replay

import (
	"github.com/replay-api/replay-common/pkg/replay/events"
	shared "github.com/resource-ownership/go-common/pkg/common"
)

// Re-export shared types for backward compatibility
type BaseEntity = shared.BaseEntity
type ResourceOwner = shared.ResourceOwner
type IntendedAudienceKey = shared.IntendedAudienceKey
type VisibilityTypeKey = shared.VisibilityTypeKey
type ResourceType = shared.ResourceType

// Re-export shared constants
const (
	ALLOW = shared.ALLOW
	DENY  = shared.DENY

	PublicVisibilityTypeKey    = shared.PublicVisibilityTypeKey
	RestrictedVisibilityTypeKey = shared.RestrictedVisibilityTypeKey
	PrivateVisibilityTypeKey    = shared.PrivateVisibilityTypeKey
	CustomVisibilityTypeKey     = shared.CustomVisibilityTypeKey

	UserAudienceIDKey              = shared.UserAudienceIDKey
	GroupAudienceIDKey             = shared.GroupAudienceIDKey
	ClientApplicationAudienceIDKey = shared.ClientApplicationAudienceIDKey
	TenantAudienceIDKey            = shared.TenantAudienceIDKey
)

// Re-export shared error types
type ErrUnauthorized = shared.ErrUnauthorized
type ErrForbidden = shared.ErrForbidden
type ErrNotFound = shared.ErrNotFound
type ErrAlreadyExists = shared.ErrAlreadyExists
type ErrInvalidInput = shared.ErrInvalidInput
type ErrBadRequest = shared.ErrBadRequest

// Re-export shared error constructors
var (
	NewErrUnauthorized  = shared.NewErrUnauthorized
	NewErrForbidden     = shared.NewErrForbidden
	NewErrNotFound      = shared.NewErrNotFound
	NewErrAlreadyExists = shared.NewErrAlreadyExists
	NewErrInvalidInput  = shared.NewErrInvalidInput
	NewErrBadRequest    = shared.NewErrBadRequest
)

// Re-export shared error checkers
var (
	IsUnauthorizedError  = shared.IsUnauthorizedError
	IsForbiddenError     = shared.IsForbiddenError
	IsNotFoundError      = shared.IsNotFoundError
	IsBadRequestError    = shared.IsBadRequestError
	IsInvalidInputError  = shared.IsInvalidInputError
)

// Re-export shared functions
var (
	NewEntity            = shared.NewEntity
	NewUnrestrictedEntity = shared.NewUnrestrictedEntity
	NewRestrictedEntity   = shared.NewRestrictedEntity
	NewPrivateEntity      = shared.NewPrivateEntity
	NewResourceOwner      = shared.NewResourceOwner
	CanAccessResource     = shared.CanAccessResource
	GetResourceOwner      = shared.GetResourceOwner
)

// Re-export shared constants
var (
	TeamPROTenantID    = shared.TeamPROTenantID
	TeamPROAppClientID = shared.TeamPROAppClientID
	ServerClientID     = shared.ServerClientID
	AudienceKey        = shared.AudienceKey
)

// Gaming-specific types
type StatType = string

const (
	ClutchStatTypeKey      StatType = "Clutch"
	EconomyStatTypeKey     StatType = "Economy"
	StrategyStatTypeKey    StatType = "Strategy"
	PlayerStatTypeKey      StatType = "Player"
	PositioningStatTypeKey StatType = "Positioning"
	UtilityStatTypeKey     StatType = "Utility"
	BattleStatTypeKey      StatType = "Battle"
	GameSenseStatTypeKey   StatType = "Game Sense"
	HighlightStatTypeKey   StatType = "Highlight"
	AreaStatTypeKey        StatType = "Area"
)

type RegionIDKey = string

const (
	SouthAmerica_RegionIDKey RegionIDKey = "SA"
	NorthAmerica_RegionIDKey RegionIDKey = "NA"
	Asia_RegionIDKey         RegionIDKey = "AS"
	Global_RegionIDKey       RegionIDKey = "GL"
)

// Re-export event types
type EventIDKey = events.EventIDKey

// Re-export event constants
const (
	Event_MatchStartID           = events.Event_MatchStartID
	Event_RoundMVPAnnouncementID = events.Event_RoundMVPAnnouncementID
	Event_RoundEndID             = events.Event_RoundEndID
	Event_FragOrScoreID          = events.Event_FragOrScoreID
	Event_WeaponFireID           = events.Event_WeaponFireID
	Event_HitID                  = events.Event_HitID
	Event_GenericGameEventID     = events.Event_GenericGameEventID
	Event_ClutchStartID          = events.Event_ClutchStartID
	Event_ClutchProgressID       = events.Event_ClutchProgressID
	Event_ClutchEndID            = events.Event_ClutchEndID
	Event_Economy                = events.Event_Economy
)