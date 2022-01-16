package handler

import (
	"context"
	players "players/proto"
	"players/repository"
	"testing"
)

func TestPlayers_Save(t *testing.T) {
	type fields struct {
		MD repository.IRepository
	}
	type args struct {
		ctx context.Context
		req *players.SaveRequest
		rsp *players.SaveResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				MD: nil,
			},
			args: args{
				ctx: nil,
				req: &players.SaveRequest{
					Name:     "Virgil Van Dijk",
					Position: players.Postiton_DEFENDER,
					ClubName: "Liverpool",
					Age:      30,
				},
				rsp: &players.SaveResponse{
					Id: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Players{
				MD: tt.fields.MD,
			}
			if err := p.Save(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Players.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
