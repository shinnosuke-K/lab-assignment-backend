package datastore

import (
	"testing"

	"github.com/shinnosuke-K/lab-assignment-backend/pkg/adapter/db"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/repository"
)

func TestUserStore_GetByID(t *testing.T) {

	db.Open()

	type args struct {
		db         repository.DBHandler
		studentNum string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "存在するデータを取得",
			args: args{
				db:         db.Driver,
				studentNum: "1",
			},
			wantErr: false,
		},
		{
			name: "存在しないデータ",
			args: args{
				db:         db.Driver,
				studentNum: "0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserStore{}
			_, err := us.GetByID(tt.args.db, tt.args.studentNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserStore_UpdateEntered(t *testing.T) {

	db.Open()

	type args struct {
		db         repository.DBHandler
		studentNum int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "存在するデータを取得",
			args: args{
				db:         db.Driver,
				studentNum: 1,
			},
			wantErr: false,
		},
		{
			name: "存在しないデータ",
			args: args{
				db:         db.Driver,
				studentNum: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserStore{}
			if err := us.UpdateEntered(tt.args.db, tt.args.studentNum); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEntered() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserStore_UpdateGraduate(t *testing.T) {

	db.Open()

	type args struct {
		db         repository.DBHandler
		studentNum int
		graduate   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "データを更新(false->true)",
			args: args{
				db:         db.Driver,
				studentNum: 1,
				graduate:   true,
			},
			wantErr: false,
		},
		{
			name: "データを更新(true->false)",
			args: args{
				db:         db.Driver,
				studentNum: 2,
				graduate:   false,
			},
			wantErr: false,
		},

		{
			name: "存在しないデータ",
			args: args{
				db:         db.Driver,
				studentNum: 0,
				graduate:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserStore{}
			if err := us.UpdateGraduate(tt.args.db, tt.args.studentNum, tt.args.graduate); (err != nil) != tt.wantErr {
				t.Errorf("UpdateGraduate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
