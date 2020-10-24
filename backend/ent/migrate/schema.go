// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ProblemsColumns holds the columns for the "problems" table.
	ProblemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "probleminfo", Type: field.TypeString},
		{Name: "adddate", Type: field.TypeTime},
		{Name: "problemstatus", Type: field.TypeInt, Nullable: true},
		{Name: "problemtitle", Type: field.TypeInt, Nullable: true},
		{Name: "room_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// ProblemsTable holds the schema information for the "problems" table.
	ProblemsTable = &schema.Table{
		Name:       "problems",
		Columns:    ProblemsColumns,
		PrimaryKey: []*schema.Column{ProblemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "problems_problem_status_problemstatus2problem",
				Columns: []*schema.Column{ProblemsColumns[3]},

				RefColumns: []*schema.Column{ProblemStatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "problems_problem_titles_problemtitle2problem",
				Columns: []*schema.Column{ProblemsColumns[4]},

				RefColumns: []*schema.Column{ProblemTitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "problems_rooms_room2problem",
				Columns: []*schema.Column{ProblemsColumns[5]},

				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "problems_users_user2problem",
				Columns: []*schema.Column{ProblemsColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProblemStatusColumns holds the columns for the "problem_status" table.
	ProblemStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "problemstatus", Type: field.TypeString},
	}
	// ProblemStatusTable holds the schema information for the "problem_status" table.
	ProblemStatusTable = &schema.Table{
		Name:        "problem_status",
		Columns:     ProblemStatusColumns,
		PrimaryKey:  []*schema.Column{ProblemStatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProblemTitlesColumns holds the columns for the "problem_titles" table.
	ProblemTitlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "problemtitle", Type: field.TypeString},
	}
	// ProblemTitlesTable holds the schema information for the "problem_titles" table.
	ProblemTitlesTable = &schema.Table{
		Name:        "problem_titles",
		Columns:     ProblemTitlesColumns,
		PrimaryKey:  []*schema.Column{ProblemTitlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "rooms_users_user2room",
				Columns: []*schema.Column{RoomsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProblemsTable,
		ProblemStatusTable,
		ProblemTitlesTable,
		RoomsTable,
		UsersTable,
	}
)

func init() {
	ProblemsTable.ForeignKeys[0].RefTable = ProblemStatusTable
	ProblemsTable.ForeignKeys[1].RefTable = ProblemTitlesTable
	ProblemsTable.ForeignKeys[2].RefTable = RoomsTable
	ProblemsTable.ForeignKeys[3].RefTable = UsersTable
	RoomsTable.ForeignKeys[0].RefTable = UsersTable
}
