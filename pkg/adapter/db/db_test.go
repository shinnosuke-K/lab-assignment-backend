package db

import "testing"

func TestInsertLabs(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "教授情報挿入",
			args:    args{path: "../../lab.csv"},
			wantErr: false,
		},
		{
			name:    "違うファイルパス",
			args:    args{path: "lab.csv"},
			wantErr: true,
		},
		{
			name:    "csvファイルではない",
			args:    args{path: "../../main.go"},
			wantErr: true,
		},
	}

	Open()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertLabs(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("InsertLabs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertUsers(t *testing.T) {
	Open()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "学生情報挿入",
			args:    args{path: "../../student.csv"},
			wantErr: false,
		},
		{
			name:    "違うファイルパス",
			args:    args{path: "student.csv"},
			wantErr: true,
		},
		{
			name:    "csvファイルではない",
			args:    args{path: "../../main.go"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertUsers(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("InsertUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
