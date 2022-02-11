package service

import (
	"context"
	"github.com/OpenCal-FYDP/PreferenceManagement/internal/storage"
	"github.com/OpenCal-FYDP/PreferenceManagement/rpc"
)

type PreferencesServer struct {
	storage *storage.Storage
}

func (p *PreferencesServer) GetUserProfile(ctx context.Context, req *rpc.GetUserProfileReq) (*rpc.GetUserProfileRes, error) {
	res := new(rpc.GetUserProfileRes)
	err := p.storage.GetUserProfile(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PreferencesServer) GetAvailability(ctx context.Context, req *rpc.GetAvailabilityReq) (*rpc.GetAvailabilityRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) GetBooking(ctx context.Context, req *rpc.GetBookingReq) (*rpc.GetBookingRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) SetUserProfile(ctx context.Context, req *rpc.SetUserProfileReq) (*rpc.SetUserProfileRes, error) {
	res := new(rpc.SetUserProfileRes)
	err := p.storage.SetUserProfile(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PreferencesServer) SetAvailability(ctx context.Context, req *rpc.SetAvailabilityReq) (*rpc.SetAvailabilityRes, error) {
	panic("implement me")
}

func (p *PreferencesServer) SetBooking(ctx context.Context, req *rpc.SetBookingReq) (*rpc.SetBookingRes, error) {
	panic("implement me")
}

func New(s *storage.Storage) *PreferencesServer {
	return &PreferencesServer{
		s,
	}
}
