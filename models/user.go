package models

import (
	"github.com/bnch/bancho/packets"
	"github.com/jinzhu/gorm"
)

// User Permissions
const (
	PermissionBanned = 1 << iota
	PermissionModerator
	PermissionAdmin
)

// User is an user on bancho.
type User struct {
	ID          int
	Username    string `sql:"size:20"`
	Permissions uint32
	Email       string
	Password    string
}

// IsBanned tells whether a user is banned.
func (u User) IsBanned() bool {
	return u.Permissions&PermissionBanned != 0
}

// IsModerator tells whether a user is a moderator.
func (u User) IsModerator() bool {
	return u.Permissions&PermissionModerator != 0
}

// IsAdmin tells whether a user is an admin.
func (u User) IsAdmin() bool {
	return u.Permissions&PermissionAdmin != 0
}

// GetColour gets an user's colour in the chat.
func (u User) GetColour() byte {
	switch {
	case u.IsModerator():
		return packets.ColourMod
	case u.IsAdmin():
		return packets.ColourAdmin
	default:
		return packets.ColourSupporter
	}
}

// RelatedUserStats finds the UserStats of an user.
func (u User) RelatedUserStats(db *gorm.DB) UserStats {
	user := UserStats{}
	db.First(&user, u.ID)
	return user
}
