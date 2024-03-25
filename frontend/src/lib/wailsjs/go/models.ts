export namespace controllers {
	
	export class NextPrevious {
	    next_start_date: string;
	    next_end_date: string;
	    previous_start_date: string;
	    previous_end_date: string;
	
	    static createFrom(source: any = {}) {
	        return new NextPrevious(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.next_start_date = source["next_start_date"];
	        this.next_end_date = source["next_end_date"];
	        this.previous_start_date = source["previous_start_date"];
	        this.previous_end_date = source["previous_end_date"];
	    }
	}
	export class authRequest {
	    username: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new authRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}

}

export namespace measurement {
	
	export class Measurement {
	    unit: string;
	    length: number;
	
	    static createFrom(source: any = {}) {
	        return new Measurement(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.unit = source["unit"];
	        this.length = source["length"];
	    }
	}

}

export namespace model {
	
	export class Completed {
	    activity_uuid: string;
	    distance?: measurement.Measurement;
	    duration: number;
	    vertical?: measurement.Measurement;
	    pace: string;
	
	    static createFrom(source: any = {}) {
	        return new Completed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activity_uuid = source["activity_uuid"];
	        this.distance = this.convertValues(source["distance"], measurement.Measurement);
	        this.duration = source["duration"];
	        this.vertical = this.convertValues(source["vertical"], measurement.Measurement);
	        this.pace = source["pace"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Planned {
	    activity_uuid: string;
	    distance?: measurement.Measurement;
	    duration: number;
	    vertical?: measurement.Measurement;
	    pace: string;
	
	    static createFrom(source: any = {}) {
	        return new Planned(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activity_uuid = source["activity_uuid"];
	        this.distance = this.convertValues(source["distance"], measurement.Measurement);
	        this.duration = source["duration"];
	        this.vertical = this.convertValues(source["vertical"], measurement.Measurement);
	        this.pace = source["pace"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class EquipmentType {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new EquipmentType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class Equipment {
	    id: number;
	    user_uuid: string;
	    equipment_type?: EquipmentType;
	    name: string;
	    brand: string;
	    model: string;
	    cost: number;
	    size: string;
	    purchase_date: string;
	    notes: string;
	    mileage?: measurement.Measurement;
	    is_retired: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Equipment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.user_uuid = source["user_uuid"];
	        this.equipment_type = this.convertValues(source["equipment_type"], EquipmentType);
	        this.name = source["name"];
	        this.brand = source["brand"];
	        this.model = source["model"];
	        this.cost = source["cost"];
	        this.size = source["size"];
	        this.purchase_date = source["purchase_date"];
	        this.notes = source["notes"];
	        this.mileage = this.convertValues(source["mileage"], measurement.Measurement);
	        this.is_retired = source["is_retired"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ActivityEquipment {
	    id: number;
	    activity_uuid: string;
	    equipment?: Equipment;
	    assigned_mileage?: measurement.Measurement;
	
	    static createFrom(source: any = {}) {
	        return new ActivityEquipment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.activity_uuid = source["activity_uuid"];
	        this.equipment = this.convertValues(source["equipment"], Equipment);
	        this.assigned_mileage = this.convertValues(source["assigned_mileage"], measurement.Measurement);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ActivitySubtype {
	    id: number;
	    supertype_id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new ActivitySubtype(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.supertype_id = source["supertype_id"];
	        this.name = source["name"];
	    }
	}
	export class ActivityTypeSubtype {
	    id: number;
	    activity_uuid: string;
	    activity_type?: ActivityType;
	    // Go type: ActivitySubtype
	    activity_subtype?: any;
	
	    static createFrom(source: any = {}) {
	        return new ActivityTypeSubtype(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.activity_uuid = source["activity_uuid"];
	        this.activity_type = this.convertValues(source["activity_type"], ActivityType);
	        this.activity_subtype = this.convertValues(source["activity_subtype"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ActivityType {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new ActivityType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class Activity {
	    uuid: string;
	    date: string;
	    order: number;
	    name: string;
	    description: string;
	    notes: string;
	    activity_type?: ActivityType;
	    type_subtype_list: ActivityTypeSubtype[];
	    equipment_list: ActivityEquipment[];
	    planned?: Planned;
	    completed?: Completed;
	    color: string;
	    is_race: boolean;
	    num_strides: number;
	
	    static createFrom(source: any = {}) {
	        return new Activity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uuid = source["uuid"];
	        this.date = source["date"];
	        this.order = source["order"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.notes = source["notes"];
	        this.activity_type = this.convertValues(source["activity_type"], ActivityType);
	        this.type_subtype_list = this.convertValues(source["type_subtype_list"], ActivityTypeSubtype);
	        this.equipment_list = this.convertValues(source["equipment_list"], ActivityEquipment);
	        this.planned = this.convertValues(source["planned"], Planned);
	        this.completed = this.convertValues(source["completed"], Completed);
	        this.color = source["color"];
	        this.is_race = source["is_race"];
	        this.num_strides = source["num_strides"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class ActivityTypeWithSubtypes {
	    activity_type?: ActivityType;
	    subtype_list: ActivitySubtype[];
	
	    static createFrom(source: any = {}) {
	        return new ActivityTypeWithSubtypes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activity_type = this.convertValues(source["activity_type"], ActivityType);
	        this.subtype_list = this.convertValues(source["subtype_list"], ActivitySubtype);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	export class TotalByActivityTypeAndDate {
	    activity_type?: ActivityType;
	    // Go type: dateutil
	    date?: any;
	    total_planned_distance?: measurement.Measurement;
	    total_planned_duration: number;
	    total_planned_vertical?: measurement.Measurement;
	    total_completed_distance?: measurement.Measurement;
	    total_completed_duration: number;
	    total_completed_vertical?: measurement.Measurement;
	
	    static createFrom(source: any = {}) {
	        return new TotalByActivityTypeAndDate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activity_type = this.convertValues(source["activity_type"], ActivityType);
	        this.date = this.convertValues(source["date"], null);
	        this.total_planned_distance = this.convertValues(source["total_planned_distance"], measurement.Measurement);
	        this.total_planned_duration = source["total_planned_duration"];
	        this.total_planned_vertical = this.convertValues(source["total_planned_vertical"], measurement.Measurement);
	        this.total_completed_distance = this.convertValues(source["total_completed_distance"], measurement.Measurement);
	        this.total_completed_duration = source["total_completed_duration"];
	        this.total_completed_vertical = this.convertValues(source["total_completed_vertical"], measurement.Measurement);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TotalByActivityTypeDifferences {
	    activity_type?: ActivityType;
	    difference_planned_distance?: measurement.Measurement;
	    percent_difference_planned_distance: number;
	    difference_planned_duration: number;
	    percent_difference_planned_duration: number;
	    difference_planned_vertical?: measurement.Measurement;
	    percent_difference_planned_vertical: number;
	    difference_completed_distance?: measurement.Measurement;
	    percent_difference_completed_distance: number;
	    difference_completed_duration: number;
	    percent_difference_completed_duration: number;
	    difference_completed_vertical?: measurement.Measurement;
	    percent_difference_completed_vertical: number;
	
	    static createFrom(source: any = {}) {
	        return new TotalByActivityTypeDifferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activity_type = this.convertValues(source["activity_type"], ActivityType);
	        this.difference_planned_distance = this.convertValues(source["difference_planned_distance"], measurement.Measurement);
	        this.percent_difference_planned_distance = source["percent_difference_planned_distance"];
	        this.difference_planned_duration = source["difference_planned_duration"];
	        this.percent_difference_planned_duration = source["percent_difference_planned_duration"];
	        this.difference_planned_vertical = this.convertValues(source["difference_planned_vertical"], measurement.Measurement);
	        this.percent_difference_planned_vertical = source["percent_difference_planned_vertical"];
	        this.difference_completed_distance = this.convertValues(source["difference_completed_distance"], measurement.Measurement);
	        this.percent_difference_completed_distance = source["percent_difference_completed_distance"];
	        this.difference_completed_duration = source["difference_completed_duration"];
	        this.percent_difference_completed_duration = source["percent_difference_completed_duration"];
	        this.difference_completed_vertical = this.convertValues(source["difference_completed_vertical"], measurement.Measurement);
	        this.percent_difference_completed_vertical = source["percent_difference_completed_vertical"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TotalByActivityType {
	    activity_type?: ActivityType;
	    total_planned_distance?: measurement.Measurement;
	    total_planned_duration: number;
	    total_planned_vertical?: measurement.Measurement;
	    planned_pace: string;
	    total_completed_distance?: measurement.Measurement;
	    total_completed_duration: number;
	    total_completed_vertical?: measurement.Measurement;
	    completed_pace: string;
	
	    static createFrom(source: any = {}) {
	        return new TotalByActivityType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activity_type = this.convertValues(source["activity_type"], ActivityType);
	        this.total_planned_distance = this.convertValues(source["total_planned_distance"], measurement.Measurement);
	        this.total_planned_duration = source["total_planned_duration"];
	        this.total_planned_vertical = this.convertValues(source["total_planned_vertical"], measurement.Measurement);
	        this.planned_pace = source["planned_pace"];
	        this.total_completed_distance = this.convertValues(source["total_completed_distance"], measurement.Measurement);
	        this.total_completed_duration = source["total_completed_duration"];
	        this.total_completed_vertical = this.convertValues(source["total_completed_vertical"], measurement.Measurement);
	        this.completed_pace = source["completed_pace"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TotalsDifferences {
	    difference_planned_distance?: measurement.Measurement;
	    percent_difference_planned_distance: number;
	    difference_planned_duration: number;
	    percent_difference_planned_duration: number;
	    difference_planned_vertical?: measurement.Measurement;
	    percent_difference_planned_vertical: number;
	    difference_completed_distance?: measurement.Measurement;
	    percent_difference_completed_distance: number;
	    difference_completed_duration: number;
	    percent_difference_completed_duration: number;
	    difference_completed_vertical?: measurement.Measurement;
	    percent_difference_completed_vertical: number;
	
	    static createFrom(source: any = {}) {
	        return new TotalsDifferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.difference_planned_distance = this.convertValues(source["difference_planned_distance"], measurement.Measurement);
	        this.percent_difference_planned_distance = source["percent_difference_planned_distance"];
	        this.difference_planned_duration = source["difference_planned_duration"];
	        this.percent_difference_planned_duration = source["percent_difference_planned_duration"];
	        this.difference_planned_vertical = this.convertValues(source["difference_planned_vertical"], measurement.Measurement);
	        this.percent_difference_planned_vertical = source["percent_difference_planned_vertical"];
	        this.difference_completed_distance = this.convertValues(source["difference_completed_distance"], measurement.Measurement);
	        this.percent_difference_completed_distance = source["percent_difference_completed_distance"];
	        this.difference_completed_duration = source["difference_completed_duration"];
	        this.percent_difference_completed_duration = source["percent_difference_completed_duration"];
	        this.difference_completed_vertical = this.convertValues(source["difference_completed_vertical"], measurement.Measurement);
	        this.percent_difference_completed_vertical = source["percent_difference_completed_vertical"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Totals {
	    total_planned_distance?: measurement.Measurement;
	    total_planned_duration: number;
	    total_planned_vertical?: measurement.Measurement;
	    total_completed_distance?: measurement.Measurement;
	    total_completed_duration: number;
	    total_completed_vertical?: measurement.Measurement;
	
	    static createFrom(source: any = {}) {
	        return new Totals(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total_planned_distance = this.convertValues(source["total_planned_distance"], measurement.Measurement);
	        this.total_planned_duration = source["total_planned_duration"];
	        this.total_planned_vertical = this.convertValues(source["total_planned_vertical"], measurement.Measurement);
	        this.total_completed_distance = this.convertValues(source["total_completed_distance"], measurement.Measurement);
	        this.total_completed_duration = source["total_completed_duration"];
	        this.total_completed_vertical = this.convertValues(source["total_completed_vertical"], measurement.Measurement);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MicrocycleSummary {
	    totals?: Totals;
	    previous_totals?: Totals;
	    average_previous_totals?: Totals;
	    totals_differences?: TotalsDifferences;
	    weighted_totals_differences?: TotalsDifferences;
	    totals_by_activity_type: TotalByActivityType[];
	    totals_by_activity_type_differences: TotalByActivityTypeDifferences[];
	    previous_totals_by_activity_type: TotalByActivityType[];
	    totals_by_activity_type_and_date: TotalByActivityTypeAndDate[];
	
	    static createFrom(source: any = {}) {
	        return new MicrocycleSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.totals = this.convertValues(source["totals"], Totals);
	        this.previous_totals = this.convertValues(source["previous_totals"], Totals);
	        this.average_previous_totals = this.convertValues(source["average_previous_totals"], Totals);
	        this.totals_differences = this.convertValues(source["totals_differences"], TotalsDifferences);
	        this.weighted_totals_differences = this.convertValues(source["weighted_totals_differences"], TotalsDifferences);
	        this.totals_by_activity_type = this.convertValues(source["totals_by_activity_type"], TotalByActivityType);
	        this.totals_by_activity_type_differences = this.convertValues(source["totals_by_activity_type_differences"], TotalByActivityTypeDifferences);
	        this.previous_totals_by_activity_type = this.convertValues(source["previous_totals_by_activity_type"], TotalByActivityType);
	        this.totals_by_activity_type_and_date = this.convertValues(source["totals_by_activity_type_and_date"], TotalByActivityTypeAndDate);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ActivityList {
	    // Go type: dateutil
	    date_object?: any;
	    activity_list: Activity[];
	
	    static createFrom(source: any = {}) {
	        return new ActivityList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date_object = this.convertValues(source["date_object"], null);
	        this.activity_list = this.convertValues(source["activity_list"], Activity);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Microcycle {
	    start_date: string;
	    end_date: string;
	    cycle_activities?: ActivityList[];
	    summary?: MicrocycleSummary;
	
	    static createFrom(source: any = {}) {
	        return new Microcycle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.start_date = source["start_date"];
	        this.end_date = source["end_date"];
	        this.cycle_activities = this.convertValues(source["cycle_activities"], ActivityList);
	        this.summary = this.convertValues(source["summary"], MicrocycleSummary);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class NCycleSummary {
	    start_date_list: dateutil.DateObject[];
	    planned_distances: measurement.Measurement[];
	    planned_durations: number[];
	    planned_verticals: measurement.Measurement[];
	    completed_distances: measurement.Measurement[];
	    completed_durations: number[];
	    completed_verticals: measurement.Measurement[];
	
	    static createFrom(source: any = {}) {
	        return new NCycleSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.start_date_list = this.convertValues(source["start_date_list"], dateutil.DateObject);
	        this.planned_distances = this.convertValues(source["planned_distances"], measurement.Measurement);
	        this.planned_durations = source["planned_durations"];
	        this.planned_verticals = this.convertValues(source["planned_verticals"], measurement.Measurement);
	        this.completed_distances = this.convertValues(source["completed_distances"], measurement.Measurement);
	        this.completed_durations = source["completed_durations"];
	        this.completed_verticals = this.convertValues(source["completed_verticals"], measurement.Measurement);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	export class User {
	    uuid: string;
	    username: string;
	    password: string;
	    units: string;
	    cycle_start: string;
	    cycle_days: number;
	    initial_cycle_start: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uuid = source["uuid"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.units = source["units"];
	        this.cycle_start = source["cycle_start"];
	        this.cycle_days = source["cycle_days"];
	        this.initial_cycle_start = source["initial_cycle_start"];
	    }
	}

}

