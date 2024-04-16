package sqlcode

import (
	"embed"
	//"log/slog"
)

/*
	This file holds file paths for all sql queries that are read in.
	This way, if the paths change or need to be reused, everything is in once place and only needs to be changed once.
*/

//go:embed all:sql
var sql embed.FS

const (
	//sqlPath = "/Users/josephcocozza/Repositories/unnamed-app/cassidy-wails/internal"

	//Activity
	Activity_create = "sql/activity/create.sql"
	Activity_update = "sql/activity/update.sql"
	Activity_delete = "sql/activity/delete.sql"

	// Planned
	Planned_create     = "sql/planned/create.sql"
	Planned_read       = "sql/planned/read.sql"
	Planned_update     = "sql/planned/update.sql"
	Planned_delete     = "sql/planned/delete.sql"
	// Completed
	Completed_create   = "sql/completed/create.sql"
	Completed_read     = "sql/completed/read.sql"
	Completed_update   = "sql/completed/update.sql"
	Completed_delete   = "sql/completed/delete.sql"
	// User
	User_create        = "sql/user/create.sql"
	User_read          = "sql/user/read.sql"
	User_update        = "sql/user/update.sql"
	User_delete        = "sql/user/delete.sql"
	User_preferences   = "sql/user/read_preferences.sql"
	// User strava token
	Strava_token_create = "sql/user/create_strava_token.sql"
	Strava_token_update = "sql/user/update_strava_token.sql"
	Strava_token_read   = "sql/user/read_strava_token.sql"
	// Equipment Type
	EquipmentType_list = "sql/equipment_type/list.sql"
	// Activity Equipment
	ActivityEquipment_create = "sql/activity_equipment/create.sql"
	ActivityEquipment_update = "sql/activity_equipment/update.sql"
	ActivityEquipment_delete = "sql/activity_equipment/delete.sql"
	ActivityEquipment_list = "sql/activity_equipment/list_by_activity_uuid.sql"

	// Activity type
	ActivityType_list_with_subtype = "sql/activity_type/list_type_with_subtype.sql"

	// Activity Type Subtype
	ActivityTypeSubtype_list = "sql/activity_type_subtype/list_by_activity_uuid.sql"
	ActivityTypeSubtype_create = "sql/activity_type_subtype/create.sql"
	ActivityTypeSubtype_delete_all_by_activity_uuid = "sql/activity_type_subtype/delete_all_by_activity_uuid.sql"

	// Equipment
	Equipment_create = "sql/equipment/create.sql"
	Equipment_update = "sql/equipment/update.sql"
	Equipment_delete = "sql/equipment/delete.sql"
	Equipment_list = "sql/equipment/list.sql"
	// Microcycle
	Microcycle_read_activity_list = "sql/microcycle/read_activity_list.sql"
	Microcycle_read_totals = "sql/microcycle/read_totals.sql"
	Microcycle_read_totals_current_previous = "sql/microcycle/read_totals_current_previous.sql"
	Microcycle_read_totals_by_activity_type = "sql/microcycle/read_totals_by_activity_type.sql"
	Microcycle_read_totals_by_activity_type_current_previous = "sql/microcycle/read_totals_by_activity_type_current_previous.sql"
	Microcycle_read_totals_by_activity_type_and_date = "sql/microcycle/read_totals_by_activity_type_and_date.sql"

	Microcycle_read_totals_date_range = "sql/microcycle/read_totals_date_range.sql"
	// Misc
	N_cycle_summary = "sql/misc/n_cycle_summary.sql"
)
// Read in a sql file from the embeded file directory
func SQLReader(filePath string) string {
	//slog.Info("Reading sql from: " + filePath)
	sqlContent, err := sql.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	//slog.Info("SQL:" + string(sqlContent))
	return string(sqlContent)
}