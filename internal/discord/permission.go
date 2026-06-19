package discord

import "github.com/bwmarrin/discordgo"

type PermissionLevel int

const (
	Citizen PermissionLevel = iota
	Sentinel
	Pantheon
)

func GetPermissionLevel(s *discordgo.Session, guildID, userID string, sentinelRoleName, pantheonRoleName string) PermissionLevel {
	member, err := s.GuildMember(guildID, userID)
	if err != nil {
		return Citizen
	}

	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		return Citizen
	}

	roleNameByID := make(map[string]string)

	for _, role := range guildRoles {
		roleNameByID[role.ID] = role.Name
	}

	levelNow := Citizen
	for _, roleID := range member.Roles {
		roleName := roleNameByID[roleID]
		if roleName == pantheonRoleName {
			return Pantheon
		}
		if roleName == sentinelRoleName {
			levelNow = Sentinel
		}
	}
	return levelNow
}

func (p PermissionLevel) CanModerate() bool {
	return p == Sentinel || p == Pantheon
}

func (p PermissionLevel) String() string {
	switch p {
	case Pantheon:
		return "The Pantheon"
	case Sentinel:
		return "The Sentinel"
	default:
		return "Citizen"
	}
}
