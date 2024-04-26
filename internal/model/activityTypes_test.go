package model

import "testing"

func TestActivityType_Validate(t *testing.T) {
	type fields struct {
		Id   int
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, Name: "Run"}, false},
		{"invalid id", fields{Id: -1, Name: "Run"}, true},
		{"invalid name", fields{Id: 1, Name: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			at := &ActivityType{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
			}
			if err := at.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ActivityType.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivitySubtype_Validate(t *testing.T) {
	type fields struct {
		Id          int
		SuperTypeId int
		Name        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, SuperTypeId: 1, Name: "Long"}, false},
		{"invalid id", fields{Id: -1, SuperTypeId: 1, Name: "Long"}, true},
		{"invalid supertype id", fields{Id: 1, SuperTypeId: -1, Name: "Long"}, true},
		{"invalid name", fields{Id: 1, SuperTypeId: -1, Name: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := &ActivitySubtype{
				Id:          tt.fields.Id,
				SuperTypeId: tt.fields.SuperTypeId,
				Name:        tt.fields.Name,
			}
			if err := as.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ActivitySubtype.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivityTypeSubtype_Validate(t *testing.T) {
	type fields struct {
		Id              int
		ActivityUuid    string
		ActivityType    *ActivityType
		ActivitySubtype *ActivitySubtype
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, ActivityUuid: "c63fd1dc-cb8a-11ee-90d8-325096b39f47", ActivityType: &ActivityType{Id: 1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, false},
		{"invalid activity id", fields{Id: -1, ActivityUuid: "c63fd1dc-cb8a-11ee-90d8-325096b39f47", ActivityType: &ActivityType{Id: 1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, true},
		{"invalid activity uuid", fields{Id: 1, ActivityUuid: "", ActivityType: &ActivityType{Id: 1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, true},
		{"invalid activity type", fields{Id: 1, ActivityUuid: "", ActivityType: &ActivityType{Id: -1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, true},
		{"invalid activity subtype", fields{Id: 1, ActivityUuid: "", ActivityType: &ActivityType{Id: -1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: -1, SuperTypeId: 1, Name: "Long"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ats := &ActivityTypeSubtype{
				Id:              tt.fields.Id,
				ActivityUuid:    tt.fields.ActivityUuid,
				ActivityType:    tt.fields.ActivityType,
				ActivitySubtype: tt.fields.ActivitySubtype,
			}
			if err := ats.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ActivityTypeSubtype.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}