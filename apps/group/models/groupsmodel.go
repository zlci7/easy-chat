package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupsModel = (*customGroupsModel)(nil)

type (
	// GroupsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupsModel.
	GroupsModel interface {
		groupsModel
		// 新增：事务创建群组和群主
		InsertWithGroupAndMember(ctx context.Context, group *Groups, member *GroupMembers) error
		// 新增：增加群成员数量
		IncrMemberCount(ctx context.Context, groupId uint64, delta int) error
		// 新增：减少群成员数量
		DecrMemberCount(ctx context.Context, groupId uint64, delta int) error
		// 🔥 新增：事务添加成员（处理重新加入和新加入两种情况）
		AddMemberWithIncrCount(ctx context.Context, member *GroupMembers, isRejoin bool) error
	}

	customGroupsModel struct {
		*defaultGroupsModel
	}
)

// NewGroupsModel returns a model for the database table.
func NewGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupsModel {
	return &customGroupsModel{
		defaultGroupsModel: newGroupsModel(conn, c, opts...),
	}
}

// 事务创建群组和群主：优化后的实现
func (m *customGroupsModel) InsertWithGroupAndMember(ctx context.Context, group *Groups, member *GroupMembers) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 1. 插入群组 - 注意：groups 是保留字，必须用反引号
		query := "insert into `groups` (`group_id`,`name`,`owner_uid`,`type`,`avatar`,`status`,`description`,`notice`,`max_members`,`member_count`,`join_type`,`invite_confirm`,`mute_all`,`create_time`,`update_time`,`deleted_at`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

		_, err := session.ExecCtx(ctx, query,
			group.GroupId,       // 1
			group.Name,          // 2
			group.OwnerUid,      // 3
			group.Type,          // 4
			group.Avatar,        // 5
			group.Status,        // 6
			group.Description,   // 7
			group.Notice,        // 8
			group.MaxMembers,    // 9
			group.MemberCount,   // 10
			group.JoinType,      // 11
			group.InviteConfirm, // 12
			group.MuteAll,       // 13
			group.CreateTime,    // 14
			group.UpdateTime,    // 15
			group.DeletedAt,     // 16
		)
		if err != nil {
			return err
		}

		// 2. 插入群成员 - 使用生成的 SQL
		// 更好的做法：用 groupMembersRowsExpectAutoSet 变量
		queryMember := "insert into group_members (group_id, user_id, role, status, join_time, update_time) values (?, ?, ?, ?, ?, ?)"
		_, err = session.ExecCtx(ctx, queryMember,
			member.GroupId,
			member.UserId,
			member.Role,
			member.Status, // 🔥 别忘了这个
			member.JoinTime,
			member.UpdateTime, // 🔥 别忘了这个
		)
		return err
	})
}

// IncrMemberCount 增加群成员数量
func (m *customGroupsModel) IncrMemberCount(ctx context.Context, groupId uint64, delta int) error {
	// 清除缓存
	groupsGroupIdKey := fmt.Sprintf("%s%v", cacheGroupsGroupIdPrefix, groupId)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		// 注意：groups 是保留字，必须用反引号
		query := "update `groups` set member_count = member_count + ? where group_id = ?"
		return conn.ExecCtx(ctx, query, delta, groupId)
	}, groupsGroupIdKey)

	return err
}

// DecrMemberCount 减少群成员数量
func (m *customGroupsModel) DecrMemberCount(ctx context.Context, groupId uint64, delta int) error {
	// 清除缓存
	groupsGroupIdKey := fmt.Sprintf("%s%v", cacheGroupsGroupIdPrefix, groupId)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		// 注意：groups 是保留字，必须用反引号
		query := "update `groups` set member_count = member_count - ? where group_id = ? and member_count >= ?"
		return conn.ExecCtx(ctx, query, delta, groupId, delta)
	}, groupsGroupIdKey)

	return err
}

// AddMemberWithIncrCount 事务：添加成员并增加成员数
// isRejoin: true表示重新加入（更新记录），false表示新加入（插入记录）
func (m *customGroupsModel) AddMemberWithIncrCount(ctx context.Context, member *GroupMembers, isRejoin bool) error {
	// 准备缓存键
	groupsGroupIdKey := fmt.Sprintf("%s%v", cacheGroupsGroupIdPrefix, member.GroupId)
	groupMembersGroupIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheGroupMembersGroupIdUserIdPrefix, member.GroupId, member.UserId)

	// 执行事务
	err := m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 1. 更新或插入成员记录
		if isRejoin {
			// 重新加入：更新现有记录
			query := `UPDATE group_members 
                     SET status = ?, role = ?, join_time = ?, update_time = ?, leave_time = 0
                     WHERE group_id = ? AND user_id = ?`
			_, err := session.ExecCtx(ctx, query,
				member.Status, member.Role, member.JoinTime, member.UpdateTime,
				member.GroupId, member.UserId)
			if err != nil {
				return fmt.Errorf("update member failed: %w", err)
			}
		} else {
			// 新加入：插入新记录
			query := `INSERT INTO group_members 
                     (group_id, user_id, role, status, inviter_uid, join_time, update_time)
                     VALUES (?, ?, ?, ?, ?, ?, ?)`
			_, err := session.ExecCtx(ctx, query,
				member.GroupId, member.UserId, member.Role, member.Status,
				member.InviterUid, member.JoinTime, member.UpdateTime)
			if err != nil {
				return fmt.Errorf("insert member failed: %w", err)
			}
		}

		// 2. 增加群成员数量（注意：groups 是保留字，必须用反引号）
		query := `UPDATE ` + "`groups`" + ` 
                 SET member_count = member_count + 1, update_time = ?
                 WHERE group_id = ?`
		_, err := session.ExecCtx(ctx, query, member.UpdateTime, member.GroupId)
		if err != nil {
			return fmt.Errorf("incr member count failed: %w", err)
		}

		return nil
	}) // 🔥 注意：这里只有2个参数

	// 🔥 事务成功后，手动清除缓存
	if err == nil {
		m.DelCacheCtx(ctx, groupsGroupIdKey)
		m.DelCacheCtx(ctx, groupMembersGroupIdUserIdKey)
	}

	return err
}
