package fdp

import (
	"reflect"
	"testing"
)

func Test_newPktBuffer(t *testing.T) {
	type args struct {
		cap int32
	}
	tests := []struct {
		name string
		args args
		want *pktBuffer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPktBuffer(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPktBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pktBuffer_Push(t *testing.T) {
	type fields struct {
		ring []packet
		head int32
		tail int32
		cap  int32
	}
	type args struct {
		pkt packet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pktBuffer{
				ring: tt.fields.ring,
				head: tt.fields.head,
				tail: tt.fields.tail,
				cap:  tt.fields.cap,
			}
			if err := b.Push(tt.args.pkt); (err != nil) != tt.wantErr {
				t.Errorf("pktBuffer.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pktBuffer_Pop(t *testing.T) {
	type fields struct {
		ring []packet
		head int32
		tail int32
		cap  int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    packet
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pktBuffer{
				ring: tt.fields.ring,
				head: tt.fields.head,
				tail: tt.fields.tail,
				cap:  tt.fields.cap,
			}
			got, err := b.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("pktBuffer.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pktBuffer.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pktBuffer_Get(t *testing.T) {
	type fields struct {
		ring []packet
		head int32
		tail int32
		cap  int32
	}
	type args struct {
		index int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    packet
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pktBuffer{
				ring: tt.fields.ring,
				head: tt.fields.head,
				tail: tt.fields.tail,
				cap:  tt.fields.cap,
			}
			got, err := b.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("pktBuffer.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pktBuffer.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
