package dnd_dice

import (
	"reflect"
	"testing"
)

func TestD10(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D10(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D10() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D10() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD100(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D100(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D100() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D100() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD12(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D12(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D12() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D12() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD20(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D20(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D20() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D20() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD4(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D4(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D4() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD6(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D6(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D6() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D6() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD8(t *testing.T) {
	type args struct {
		rollTimes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{21}, 21, false},
		{"Invalid roll times less than 0", args{-5}, 0, true},
		{"Invalid roll times equal to 0", args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := D8(tt.args.rollTimes)
			if (err != nil) != tt.wantErr {
				t.Errorf("D8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("D8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDice(t *testing.T) {
	type args struct {
		rollTimes int
		diceSize  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid values", args{3, 6}, 3, false},
		{"Invalid roll times equal to 0", args{0, 6}, 0, true},
		{"Invalid roll times less than 0", args{-5, 8}, 0, true},
		{"Invalid dice size equal to 0", args{7, 0}, 0, true},
		{"Invalid dice size less than 0", args{2, -1}, 0, true},
		{"Invalid roll times and dice size less than 0", args{-6, -9}, 0, true},
		{"Invalid roll times and dice size equal to 0", args{0, 0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Dice(tt.args.rollTimes, tt.args.diceSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("Dice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiceExpression(t *testing.T) {
	type args struct {
		diceExpression string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Valid dice expression", args{"4d6"}, 4, false},
		{"Invalid dice expression roll times less than 0", args{"-5d20"}, 0, true},
		{"Invalid dice expression roll times equal to 0", args{"0d20"}, 0, true},
		{"Invalid dice expression dice size less than 0", args{"3d-6"}, 0, true},
		{"Invalid dice expression dice size equal to 0", args{"3d0"}, 0, true},
		{"Invalid dice expression roll times and dice size less than 0", args{"-7d-20"}, 0, true},
		{"Invalid dice expression roll times less than 0 and dice size equal 0", args{"-7d0"}, 0, true},
		{"Invalid dice expression roll times equal to 0 and dice size less than 0", args{"0d-20"}, 0, true},
		{"Invalid dice expression roll times and dice size equal to 0", args{"0d0"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DiceExpression(tt.args.diceExpression)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiceExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("DiceExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}
