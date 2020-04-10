package template

import "testing"

func TestCreateQtProjectTemplate(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TEST_NAME",
			args:    args{name: "daq"},
			wantErr: false,
		},
		{
			name:    "TEST_EMPTY_NAME",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateQtProjectTemplate(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("CreateQtProjectTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
