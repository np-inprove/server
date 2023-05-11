// Code generated by ent, DO NOT EDIT.

package redemption

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the redemption type in the database.
	Label = "redemption"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRedeemedAt holds the string denoting the redeemed_at field in the database.
	FieldRedeemedAt = "redeemed_at"
	// EdgeVoucher holds the string denoting the voucher edge name in mutations.
	EdgeVoucher = "voucher"
	// EdgeAccessory holds the string denoting the accessory edge name in mutations.
	EdgeAccessory = "accessory"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the redemption in the database.
	Table = "redemptions"
	// VoucherTable is the table that holds the voucher relation/edge.
	VoucherTable = "redemptions"
	// VoucherInverseTable is the table name for the Voucher entity.
	// It exists in this package in order to avoid circular dependency with the "voucher" package.
	VoucherInverseTable = "vouchers"
	// VoucherColumn is the table column denoting the voucher relation/edge.
	VoucherColumn = "voucher_redemptions"
	// AccessoryTable is the table that holds the accessory relation/edge.
	AccessoryTable = "redemptions"
	// AccessoryInverseTable is the table name for the Accessory entity.
	// It exists in this package in order to avoid circular dependency with the "accessory" package.
	AccessoryInverseTable = "accessories"
	// AccessoryColumn is the table column denoting the accessory relation/edge.
	AccessoryColumn = "accessory_redemptions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "redemptions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "redemption_user"
)

// Columns holds all SQL columns for redemption fields.
var Columns = []string{
	FieldID,
	FieldRedeemedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "redemptions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"accessory_redemptions",
	"redemption_user",
	"voucher_redemptions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Redemption queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRedeemedAt orders the results by the redeemed_at field.
func ByRedeemedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedeemedAt, opts...).ToFunc()
}

// ByVoucherField orders the results by voucher field.
func ByVoucherField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVoucherStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccessoryField orders the results by accessory field.
func ByAccessoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccessoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newVoucherStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VoucherInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VoucherTable, VoucherColumn),
	)
}
func newAccessoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccessoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccessoryTable, AccessoryColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
