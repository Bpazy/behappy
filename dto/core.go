package dto

import "github.com/Bpazy/behappy/ent"

func FromSubscriptionMatch(sm *ent.SubscriptionMatch) *MatchPlayerDto {
	if sm == nil {
		return nil
	}
	return &MatchPlayerDto{
		MatchID:      sm.MatchID,
		PlayerID:     sm.PlayerID,
		PlayerSlot:   sm.PlayerSlot,
		RadiantWin:   sm.RadiantWin,
		Duration:     sm.Duration,
		GameMode:     sm.GameMode,
		LobbyType:    sm.LobbyType,
		HeroID:       sm.HeroID,
		StartTime:    sm.StartTime,
		Version:      sm.Version,
		Kills:        sm.Kills,
		Deaths:       sm.Deaths,
		Assists:      sm.Assists,
		Skill:        sm.Skill,
		LeaverStatus: sm.LeaverStatus,
		PartySize:    sm.PartySize,
	}
}

type MatchPlayerDto struct {
	MatchID      int64  `json:"match_id" gorm:"index"`
	PlayerID     string `gorm:"index"`
	PlayerSlot   int    `json:"player_slot" gorm:"not null"`
	RadiantWin   bool   `json:"radiant_win" gorm:"not null"`
	Duration     int    `json:"duration" gorm:"not null"` // Seconds
	GameMode     int    `json:"game_mode" gorm:"not null"`
	LobbyType    int    `json:"lobby_type" gorm:"not null"`
	HeroID       int    `json:"hero_id" gorm:"not null"`
	StartTime    int    `json:"start_time" gorm:"not null"`
	Version      int    `json:"version" gorm:"not null"`
	Kills        int    `json:"kills" gorm:"not null"`
	Deaths       int    `json:"deaths" gorm:"not null"`
	Assists      int    `json:"assists" gorm:"not null"`
	Skill        *int   `json:"skill"`
	LeaverStatus int    `json:"leaver_status" gorm:"not null"`
	PartySize    int    `json:"party_size" gorm:"not null"`
}

func (m MatchPlayerDto) IsWin() bool {
	if m.PlayerSlot < 127 {
		return m.RadiantWin
	}
	return !m.RadiantWin
}

func (m MatchPlayerDto) SkillString() string {
	if m.Skill == nil {
		return "Unknown"
	}
	switch *m.Skill {
	case 3:
		return "Very High"
	case 2:
		return "High"
	}
	return "Normal"
}

func (m MatchPlayerDto) DurationMinutes() int {
	return m.Duration / 60
}

type SubscriptionDto struct {
	GroupID  int
	PlayerID string
	Alias    string
}

func (sp SubscriptionDto) Name() string {
	if sp.Alias != "" {
		return sp.Alias
	}
	return sp.PlayerID
}

type HeroDto struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	LocalizedName string `json:"localized_name"`
}

type SteamApiResult struct {
	Result struct {
		Heroes []HeroDto `json:"heroes"`
		Status int       `json:"status"`
		Count  int       `json:"count"`
	} `json:"result"`
}
