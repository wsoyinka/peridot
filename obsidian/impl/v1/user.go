// Copyright (c) All respective contributors to the Peridot Project. All rights reserved.
// Copyright (c) 2021-2022 Rocky Enterprise Software Foundation, Inc. All rights reserved.
// Copyright (c) 2021-2022 Ctrl IQ, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package obsidianimplv1

import (
	"context"
	"github.com/ory/hydra-client-go/client/admin"
	hydramodels "github.com/ory/hydra-client-go/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	obsidianpb "peridot.resf.org/obsidian/pb"
)

const (
	invalidCheckType        = "invalid_check_type"
	challengeRequired       = "challenge_required"
	logoutChallengeRequired = "logout_challenge_required"
)

func (s *Server) SessionStatus(_ context.Context, req *obsidianpb.SessionStatusRequest) (*obsidianpb.SessionStatusResponse, error) {
	allowedTypes := map[string]bool{"login": true, "consent": true}
	if !allowedTypes[req.CheckType] {
		return nil, status.Error(codes.InvalidArgument, invalidCheckType)
	}

	if req.Challenge == "" {
		return nil, status.Error(codes.InvalidArgument, challengeRequired)
	}

	var res *obsidianpb.SessionStatusResponse
	var err error
	switch req.CheckType {
	case "login":
		res, err = s.ProcessLoginRequest(req.Challenge)
		break
	case "consent":
		res, err = s.ProcessConsentRequest(req.Challenge)
		break
	default:
		return nil, status.Error(codes.InvalidArgument, "")
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) ConsentDecision(ctx context.Context, req *obsidianpb.ConsentDecisionRequest) (*obsidianpb.ConsentDecisionResponse, error) {
	if req.Challenge == "" {
		return nil, status.Error(codes.InvalidArgument, challengeRequired)
	}

	consentReq, err := s.hydra.Admin.GetConsentRequest(&admin.GetConsentRequestParams{
		Context:          ctx,
		ConsentChallenge: req.Challenge,
	})
	if err != nil {
		s.log.Error(err)
		return nil, status.Error(codes.Internal, authError)
	}

	var redirectURL string

	if req.Allow {
		res, err := s.AcceptConsentRequest(ctx, req.Challenge, consentReq)
		if err != nil {
			return nil, err
		}
		if !res.Valid {
			return nil, status.Error(codes.InvalidArgument, "invalid consent request")
		}
		redirectURL = res.RedirectUrl
	} else {
		res, err := s.hydra.Admin.RejectConsentRequest(&admin.RejectConsentRequestParams{
			Body: &hydramodels.RejectRequest{
				StatusCode:       int64(codes.Aborted),
				ErrorDescription: "User does not consent to data sharing",
				Error:            "no_consent",
			},
			ConsentChallenge: req.Challenge,
			Context:          ctx,
		})
		if err != nil {
			s.log.Errorf("error rejecting consent request: %s", err.Error())
			return nil, status.Error(codes.Internal, authError)
		}
		redirectURL = *res.Payload.RedirectTo
	}

	return &obsidianpb.ConsentDecisionResponse{
		RedirectUrl: redirectURL,
	}, nil
}

func (s *Server) LogoutDecision(ctx context.Context, req *obsidianpb.LogoutDecisionRequest) (*obsidianpb.LogoutDecisionResponse, error) {
	if req.Challenge == "" {
		return nil, status.Error(codes.InvalidArgument, logoutChallengeRequired)
	}

	logout, err := s.hydra.Admin.GetLogoutRequest(&admin.GetLogoutRequestParams{
		LogoutChallenge: req.Challenge,
		Context:         ctx,
	})
	if err != nil {
		s.log.Error(err)
		return nil, status.Error(codes.Internal, authError)
	}

	var redirectURL string
	if req.Accept {
		_, err = s.hydra.Admin.RevokeConsentSessions(&admin.RevokeConsentSessionsParams{
			All:     proto.Bool(true),
			Subject: logout.Payload.Subject,
			Context: ctx,
		})
		if err != nil {
			s.log.Errorf("error revoking consent sessions: %s", err.Error())
			return nil, status.Error(codes.Internal, "error revoking consent sessions")
		}

		acceptReq, err := s.hydra.Admin.AcceptLogoutRequest(&admin.AcceptLogoutRequestParams{
			LogoutChallenge: req.Challenge,
			Context:         ctx,
		})
		if err != nil {
			s.log.Errorf("error accepting logout request: %s", err.Error())
			return nil, status.Error(codes.Internal, "error accepting logout request")
		}
		redirectURL = *acceptReq.Payload.RedirectTo
	} else {
		_, err := s.hydra.Admin.RejectLogoutRequest(&admin.RejectLogoutRequestParams{
			Body: &hydramodels.RejectRequest{
				StatusCode:       int64(codes.Aborted),
				ErrorDescription: "User denied logout request",
				Error:            "user_denied",
			},
			LogoutChallenge: req.Challenge,
			Context:         nil,
			HTTPClient:      nil,
		})
		if err != nil {
			s.log.Errorf("error rejecting logout request: %s", err.Error())
			return nil, status.Error(codes.Internal, "error rejecting logout request")
		}
		redirectURL = logout.Payload.RequestURL
	}

	return &obsidianpb.LogoutDecisionResponse{
		RedirectUrl: redirectURL,
	}, nil
}
