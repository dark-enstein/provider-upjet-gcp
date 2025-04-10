// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-gcp/internal/controller/appengine/application"
	applicationurldispatchrules "github.com/upbound/provider-gcp/internal/controller/appengine/applicationurldispatchrules"
	domainmapping "github.com/upbound/provider-gcp/internal/controller/appengine/domainmapping"
	firewallrule "github.com/upbound/provider-gcp/internal/controller/appengine/firewallrule"
	flexibleappversion "github.com/upbound/provider-gcp/internal/controller/appengine/flexibleappversion"
	servicenetworksettings "github.com/upbound/provider-gcp/internal/controller/appengine/servicenetworksettings"
	servicesplittraffic "github.com/upbound/provider-gcp/internal/controller/appengine/servicesplittraffic"
	standardappversion "github.com/upbound/provider-gcp/internal/controller/appengine/standardappversion"
)

// Setup_appengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationurldispatchrules.Setup,
		domainmapping.Setup,
		firewallrule.Setup,
		flexibleappversion.Setup,
		servicenetworksettings.Setup,
		servicesplittraffic.Setup,
		standardappversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
