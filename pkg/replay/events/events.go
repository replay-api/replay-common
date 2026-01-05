package events

type EventIDKey string

const (
	Event_MatchStartID           EventIDKey = "MatchStart"
	Event_RoundMVPAnnouncementID EventIDKey = "RoundMVPAnnouncement"
	Event_RoundEndID             EventIDKey = "RoundEndID"
	Event_FragOrScoreID          EventIDKey = "FragOrScoreID"
	Event_WeaponFireID           EventIDKey = "WeaponFireID"
	Event_HitID                  EventIDKey = "HitID"
	Event_GenericGameEventID     EventIDKey = "GenericGameEvent"
	Event_ClutchStartID          EventIDKey = "ClutchStart"
	Event_ClutchProgressID       EventIDKey = "ClutchProgress"
	Event_ClutchEndID            EventIDKey = "ClutchEnd"
	Event_Economy                EventIDKey = "EconomyEvent"
)