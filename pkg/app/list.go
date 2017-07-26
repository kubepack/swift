package app

import (
	"io"

	app "github.com/appscode/wheel/pkg/apis/app/v1beta1"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/release"
)

type listCmd struct {
	filter     string
	short      bool
	limit      int
	offset     string
	byDate     bool
	sortDesc   bool
	out        io.Writer
	all        bool
	deleted    bool
	deleting   bool
	deployed   bool
	failed     bool
	namespace  string
	superseded bool
	pending    bool
	client     helm.Interface
}

func (l *listCmd) run() (app.ListReleasesResponse, error) {
	sortBy := app.ListSort_NAME
	if l.byDate {
		sortBy = app.ListSort_LAST_RELEASED
	}

	sortOrder := app.ListSort_ASC
	if l.sortDesc {
		sortOrder = app.ListSort_DESC
	}

	stats := l.statusCodes()

	res, err := l.client.ListReleases(
		helm.ReleaseListLimit(l.limit),
		helm.ReleaseListOffset(l.offset),
		helm.ReleaseListFilter(l.filter),
		helm.ReleaseListSort(int32(sortBy)),
		helm.ReleaseListOrder(int32(sortOrder)),
		helm.ReleaseListStatuses(stats),
		helm.ReleaseListNamespace(l.namespace),
	)

	if err != nil {
		return nil, prettyError(err)
	}

	return app.ListReleasesResponse(res), nil
}

// statusCodes gets the list of status codes that are to be included in the results.
func (l *listCmd) statusCodes() []release.Status_Code {
	if l.all {
		return []release.Status_Code{
			release.Status_UNKNOWN,
			release.Status_DEPLOYED,
			release.Status_DELETED,
			release.Status_DELETING,
			release.Status_FAILED,
			release.Status_PENDING_INSTALL,
			release.Status_PENDING_UPGRADE,
			release.Status_PENDING_ROLLBACK,
		}
	}
	status := []release.Status_Code{}
	if l.deployed {
		status = append(status, release.Status_DEPLOYED)
	}
	if l.deleted {
		status = append(status, release.Status_DELETED)
	}
	if l.deleting {
		status = append(status, release.Status_DELETING)
	}
	if l.failed {
		status = append(status, release.Status_FAILED)
	}
	if l.superseded {
		status = append(status, release.Status_SUPERSEDED)
	}
	if l.pending {
		status = append(status, release.Status_PENDING_INSTALL, release.Status_PENDING_UPGRADE, release.Status_PENDING_ROLLBACK)
	}

	// Default case.
	if len(status) == 0 {
		status = append(status, release.Status_DEPLOYED, release.Status_FAILED)
	}
	return status
}
