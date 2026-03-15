package workspaces_services

import (
	"dbsystemdata-backend/internal/features/audit_logs"
	"dbsystemdata-backend/internal/features/email"
	users_services "dbsystemdata-backend/internal/features/users/services"
	workspaces_interfaces "dbsystemdata-backend/internal/features/workspaces/interfaces"
	workspaces_repositories "dbsystemdata-backend/internal/features/workspaces/repositories"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	workspaceRepository  = &workspaces_repositories.WorkspaceRepository{}
	membershipRepository = &workspaces_repositories.MembershipRepository{}
)

var workspaceService = &WorkspaceService{
	workspaceRepository,
	membershipRepository,
	users_services.GetUserService(),
	audit_logs.GetAuditLogService(),
	users_services.GetSettingsService(),
	[]workspaces_interfaces.WorkspaceDeletionListener{},
}

var membershipService = &MembershipService{
	membershipRepository,
	workspaceRepository,
	users_services.GetUserService(),
	audit_logs.GetAuditLogService(),
	workspaceService,
	users_services.GetSettingsService(),
	email.GetEmailSMTPSender(),
	logger.GetLogger(),
}

func GetWorkspaceService() *WorkspaceService {
	return workspaceService
}

func GetMembershipService() *MembershipService {
	return membershipService
}
