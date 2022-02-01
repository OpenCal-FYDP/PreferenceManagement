package service

import (
	"context"
	"github.com/OpenCal-FYDP/PreferenceManagement/rpc"
)

type PreferencesServer struct{}

func (p *PreferencesServer) GetUserProfile(ctx context.Context, req *rpc.GetUserProfileReq) (*rpc.GetUserProfileRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) GetAvailability(ctx context.Context, req *rpc.GetAvailabilityReq) (*rpc.GetAvailabilityRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) GetBooking(ctx context.Context, req *rpc.GetBookingReq) (*rpc.GetBookingRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) SetUserProfile(ctx context.Context, req *rpc.SetUserProfileReq) (*rpc.SetUserProfileRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) SetAvailability(ctx context.Context, req *rpc.SetAvailabilityReq) (*rpc.SetAvailabilityRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) SetBooking(ctx context.Context, req *rpc.SetBookingReq) (*rpc.SetBookingRes, error) {
	panic("implement me")
}

func New() *PreferencesServer {
	return &PreferencesServer{}
}
