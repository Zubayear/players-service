package repository

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoDB_Save(t *testing.T) {
	type fields struct {
		collection mongo.Collection
	}
	type args struct {
		ctx    context.Context
		entity interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "DBSaveTest",
			fields: fields{
				collection: mongo.Collection{},
			},
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &MongoDB{
				collection: tt.fields.collection,
			}
			got, err := md.Save(tt.args.ctx, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoDB.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
