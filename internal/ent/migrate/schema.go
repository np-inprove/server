// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessoriesColumns holds the columns for the "accessories" table.
	AccessoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "points_required", Type: field.TypeInt},
		{Name: "institution_accessories", Type: field.TypeInt},
	}
	// AccessoriesTable holds the schema information for the "accessories" table.
	AccessoriesTable = &schema.Table{
		Name:       "accessories",
		Columns:    AccessoriesColumns,
		PrimaryKey: []*schema.Column{AccessoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accessories_institutions_accessories",
				Columns:    []*schema.Column{AccessoriesColumns[4]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DeadlinesColumns holds the columns for the "deadlines" table.
	DeadlinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "due_time", Type: field.TypeTime, Nullable: true},
		{Name: "deadline_author", Type: field.TypeInt},
		{Name: "group_deadlines", Type: field.TypeInt},
	}
	// DeadlinesTable holds the schema information for the "deadlines" table.
	DeadlinesTable = &schema.Table{
		Name:       "deadlines",
		Columns:    DeadlinesColumns,
		PrimaryKey: []*schema.Column{DeadlinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deadlines_users_author",
				Columns:    []*schema.Column{DeadlinesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "deadlines_groups_deadlines",
				Columns:    []*schema.Column{DeadlinesColumns[4]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "repeat_pattern", Type: field.TypeString, Nullable: true},
		{Name: "group_events", Type: field.TypeInt},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_groups_events",
				Columns:    []*schema.Column{EventsColumns[6]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ForumsColumns holds the columns for the "forums" table.
	ForumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "short_name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "group_forums", Type: field.TypeInt},
	}
	// ForumsTable holds the schema information for the "forums" table.
	ForumsTable = &schema.Table{
		Name:       "forums",
		Columns:    ForumsColumns,
		PrimaryKey: []*schema.Column{ForumsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "forums_groups_forums",
				Columns:    []*schema.Column{ForumsColumns[4]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ForumPostsColumns holds the columns for the "forum_posts" table.
	ForumPostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "mentioned_users_json", Type: field.TypeJSON},
		{Name: "forum_posts", Type: field.TypeInt},
		{Name: "forum_post_author", Type: field.TypeInt},
		{Name: "forum_post_children", Type: field.TypeInt, Nullable: true},
	}
	// ForumPostsTable holds the schema information for the "forum_posts" table.
	ForumPostsTable = &schema.Table{
		Name:       "forum_posts",
		Columns:    ForumPostsColumns,
		PrimaryKey: []*schema.Column{ForumPostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "forum_posts_forums_posts",
				Columns:    []*schema.Column{ForumPostsColumns[4]},
				RefColumns: []*schema.Column{ForumsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "forum_posts_users_author",
				Columns:    []*schema.Column{ForumPostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "forum_posts_forum_posts_children",
				Columns:    []*schema.Column{ForumPostsColumns[6]},
				RefColumns: []*schema.Column{ForumPostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "short_name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "institution_groups", Type: field.TypeInt},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_institutions_groups",
				Columns:    []*schema.Column{GroupsColumns[4]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupInviteLinksColumns holds the columns for the "group_invite_links" table.
	GroupInviteLinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"owner", "educator", "member"}},
		{Name: "group_invites", Type: field.TypeInt},
	}
	// GroupInviteLinksTable holds the schema information for the "group_invite_links" table.
	GroupInviteLinksTable = &schema.Table{
		Name:       "group_invite_links",
		Columns:    GroupInviteLinksColumns,
		PrimaryKey: []*schema.Column{GroupInviteLinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_invite_links_groups_invites",
				Columns:    []*schema.Column{GroupInviteLinksColumns[3]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "role", Type: field.TypeEnum, Enums: []string{"owner", "educator", "member"}},
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[1], GroupUsersColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_groups_group",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_users_users_user",
				Columns:    []*schema.Column{GroupUsersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// InstitutionsColumns holds the columns for the "institutions" table.
	InstitutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "short_name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: ""},
	}
	// InstitutionsTable holds the schema information for the "institutions" table.
	InstitutionsTable = &schema.Table{
		Name:       "institutions",
		Columns:    InstitutionsColumns,
		PrimaryKey: []*schema.Column{InstitutionsColumns[0]},
	}
	// InstitutionInviteLinksColumns holds the columns for the "institution_invite_links" table.
	InstitutionInviteLinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "educator", "member"}},
		{Name: "institution_invites", Type: field.TypeInt},
	}
	// InstitutionInviteLinksTable holds the schema information for the "institution_invite_links" table.
	InstitutionInviteLinksTable = &schema.Table{
		Name:       "institution_invite_links",
		Columns:    InstitutionInviteLinksColumns,
		PrimaryKey: []*schema.Column{InstitutionInviteLinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "institution_invite_links_institutions_invites",
				Columns:    []*schema.Column{InstitutionInviteLinksColumns[3]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// JwtRevocationsColumns holds the columns for the "jwt_revocations" table.
	JwtRevocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "jti", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
	}
	// JwtRevocationsTable holds the schema information for the "jwt_revocations" table.
	JwtRevocationsTable = &schema.Table{
		Name:       "jwt_revocations",
		Columns:    JwtRevocationsColumns,
		PrimaryKey: []*schema.Column{JwtRevocationsColumns[0]},
	}
	// MilestonesColumns holds the columns for the "milestones" table.
	MilestonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "target_completion_time", Type: field.TypeTime},
		{Name: "study_plan_milestones", Type: field.TypeInt},
	}
	// MilestonesTable holds the schema information for the "milestones" table.
	MilestonesTable = &schema.Table{
		Name:       "milestones",
		Columns:    MilestonesColumns,
		PrimaryKey: []*schema.Column{MilestonesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "milestones_study_plans_milestones",
				Columns:    []*schema.Column{MilestonesColumns[3]},
				RefColumns: []*schema.Column{StudyPlansColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "svg_raw", Type: field.TypeString},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
	}
	// ReactionsColumns holds the columns for the "reactions" table.
	ReactionsColumns = []*schema.Column{
		{Name: "emoji", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "forum_post_id", Type: field.TypeInt},
	}
	// ReactionsTable holds the schema information for the "reactions" table.
	ReactionsTable = &schema.Table{
		Name:       "reactions",
		Columns:    ReactionsColumns,
		PrimaryKey: []*schema.Column{ReactionsColumns[2], ReactionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reactions_users_user",
				Columns:    []*schema.Column{ReactionsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "reactions_forum_posts_forum_post",
				Columns:    []*schema.Column{ReactionsColumns[2]},
				RefColumns: []*schema.Column{ForumPostsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RedemptionsColumns holds the columns for the "redemptions" table.
	RedemptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "redeemed_at", Type: field.TypeTime},
		{Name: "accessory_redemptions", Type: field.TypeInt, Nullable: true},
		{Name: "redemption_user", Type: field.TypeInt},
		{Name: "voucher_redemptions", Type: field.TypeInt, Nullable: true},
	}
	// RedemptionsTable holds the schema information for the "redemptions" table.
	RedemptionsTable = &schema.Table{
		Name:       "redemptions",
		Columns:    RedemptionsColumns,
		PrimaryKey: []*schema.Column{RedemptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "redemptions_accessories_redemptions",
				Columns:    []*schema.Column{RedemptionsColumns[2]},
				RefColumns: []*schema.Column{AccessoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "redemptions_users_user",
				Columns:    []*schema.Column{RedemptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "redemptions_vouchers_redemptions",
				Columns:    []*schema.Column{RedemptionsColumns[4]},
				RefColumns: []*schema.Column{VouchersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudyPlansColumns holds the columns for the "study_plans" table.
	StudyPlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString, Nullable: true},
		{Name: "study_plan_author", Type: field.TypeInt},
	}
	// StudyPlansTable holds the schema information for the "study_plans" table.
	StudyPlansTable = &schema.Table{
		Name:       "study_plans",
		Columns:    StudyPlansColumns,
		PrimaryKey: []*schema.Column{StudyPlansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "study_plans_users_author",
				Columns:    []*schema.Column{StudyPlansColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeJSON},
		{Name: "points", Type: field.TypeInt, Default: 0},
		{Name: "points_awarded_count", Type: field.TypeInt, Default: 0},
		{Name: "points_awarded_reset_time", Type: field.TypeTime, Nullable: true},
		{Name: "god_mode", Type: field.TypeBool},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "educator", "member"}},
		{Name: "institution_users", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_institutions_users",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPetsColumns holds the columns for the "user_pets" table.
	UserPetsColumns = []*schema.Column{
		{Name: "hunger_percentage", Type: field.TypeFloat64},
		{Name: "enabled_svg_group_element_ids", Type: field.TypeJSON},
		{Name: "pet_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// UserPetsTable holds the schema information for the "user_pets" table.
	UserPetsTable = &schema.Table{
		Name:       "user_pets",
		Columns:    UserPetsColumns,
		PrimaryKey: []*schema.Column{UserPetsColumns[2], UserPetsColumns[3]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_pets_pets_pet",
				Columns:    []*schema.Column{UserPetsColumns[2]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_pets_users_user",
				Columns:    []*schema.Column{UserPetsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userpet_user_id",
				Unique:  true,
				Columns: []*schema.Column{UserPetsColumns[3]},
			},
		},
	}
	// VouchersColumns holds the columns for the "vouchers" table.
	VouchersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "points_required", Type: field.TypeInt},
		{Name: "institution_vouchers", Type: field.TypeInt},
	}
	// VouchersTable holds the schema information for the "vouchers" table.
	VouchersTable = &schema.Table{
		Name:       "vouchers",
		Columns:    VouchersColumns,
		PrimaryKey: []*schema.Column{VouchersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vouchers_institutions_vouchers",
				Columns:    []*schema.Column{VouchersColumns[4]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DeadlineVotedUsersColumns holds the columns for the "deadline_voted_users" table.
	DeadlineVotedUsersColumns = []*schema.Column{
		{Name: "deadline_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// DeadlineVotedUsersTable holds the schema information for the "deadline_voted_users" table.
	DeadlineVotedUsersTable = &schema.Table{
		Name:       "deadline_voted_users",
		Columns:    DeadlineVotedUsersColumns,
		PrimaryKey: []*schema.Column{DeadlineVotedUsersColumns[0], DeadlineVotedUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deadline_voted_users_deadline_id",
				Columns:    []*schema.Column{DeadlineVotedUsersColumns[0]},
				RefColumns: []*schema.Column{DeadlinesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "deadline_voted_users_user_id",
				Columns:    []*schema.Column{DeadlineVotedUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessoriesTable,
		DeadlinesTable,
		EventsTable,
		ForumsTable,
		ForumPostsTable,
		GroupsTable,
		GroupInviteLinksTable,
		GroupUsersTable,
		InstitutionsTable,
		InstitutionInviteLinksTable,
		JwtRevocationsTable,
		MilestonesTable,
		PetsTable,
		ReactionsTable,
		RedemptionsTable,
		StudyPlansTable,
		UsersTable,
		UserPetsTable,
		VouchersTable,
		DeadlineVotedUsersTable,
	}
)

func init() {
	AccessoriesTable.ForeignKeys[0].RefTable = InstitutionsTable
	DeadlinesTable.ForeignKeys[0].RefTable = UsersTable
	DeadlinesTable.ForeignKeys[1].RefTable = GroupsTable
	EventsTable.ForeignKeys[0].RefTable = GroupsTable
	ForumsTable.ForeignKeys[0].RefTable = GroupsTable
	ForumPostsTable.ForeignKeys[0].RefTable = ForumsTable
	ForumPostsTable.ForeignKeys[1].RefTable = UsersTable
	ForumPostsTable.ForeignKeys[2].RefTable = ForumPostsTable
	GroupsTable.ForeignKeys[0].RefTable = InstitutionsTable
	GroupInviteLinksTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	InstitutionInviteLinksTable.ForeignKeys[0].RefTable = InstitutionsTable
	MilestonesTable.ForeignKeys[0].RefTable = StudyPlansTable
	ReactionsTable.ForeignKeys[0].RefTable = UsersTable
	ReactionsTable.ForeignKeys[1].RefTable = ForumPostsTable
	RedemptionsTable.ForeignKeys[0].RefTable = AccessoriesTable
	RedemptionsTable.ForeignKeys[1].RefTable = UsersTable
	RedemptionsTable.ForeignKeys[2].RefTable = VouchersTable
	StudyPlansTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = InstitutionsTable
	UserPetsTable.ForeignKeys[0].RefTable = PetsTable
	UserPetsTable.ForeignKeys[1].RefTable = UsersTable
	VouchersTable.ForeignKeys[0].RefTable = InstitutionsTable
	DeadlineVotedUsersTable.ForeignKeys[0].RefTable = DeadlinesTable
	DeadlineVotedUsersTable.ForeignKeys[1].RefTable = UsersTable
}
