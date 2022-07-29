package service

import (
	"context"
	v1 "github.com/the-zion/matrix-core/api/user/service/v1"
	"github.com/the-zion/matrix-core/app/user/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetAccount(ctx context.Context, req *v1.GetAccountReq) (*v1.GetAccountReply, error) {
	user, err := s.uc.GetAccount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetAccountReply{
		Phone:    user.Phone,
		Email:    user.Email,
		Qq:       user.Qq,
		Wechat:   user.Wechat,
		Weibo:    user.Weibo,
		Github:   user.Github,
		Password: user.Password,
	}, nil
}

func (s *UserService) GetProfile(ctx context.Context, req *v1.GetProfileReq) (*v1.GetProfileReply, error) {
	profile, err := s.uc.GetProfile(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileReply{
		Uuid:      profile.Uuid,
		Username:  profile.Username,
		Avatar:    profile.Avatar,
		School:    profile.School,
		Company:   profile.Company,
		Job:       profile.Job,
		Homepage:  profile.Homepage,
		Introduce: profile.Introduce,
		Created:   profile.Created,
	}, nil
}

func (s *UserService) GetProfileList(ctx context.Context, req *v1.GetProfileListReq) (*v1.GetProfileListReply, error) {
	reply := &v1.GetProfileListReply{Profile: make([]*v1.GetProfileListReply_Profile, 0)}
	profileList, err := s.uc.GetProfileList(ctx, req.Uuids)
	if err != nil {
		return nil, err
	}
	for _, item := range profileList {
		reply.Profile = append(reply.Profile, &v1.GetProfileListReply_Profile{
			Uuid:      item.Uuid,
			Username:  item.Username,
			Introduce: item.Introduce,
		})
	}
	return reply, nil
}

func (s *UserService) GetProfileUpdate(ctx context.Context, req *v1.GetProfileUpdateReq) (*v1.GetProfileUpdateReply, error) {
	profile, err := s.uc.GetProfileUpdate(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileUpdateReply{
		Username:  profile.Username,
		Avatar:    profile.Avatar,
		School:    profile.School,
		Company:   profile.Company,
		Job:       profile.Job,
		Homepage:  profile.Homepage,
		Introduce: profile.Introduce,
		Status:    profile.Status,
	}, nil
}

func (s *UserService) GetUserFollow(ctx context.Context, req *v1.GetUserFollowReq) (*v1.GetUserFollowReply, error) {
	follow, err := s.uc.GetUserFollow(ctx, req.Uuid, req.UserUuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserFollowReply{
		Follow: follow,
	}, nil
}

func (s *UserService) GetFollowList(ctx context.Context, req *v1.GetFollowListReq) (*v1.GetFollowListReply, error) {
	reply := &v1.GetFollowListReply{Follow: make([]*v1.GetFollowListReply_Follow, 0)}
	followList, err := s.uc.GetFollowList(ctx, req.Page, req.Uuid)
	if err != nil {
		return nil, err
	}
	for _, item := range followList {
		reply.Follow = append(reply.Follow, &v1.GetFollowListReply_Follow{
			Uuid: item.Follow,
		})
	}
	return reply, nil
}

func (s *UserService) GetFollowListCount(ctx context.Context, req *v1.GetFollowListCountReq) (*v1.GetFollowListCountReply, error) {
	count, err := s.uc.GetFollowListCount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetFollowListCountReply{
		Count: count,
	}, nil
}

func (s *UserService) GetFollowedList(ctx context.Context, req *v1.GetFollowedListReq) (*v1.GetFollowedListReply, error) {
	reply := &v1.GetFollowedListReply{Follow: make([]*v1.GetFollowedListReply_Follow, 0)}
	followedList, err := s.uc.GetFollowedList(ctx, req.Page, req.Uuid)
	if err != nil {
		return nil, err
	}
	for _, item := range followedList {
		reply.Follow = append(reply.Follow, &v1.GetFollowedListReply_Follow{
			Uuid: item.Followed,
		})
	}
	return reply, nil
}

func (s *UserService) GetFollowedListCount(ctx context.Context, req *v1.GetFollowedListCountReq) (*v1.GetFollowedListCountReply, error) {
	count, err := s.uc.GetFollowedListCount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetFollowedListCountReply{
		Count: count,
	}, nil
}

func (s *UserService) GetUserFollows(ctx context.Context, req *v1.GetUserFollowsReq) (*v1.GetUserFollowsReply, error) {
	reply := &v1.GetUserFollowsReply{Follows: make([]*v1.GetUserFollowsReply_Follows, 0)}
	followsList, err := s.uc.GetUserFollows(ctx, req.UserId, req.Uuids)
	if err != nil {
		return nil, err
	}
	for _, item := range followsList {
		reply.Follows = append(reply.Follows, &v1.GetUserFollowsReply_Follows{
			Uuid:        item.Uuid,
			FollowJudge: item.Follow,
		})
	}
	return reply, nil
}

func (s *UserService) SetProfileUpdate(ctx context.Context, req *v1.SetProfileUpdateReq) (*emptypb.Empty, error) {
	profile := &biz.ProfileUpdate{}
	profile.Uuid = req.Uuid
	profile.Username = req.Username
	profile.School = req.School
	profile.Company = req.Company
	profile.Job = req.Job
	profile.Homepage = req.Homepage
	profile.Introduce = req.Introduce
	err := s.uc.SetProfileUpdate(ctx, profile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) SetUserFollow(ctx context.Context, req *v1.SetUserFollowReq) (*emptypb.Empty, error) {
	err := s.uc.SetUserFollow(ctx, req.Uuid, req.UserUuid)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) CancelUserFollow(ctx context.Context, req *v1.CancelUserFollowReq) (*emptypb.Empty, error) {
	err := s.uc.CancelUserFollow(ctx, req.Uuid, req.UserUuid)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) ProfileReviewPass(ctx context.Context, req *v1.ProfileReviewPassReq) (*emptypb.Empty, error) {
	err := s.uc.ProfileReviewPass(ctx, req.Uuid, req.Update)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) ProfileReviewNotPass(ctx context.Context, req *v1.ProfileReviewNotPassReq) (*emptypb.Empty, error) {
	err := s.uc.ProfileReviewNotPass(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
