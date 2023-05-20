package tools

//go:generate mockgen -package security -destination ../pkg/security/mocks.go github.com/guidomantilla/go-feather-web/pkg/security AuthenticationEndpoint,AuthenticationService,AuthenticationDelegate,AuthorizationFilter,AuthorizationService,AuthorizationDelegate,PrincipalManager,TokenManager
//go:generate mockgen -package server -destination ../pkg/server/mocks.go github.com/qmdx00/lifecycle Server
//go:generate mockgen -package server -destination ../pkg/server/mocks.go github.com/qmdx00/lifecycle Server
