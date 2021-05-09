package cowinutils

import (
	"reflect"
	"t2findmyvaccinebot/pkg/common"
	"testing"
)

func TestGetStates(t *testing.T) {
	tests := []struct {
		name    string
		want    common.StateList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"t1", common.StateList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStates()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDistricts(t *testing.T) {
	type args struct {
		stateID string
	}
	tests := []struct {
		name    string
		args    args
		want    common.DistrictList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"t1", args{stateID: "1"}, common.DistrictList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDistricts(tt.args.stateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDistricts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDistricts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSessionByDist(t *testing.T) {
	type args struct {
		ID   string
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    common.SessionList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"t1", args{ID: "140", date: "06-05-2021"}, common.SessionList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSessionByDist(tt.args.ID, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionByDist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSessionByDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSessionByPin(t *testing.T) {
	type args struct {
		pinID string
		date  string
	}
	tests := []struct {
		name    string
		args    args
		want    common.SessionList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"t1", args{pinID: "", date: ""}, common.SessionList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSessionByPin(tt.args.pinID, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionByPin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSessionByPin() = %v, want %v", got, tt.want)
			}
		})
	}
}
