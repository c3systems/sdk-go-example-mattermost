// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c3systems/c3-sdk-go-example-mattermost/model"
	"github.com/c3systems/c3-sdk-go-example-mattermost/store/storetest"
)

func TestChannelStore(t *testing.T) {
	StoreTest(t, storetest.TestChannelStore)
}

func TestChannelStoreInternalDataTypes(t *testing.T) {
	t.Run("NewChannelMemberFromModel", func(t *testing.T) { testNewChannelMemberFromModel(t) })
	t.Run("ChannelMemberWithSchemeRolesToModel", func(t *testing.T) { testChannelMemberWithSchemeRolesToModel(t) })
	t.Run("AllChannelMemberProcess", func(t *testing.T) { testAllChannelMemberProcess(t) })
}

func testNewChannelMemberFromModel(t *testing.T) {
	m := model.ChannelMember{
		ChannelId:     model.NewId(),
		UserId:        model.NewId(),
		Roles:         "channel_user channel_admin custom_role",
		LastViewedAt:  12345,
		MsgCount:      2,
		MentionCount:  1,
		NotifyProps:   model.StringMap{"key": "value"},
		LastUpdateAt:  54321,
		SchemeUser:    true,
		SchemeAdmin:   true,
		ExplicitRoles: "custom_role",
	}

	db := NewChannelMemberFromModel(&m)

	assert.Equal(t, m.ChannelId, db.ChannelId)
	assert.Equal(t, m.UserId, db.UserId)
	assert.Equal(t, m.LastViewedAt, db.LastViewedAt)
	assert.Equal(t, m.MsgCount, db.MsgCount)
	assert.Equal(t, m.MentionCount, db.MentionCount)
	assert.Equal(t, m.NotifyProps, db.NotifyProps)
	assert.Equal(t, m.LastUpdateAt, db.LastUpdateAt)
	assert.Equal(t, true, db.SchemeUser.Valid)
	assert.Equal(t, true, db.SchemeAdmin.Valid)
	assert.Equal(t, m.SchemeUser, db.SchemeUser.Bool)
	assert.Equal(t, m.SchemeAdmin, db.SchemeAdmin.Bool)
	assert.Equal(t, m.ExplicitRoles, db.Roles)
}

func testChannelMemberWithSchemeRolesToModel(t *testing.T) {
	t.Run("BasicProperties", func(t *testing.T) {
		// Test all the non-roles properties here.
		db := channelMemberWithSchemeRoles{
			ChannelId:                     model.NewId(),
			UserId:                        model.NewId(),
			Roles:                         "custom_role",
			LastViewedAt:                  12345,
			MsgCount:                      2,
			MentionCount:                  1,
			NotifyProps:                   model.StringMap{"key": "value"},
			LastUpdateAt:                  54321,
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		m := db.ToModel()

		assert.Equal(t, db.ChannelId, m.ChannelId)
		assert.Equal(t, db.UserId, m.UserId)
		assert.Equal(t, "custom_role channel_user channel_admin", m.Roles)
		assert.Equal(t, db.LastViewedAt, m.LastViewedAt)
		assert.Equal(t, db.MsgCount, m.MsgCount)
		assert.Equal(t, db.MentionCount, m.MentionCount)
		assert.Equal(t, db.NotifyProps, m.NotifyProps)
		assert.Equal(t, db.LastUpdateAt, m.LastUpdateAt)
		assert.Equal(t, db.SchemeUser.Bool, m.SchemeUser)
		assert.Equal(t, db.SchemeAdmin.Bool, m.SchemeAdmin)
		assert.Equal(t, db.Roles, m.ExplicitRoles)
	})

	// Example data *before* the Phase 2 migration has taken place.
	t.Run("Unmigrated_NoScheme_User", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "channel_user",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "channel_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Unmigrated_NoScheme_Admin", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "channel_admin channel_user",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "channel_admin channel_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Unmigrated_NoScheme_CustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Unmigrated_NoScheme_UserAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "channel_user custom_role",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "channel_user custom_role", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Unmigrated_NoScheme_AdminAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "channel_user channel_admin custom_role",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "channel_user channel_admin custom_role", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Unmigrated_NoScheme_NoRoles", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	// Example data *after* the Phase 2 migration has taken place.
	t.Run("Migrated_NoScheme_User", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "channel_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_NoScheme_Admin", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "channel_user channel_admin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_NoScheme_CustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_NoScheme_UserAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role channel_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_NoScheme_AdminAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role channel_user channel_admin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_NoScheme_NoRoles", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	// Example data with a channel scheme.
	t.Run("Migrated_ChannelScheme_User", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "cscheme_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_ChannelScheme_Admin", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "cscheme_user cscheme_admin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_ChannelScheme_CustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_ChannelScheme_UserAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role cscheme_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_ChannelScheme_AdminAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role cscheme_user cscheme_admin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_ChannelScheme_NoRoles", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	// Example data with a team scheme.
	t.Run("Migrated_TeamScheme_User", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "tscheme_channeluser", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamScheme_Admin", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "tscheme_channeluser tscheme_channeladmin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamScheme_CustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamScheme_UserAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role tscheme_channeluser", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamScheme_AdminAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role tscheme_channeluser tscheme_channeladmin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamScheme_NoRoles", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		cm := db.ToModel()

		assert.Equal(t, "", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	// Example data with a team and channel scheme.
	t.Run("Migrated_TeamAndChannelScheme_User", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "cscheme_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamAndChannelScheme_Admin", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "cscheme_user cscheme_admin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamAndChannelScheme_CustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamAndChannelScheme_UserAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role cscheme_user", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamAndChannelScheme_AdminAndCustomRole", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "custom_role",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "custom_role cscheme_user cscheme_admin", cm.Roles)
		assert.Equal(t, true, cm.SchemeUser)
		assert.Equal(t, true, cm.SchemeAdmin)
		assert.Equal(t, "custom_role", cm.ExplicitRoles)
	})

	t.Run("Migrated_TeamAndChannelScheme_NoRoles", func(t *testing.T) {
		db := channelMemberWithSchemeRoles{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		cm := db.ToModel()

		assert.Equal(t, "", cm.Roles)
		assert.Equal(t, false, cm.SchemeUser)
		assert.Equal(t, false, cm.SchemeAdmin)
		assert.Equal(t, "", cm.ExplicitRoles)
	})
}

func testAllChannelMemberProcess(t *testing.T) {
	t.Run("Unmigrated_User", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "channel_user",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "channel_user", roles)
	})

	t.Run("Unmigrated_Admin", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "channel_user channel_admin",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "channel_user channel_admin", roles)
	})

	t.Run("Unmigrated_Neither", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "", roles)
	})

	t.Run("Unmigrated_Custom", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "custom",
			SchemeUser:                    sql.NullBool{Valid: false, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: false, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "custom", roles)
	})

	t.Run("MigratedNoScheme_User", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "channel_user", roles)
	})

	t.Run("MigratedNoScheme_Admin", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "channel_user channel_admin", roles)
	})

	t.Run("MigratedNoScheme_Neither", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "", roles)
	})

	t.Run("MigratedChannelScheme_User", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		_, roles := db.Process()

		assert.Equal(t, "cscheme_user", roles)
	})

	t.Run("MigratedChannelScheme_Admin", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		_, roles := db.Process()

		assert.Equal(t, "cscheme_user cscheme_admin", roles)
	})

	t.Run("MigratedChannelScheme_Neither", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		_, roles := db.Process()

		assert.Equal(t, "", roles)
	})

	t.Run("MigratedTeamScheme_User", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "tscheme_channeluser", roles)
	})

	t.Run("MigratedTeamScheme_Admin", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "tscheme_channeluser tscheme_channeladmin", roles)
	})

	t.Run("MigratedTeamScheme_Neither", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "", roles)
	})

	t.Run("MigratedTeamAndChannelScheme_User", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		_, roles := db.Process()

		assert.Equal(t, "cscheme_user", roles)
	})

	t.Run("MigratedTeamAndChannelScheme_Admin", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: true},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		_, roles := db.Process()

		assert.Equal(t, "cscheme_user cscheme_admin", roles)
	})

	t.Run("MigratedTeamAndChannelScheme_Neither", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: false},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: true, String: "tscheme_channeluser"},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: true, String: "tscheme_channeladmin"},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: true, String: "cscheme_user"},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: true, String: "cscheme_admin"},
		}

		_, roles := db.Process()

		assert.Equal(t, "", roles)
	})

	t.Run("DeduplicationCheck", func(t *testing.T) {
		db := allChannelMember{
			Roles:                         "channel_user",
			SchemeUser:                    sql.NullBool{Valid: true, Bool: true},
			SchemeAdmin:                   sql.NullBool{Valid: true, Bool: false},
			TeamSchemeDefaultUserRole:     sql.NullString{Valid: false},
			TeamSchemeDefaultAdminRole:    sql.NullString{Valid: false},
			ChannelSchemeDefaultUserRole:  sql.NullString{Valid: false},
			ChannelSchemeDefaultAdminRole: sql.NullString{Valid: false},
		}

		_, roles := db.Process()

		assert.Equal(t, "channel_user", roles)
	})
}
