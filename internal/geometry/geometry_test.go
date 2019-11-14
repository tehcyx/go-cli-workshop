package geometry

import (
	"reflect"
	"testing"
)

func TestNewRect(t *testing.T) {
	type args struct {
		width  float64
		height float64
	}
	tests := []struct {
		name string
		args args
		want Rect
	}{
		{name: "When I pass 1.0 and 2.0 I should get a rectangle with these values", args: args{width: 1.0, height: 2.0}, want: Rect{Width: 1.0, Height: 2.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRect(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCircle(t *testing.T) {
	type args struct {
		radius float64
	}
	tests := []struct {
		name string
		args args
		want Circle
	}{
		{name: "When I pass 2.0 I should get a rectangle with this values", args: args{radius: 2.0}, want: Circle{Radius: 2.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCircle(tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
