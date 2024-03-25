package sqlcode
/*
	This file holds file paths for all sql queries that are read in.
	This way, if the paths change or need to be reused, everything is in once place and only needs to be changed once.
*/
const (
	sqlPath = "/Users/josephcocozza/Repositories/unnamed-app/cassidy-wails/internal"
	//Activity

	Activity_create = sqlPath + "/sqlcode/activity/create.sql"
	Activity_update = sqlPath + "/sqlcode/activity/update.sql"
	Activity_delete = sqlPath + "/sqlcode/activity/delete.sql"

	// Planned
	Planned_create     = sqlPath + "/sqlcode/planned/create.sql"
	Planned_read       = sqlPath + "/sqlcode/planned/read.sql"
	Planned_update     = sqlPath + "/sqlcode/planned/update.sql"
	Planned_delete     = sqlPath + "/sqlcode/planned/delete.sql"
	// Completed
	Completed_create   = sqlPath + "/sqlcode/completed/create.sql"
	Completed_read     = sqlPath + "/sqlcode/completed/read.sql"
	Completed_update   = sqlPath + "/sqlcode/completed/update.sql"
	Completed_delete   = sqlPath + "/sqlcode/completed/delete.sql"
	// User
	User_create        = sqlPath + "/sqlcode/user/create.sql"
	User_read          = sqlPath + "/sqlcode/user/read.sql"
	User_update        = sqlPath + "/sqlcode/user/update.sql"
	User_delete        = sqlPath + "/sqlcode/user/delete.sql"
	User_preferences   = sqlPath + "/sqlcode/user/read_preferences.sql"
	// Equipment Type
	EquipmentType_list = sqlPath + "/sqlcode/equipment_type/list.sql"
	// Activity Equipment
	ActivityEquipment_create = sqlPath + "/sqlcode/activity_equipment/create.sql"
	ActivityEquipment_update = sqlPath + "/sqlcode/activity_equipment/update.sql"
	ActivityEquipment_delete = sqlPath + "/sqlcode/activity_equipment/delete.sql"
	ActivityEquipment_list = sqlPath + "/sqlcode/activity_equipment/list_by_activity_uuid.sql"

	// Activity type
	ActivityType_list_with_subtype = sqlPath + "/sqlcode/activity_type/list_type_with_subtype.sql"

	// Activity Type Subtype
	ActivityTypeSubtype_list = sqlPath + "/sqlcode/activity_type_subtype/list_by_activity_uuid.sql"
	ActivityTypeSubtype_create = sqlPath + "/sqlcode/activity_type_subtype/create.sql"
	ActivityTypeSubtype_delete_all_by_activity_uuid = sqlPath + "/sqlcode/activity_type_subtype/delete_all_by_activity_uuid.sql"

	// Equipment
	Equipment_create = sqlPath + "/sqlcode/equipment/create.sql"
	Equipment_update = sqlPath + "/sqlcode/equipment/update.sql"
	Equipment_delete = sqlPath + "/sqlcode/equipment/delete.sql"
	Equipment_list = sqlPath + "/sqlcode/equipment/list.sql"
	// Microcycle
	Microcycle_read_activity_list = sqlPath + "/sqlcode/microcycle/read_activity_list.sql"
	Microcycle_read_totals = sqlPath + "/sqlcode/microcycle/read_totals.sql"
	Microcycle_read_totals_current_previous = sqlPath + "/sqlcode/microcycle/read_totals_current_previous.sql"
	Microcycle_read_totals_by_activity_type = sqlPath + "/sqlcode/microcycle/read_totals_by_activity_type.sql"
	Microcycle_read_totals_by_activity_type_current_previous = sqlPath + "/sqlcode/microcycle/read_totals_by_activity_type_current_previous.sql"
	Microcycle_read_totals_by_activity_type_and_date = sqlPath + "/sqlcode/microcycle/read_totals_by_activity_type_and_date.sql"

	Microcycle_read_totals_date_range = sqlPath + "/sqlcode/microcycle/read_totals_date_range.sql"
	// Misc

	N_cycle_summary = sqlPath + "/sqlcode/misc/n_cycle_summary.sql"
)