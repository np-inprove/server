// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AcademicSchoolsColumns holds the columns for the "academic_schools" table.
	AcademicSchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "school_code", Type: field.TypeString},
		{Name: "institution_academic_schools", Type: field.TypeInt, Nullable: true},
	}
	// AcademicSchoolsTable holds the schema information for the "academic_schools" table.
	AcademicSchoolsTable = &schema.Table{
		Name:       "academic_schools",
		Columns:    AcademicSchoolsColumns,
		PrimaryKey: []*schema.Column{AcademicSchoolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "academic_schools_institutions_academic_schools",
				Columns:    []*schema.Column{AcademicSchoolsColumns[3]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "course_id", Type: field.TypeString},
		{Name: "academic_school_courses", Type: field.TypeInt, Nullable: true},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "courses_academic_schools_courses",
				Columns:    []*schema.Column{CoursesColumns[3]},
				RefColumns: []*schema.Column{AcademicSchoolsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InstitutionsColumns holds the columns for the "institutions" table.
	InstitutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// InstitutionsTable holds the schema information for the "institutions" table.
	InstitutionsTable = &schema.Table{
		Name:       "institutions",
		Columns:    InstitutionsColumns,
		PrimaryKey: []*schema.Column{InstitutionsColumns[0]},
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
	// PrizesColumns holds the columns for the "prizes" table.
	PrizesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "points_required", Type: field.TypeInt},
		{Name: "discriminator", Type: field.TypeEnum, Enums: []string{"voucher", "accessory"}},
		{Name: "institution_prizes", Type: field.TypeInt, Nullable: true},
	}
	// PrizesTable holds the schema information for the "prizes" table.
	PrizesTable = &schema.Table{
		Name:       "prizes",
		Columns:    PrizesColumns,
		PrimaryKey: []*schema.Column{PrizesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prizes_institutions_prizes",
				Columns:    []*schema.Column{PrizesColumns[5]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PrizeRedemptionsColumns holds the columns for the "prize_redemptions" table.
	PrizeRedemptionsColumns = []*schema.Column{
		{Name: "redeemed_at", Type: field.TypeTime},
		{Name: "prize_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// PrizeRedemptionsTable holds the schema information for the "prize_redemptions" table.
	PrizeRedemptionsTable = &schema.Table{
		Name:       "prize_redemptions",
		Columns:    PrizeRedemptionsColumns,
		PrimaryKey: []*schema.Column{PrizeRedemptionsColumns[1], PrizeRedemptionsColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prize_redemptions_prizes_prize",
				Columns:    []*schema.Column{PrizeRedemptionsColumns[1]},
				RefColumns: []*schema.Column{PrizesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "prize_redemptions_users_user",
				Columns:    []*schema.Column{PrizeRedemptionsColumns[2]},
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
		{Name: "email", Type: field.TypeString},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "points", Type: field.TypeInt},
		{Name: "points_awarded_count", Type: field.TypeInt},
		{Name: "points_awarded_reset_time", Type: field.TypeTime, Nullable: true},
		{Name: "superuser", Type: field.TypeBool},
		{Name: "course_students", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_courses_students",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.SetNull,
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
	}
	// InstitutionAdminsColumns holds the columns for the "institution_admins" table.
	InstitutionAdminsColumns = []*schema.Column{
		{Name: "institution_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// InstitutionAdminsTable holds the schema information for the "institution_admins" table.
	InstitutionAdminsTable = &schema.Table{
		Name:       "institution_admins",
		Columns:    InstitutionAdminsColumns,
		PrimaryKey: []*schema.Column{InstitutionAdminsColumns[0], InstitutionAdminsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "institution_admins_institution_id",
				Columns:    []*schema.Column{InstitutionAdminsColumns[0]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "institution_admins_user_id",
				Columns:    []*schema.Column{InstitutionAdminsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AcademicSchoolsTable,
		CoursesTable,
		InstitutionsTable,
		PetsTable,
		PrizesTable,
		PrizeRedemptionsTable,
		UsersTable,
		UserPetsTable,
		InstitutionAdminsTable,
	}
)

func init() {
	AcademicSchoolsTable.ForeignKeys[0].RefTable = InstitutionsTable
	CoursesTable.ForeignKeys[0].RefTable = AcademicSchoolsTable
	PrizesTable.ForeignKeys[0].RefTable = InstitutionsTable
	PrizeRedemptionsTable.ForeignKeys[0].RefTable = PrizesTable
	PrizeRedemptionsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = CoursesTable
	UserPetsTable.ForeignKeys[0].RefTable = PetsTable
	UserPetsTable.ForeignKeys[1].RefTable = UsersTable
	InstitutionAdminsTable.ForeignKeys[0].RefTable = InstitutionsTable
	InstitutionAdminsTable.ForeignKeys[1].RefTable = UsersTable
}
