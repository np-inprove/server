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
				OnDelete:   schema.NoAction,
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
		{Name: "god_mode", Type: field.TypeBool},
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
		AccessoriesTable,
		CoursesTable,
		InstitutionsTable,
		PetsTable,
		RedemptionsTable,
		UsersTable,
		UserPetsTable,
		VouchersTable,
		InstitutionAdminsTable,
	}
)

func init() {
	AcademicSchoolsTable.ForeignKeys[0].RefTable = InstitutionsTable
	AccessoriesTable.ForeignKeys[0].RefTable = InstitutionsTable
	CoursesTable.ForeignKeys[0].RefTable = AcademicSchoolsTable
	RedemptionsTable.ForeignKeys[0].RefTable = AccessoriesTable
	RedemptionsTable.ForeignKeys[1].RefTable = UsersTable
	RedemptionsTable.ForeignKeys[2].RefTable = VouchersTable
	UsersTable.ForeignKeys[0].RefTable = CoursesTable
	UserPetsTable.ForeignKeys[0].RefTable = PetsTable
	UserPetsTable.ForeignKeys[1].RefTable = UsersTable
	VouchersTable.ForeignKeys[0].RefTable = InstitutionsTable
	InstitutionAdminsTable.ForeignKeys[0].RefTable = InstitutionsTable
	InstitutionAdminsTable.ForeignKeys[1].RefTable = UsersTable
}
