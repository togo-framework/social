// Package social is the Social content pack for togo: SocialLink resource for footer/nav profile links.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package social

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("social", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
