package dash

import "context"

type Reader interface {
	FindInstitutionByAdminDomain(ctx context.Context, domain string) (*Institution, error)
	FindGroupsByUser(ctx context.Context, email string) ([]*Group, error)
	FindGroupTypes() ([]*GroupType, error)
}

type Writer interface {
	CreateGroup(ctx context.Context, groupType GroupType, opts ...CreateGroupOption) (*Group, error)
	DeleteGroup(ctx context.Context, path string) error
}

type Repository interface {
	Reader
	Writer
}
