package unitypes_test

import (
	"encoding/json"
	"fmt"
	"github.com/mistermelphin/go-unitypes"
	"reflect"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	type subtest struct {
		wantErr bool
		want    any
	}
	tests := []struct {
		name     string
		jsn      string
		subtests []subtest
	}{
		// INT
		{
			name: "string is integer",
			jsn:  "\"123\"",
			subtests: []subtest{
				{
					wantErr: false,
					want:    unitypes.Int(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int8(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int16(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int32(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int64(123),
				},
				{
					wantErr: false,
					want:    unitypes.Float32(123),
				},
				{
					wantErr: false,
					want:    unitypes.Float64(123),
				},
			},
		},
		{
			name: "string is no integer",
			jsn:  "\"abd\"",
			subtests: []subtest{
				{
					wantErr: true,
					want:    unitypes.Int(0),
				},
				{
					wantErr: true,
					want:    unitypes.Int8(0),
				},
				{
					wantErr: true,
					want:    unitypes.Int16(0),
				},
				{
					wantErr: true,
					want:    unitypes.Int32(0),
				},
				{
					wantErr: true,
					want:    unitypes.Int64(0),
				},
				{
					wantErr: true,
					want:    unitypes.Float32(0),
				},
				{
					wantErr: true,
					want:    unitypes.Float64(0),
				},
			},
		},
		{
			name: "string is float",
			jsn:  "\"123.345\"",
			subtests: []subtest{
				{
					wantErr: false,
					want:    unitypes.Int(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int8(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int16(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int32(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int64(123),
				},
				{
					wantErr: false,
					want:    unitypes.Float32(123.345),
				},
				{
					wantErr: false,
					want:    unitypes.Float64(123.345),
				},
			},
		},
		{
			name: "integer",
			jsn:  "123",
			subtests: []subtest{
				{
					wantErr: false,
					want:    unitypes.Int(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int8(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int16(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int32(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int64(123),
				},
				{
					wantErr: false,
					want:    unitypes.Float32(123),
				},
				{
					wantErr: false,
					want:    unitypes.Float64(123),
				},
			},
		},
		{
			name: "float",
			jsn:  "123.456",
			subtests: []subtest{
				{
					wantErr: false,
					want:    unitypes.Int(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int8(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int16(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int32(123),
				},
				{
					wantErr: false,
					want:    unitypes.Int64(123),
				},
				{
					wantErr: false,
					want:    unitypes.Float32(123.456),
				},
				{
					wantErr: false,
					want:    unitypes.Float64(123.456),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, ttt := range tt.subtests {
				typ := reflect.TypeOf(ttt.want)
				t.Run(fmt.Sprintf("to %v", typ.Name()), func(t *testing.T) {
					ii := reflect.New(typ).Interface()
					if err := json.Unmarshal([]byte(tt.jsn), ii); (err != nil) != ttt.wantErr {
						t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, ttt.wantErr)
						return
					}
					ii = reflect.ValueOf(ii).Elem().Interface()
					if ii != ttt.want {
						t.Errorf("UnmarshalJSON() got = %v, want %v", ii, ttt.want)
					}
					t.Logf("%v -> %v (%v)", tt.jsn, ii, reflect.TypeOf(ii).Name())
				})
			}
		})
	}
}

func ptr[T comparable](v T) *T {
	return &v
}

func TestJsonStructure(t *testing.T) {
	type testObj struct {
		I  unitypes.Int  `json:"i"`
		Ip *unitypes.Int `json:"ip"`
	}
	tests := []struct {
		name string
		jsn  string
		want testObj
	}{
		{
			name: "all fields",
			jsn:  `{"i":"123","ip":"1234"}`,
			want: testObj{
				I:  123,
				Ip: ptr[unitypes.Int](1234),
			},
		},
		{
			name: "with nil",
			jsn:  `{"i":"123"}`,
			want: testObj{
				I: 123,
			},
		},
		{
			name: "with no fields",
			jsn:  "{}",
			want: testObj{},
		},
		{
			name: "with null",
			jsn:  `{"ip": null}`,
			want: testObj{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := testObj{}
			if err := json.Unmarshal([]byte(tt.jsn), &res); err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("aren't different")
			}
			t.Logf("'%v' -> %+v", tt.jsn, res)
		})
	}
}

func TestJsonMarshaling(t *testing.T) {
	obj := struct {
		I   unitypes.Int  `json:"i"`
		Ip  *unitypes.Int `json:"ip"`
		Ipn *unitypes.Int `json:"ipn"`
	}{
		I:  123,
		Ip: ptr[unitypes.Int](1234),
	}
	res, err := json.Marshal(&obj)
	if err != nil {
		t.Error(err)
	}
	want := `{"i":123,"ip":1234,"ipn":null}`
	if !reflect.DeepEqual(want, string(res)) {
		t.Errorf("aren't equal. got = %v, want = %v", string(res), want)
	}
}
