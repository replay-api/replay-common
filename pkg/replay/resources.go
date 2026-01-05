package replay

import (
	"fmt"
	"strings"

	shared "github.com/resource-ownership/go-common/pkg/common"
)

// Gaming-specific resource types
const (
	ResourceTypeGameEvent      shared.ResourceType = "GameEvents"
	ResourceTypeBadge          shared.ResourceType = "Badges"
	ResourceTypeReplayFile     shared.ResourceType = "ReplayFiles"
	ResourceTypeMatch          shared.ResourceType = "Matches"
	ResourceTypeRound          shared.ResourceType = "Rounds"
	ResourceTypeGame           shared.ResourceType = "Games"
	ResourceTypePlayerProfile  shared.ResourceType = "Players"        // composition of user
	ResourceTypePlayerMetadata shared.ResourceType = "PlayerMetadata" // composition of user
	ResourceTypeTeam           shared.ResourceType = "Teams"          // specification of user group
	ResourceTypeSquad          shared.ResourceType = "Squads"         // specification of user group
	ResourceTypeChannel        shared.ResourceType = "Channels"       // specification of user group
	ResourceTypeLeague         shared.ResourceType = "Leagues"        // specification of user group
	ResourceTypeTournament     shared.ResourceType = "Tournaments"    // specification of user group
	ResourceTypeProfile        shared.ResourceType = "Profiles"       // specification of user group
	ResourceTypeMembership     shared.ResourceType = "Memberships"    // specification of user group
	ResourceTypeFriends        shared.ResourceType = "Friends"
	ResourceTypeChallenge      shared.ResourceType = "Challenges" // Bug reports, VAR, round restart requests
)

var GamingResourceTypes = []shared.ResourceType{
	ResourceTypeGameEvent,
	ResourceTypeBadge,
	ResourceTypeReplayFile,
	ResourceTypeMatch,
	ResourceTypeRound,
	ResourceTypeGame,
	ResourceTypePlayerProfile,
	ResourceTypePlayerMetadata,
	ResourceTypeTeam,
	ResourceTypeSquad,
	ResourceTypeChannel,
	ResourceTypeLeague,
	ResourceTypeTournament,
	ResourceTypeProfile,
	ResourceTypeMembership,
	ResourceTypeFriends,
	ResourceTypeChallenge,
}

var GamingResourceKeyMap = map[shared.ResourceType]string{
	ResourceTypeGameEvent:     "game_event_id",
	ResourceTypeBadge:         "badge_id",
	ResourceTypeReplayFile:    "replay_file_id",
	ResourceTypeMatch:         "match_id",
	ResourceTypeRound:         "round_id",
	ResourceTypeGame:          "game_id",
	ResourceTypePlayerProfile: "player_id",
	ResourceTypeTeam:          "team_id",
	ResourceTypeProfile:       "profile_id",
	ResourceTypeChallenge:     "challenge_id",
}

func GetGamingResourceFieldID(resourcePart string) (string, error) {
	for k, v := range GamingResourceKeyMap {
		if strings.EqualFold(fmt.Sprint(k), resourcePart) {
			return v, nil
		}
	}

	return "", fmt.Errorf("failed to parse ResourceIDField: Unknown gaming resource %s", resourcePart)
}