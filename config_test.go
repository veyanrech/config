package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "TestNewConfig",
			args: args{
				filepath: "dummy.test.json",
			},
			want: &config{
				Data: map[string]interface{}{
					"key": "value",
					"topic": map[string]interface{}{
						"key1": "value1",
					},
				},
			},
		},
	}
	for _, tt := range tests {

		got := NewConfig(tt.args.filepath)

		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetTopic(t *testing.T) {
	type fields struct {
		Data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Config
	}{
		{
			name: "Test_config_GetTopic",
			fields: fields{
				Data: map[string]interface{}{
					"key": "value",
					"topic": map[string]interface{}{
						"key1": "value1",
					},
				},
			},
			args: args{
				key: "topic",
			},
			want: &config{
				Data: map[string]interface{}{
					"key1": "value1",
				},
			},
		},
	}
	for _, tt := range tests {
		c := &config{
			Data: tt.fields.Data,
		}
		t.Run(tt.name, func(t *testing.T) {
			got := c.GetTopic(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("config.GetTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}
