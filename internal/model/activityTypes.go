package model

var (
	Run = ActivityType{Id: 1, Name: "Run"}
		Run_Long = ActivitySubtype{SuperTypeId: Run.Id, Name: "Long"}
		Run_Fartlek = ActivitySubtype{SuperTypeId: Run.Id, Name: "Fartlek"}
		Run_Tempo = ActivitySubtype{SuperTypeId: Run.Id, Name: "Tempo"}
		Run_Track = ActivitySubtype{SuperTypeId: Run.Id, Name: "Track"}
		Run_Intervals = ActivitySubtype{SuperTypeId: Run.Id, Name: "Intervals"}
		Run_Recovery = ActivitySubtype{SuperTypeId: Run.Id, Name: "Recovery"}
		Run_Indoor = ActivitySubtype{SuperTypeId: Run.Id, Name: "Indoor"}
		Run_Trails = ActivitySubtype{SuperTypeId: Run.Id, Name: "Trails"}
	Bike = ActivityType{Id: 2, Name: "Bike"}
		Bike_Long = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Long"}
		Bike_Velodrome = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Velodrome"}
		Bike_Recovery = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Recovery"}
		Bike_Indoor = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Indoor"}
	Swim = ActivityType{Id: 3, Name: "Swim"}
		Swim_Drills = ActivitySubtype{SuperTypeId: Swim.Id, Name: "Drills"}
		Swim_OpenWater = ActivitySubtype{SuperTypeId: Swim.Id, Name: "Open Water"}
		Swim_Recovery = ActivitySubtype{SuperTypeId: Swim.Id, Name: "Recovery"}
	Hike = ActivityType{Id: 4, Name: "Hike"}
	RestDay = ActivityType{Id: 5, Name: "Rest Day"}
	Strength = ActivityType{Id: 6, Name: "Strength"}
	Other = ActivityType{Id: 7, Name: "Other"}
	MountainBike = ActivityType{Id: 8, Name: "Mountain Bike"}
)