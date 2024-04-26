package userrepo
/*
import (
	"reflect"
	"testing"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
)

func TestIUserRepository_Create(t *testing.T) {
	d := database.InitTestDB()
	type fields struct {
		DB database.DbOperations
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"create valid", fields{d}, args{&model.User{Uuid: "07018c78-cc09-11ee-8fd9-325096b39f47", Username: "test1@test.com", Password: "test", Units: "imperial", CycleStart: "Monday", CycleDays: 7}}, true},
		{"create already existing uuid", fields{d}, args{&model.User{Uuid: "d4e0fbfe-c945-11ee-b8e2-325096b39f47", Username: "test1@test.com", Password: "test", Units: "imperial", CycleStart: "Monday", CycleDays: 7}}, true},
		{"create already existing username", fields{d}, args{&model.User{Uuid: "d4e0fbfe-c945-11ee-b8e2-325096b39f48", Username: "test1@test.com", Password: "test", Units: "imperial", CycleStart: "Monday", CycleDays: 7}}, true},
		{"create invliad", fields{d}, args{&model.User{Uuid: "", Username: "test1@test.com", Password: "test", Units: "imperial", CycleStart: "Monday", CycleDays: 7}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &IUserRepository{
				DB: tt.fields.DB,
			}
			if err := db.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("IUserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIUserRepository_Read(t *testing.T) {
	d := database.InitTestDB()
	type fields struct {
		DB database.DbOperations
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"read valid", fields{d}, args{"test1@test.com"}, &model.User{Uuid: "d4e0fbfe-c945-11ee-b8e2-325096b39f47", Username: "test1@test.com", Password: "test", Units: "imperial", CycleStart: "Monday", CycleDays: 7}, false},
		{"read invliad", fields{d}, args{"asdf@test.com"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &IUserRepository{
				DB: tt.fields.DB,
			}
			got, err := db.Read(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("IUserRepository.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IUserRepository.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/