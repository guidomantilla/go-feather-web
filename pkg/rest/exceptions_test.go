package rest

import (
	"reflect"
	"testing"
)

func TestBadRequestException(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name string
		args args
		want *Exception
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadRequestException(tt.args.message, tt.args.err...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BadRequestException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnauthorizedException(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name string
		args args
		want *Exception
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnauthorizedException(tt.args.message, tt.args.err...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnauthorizedException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotFoundException(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name string
		args args
		want *Exception
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotFoundException(tt.args.message, tt.args.err...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotFoundException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternalServerErrorException(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name string
		args args
		want *Exception
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InternalServerErrorException(tt.args.message, tt.args.err...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalServerErrorException() = %v, want %v", got, tt.want)
			}
		})
	}
}
