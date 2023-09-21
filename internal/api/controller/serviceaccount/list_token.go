// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serviceaccount

import (
	"context"

	apiauth "github.com/harness/gitness/internal/api/auth"
	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

// ListTokens lists all tokens of a service account.
func (c *Controller) ListTokens(ctx context.Context, session *auth.Session,
	saUID string) ([]*types.Token, error) {
	sa, err := findServiceAccountFromUID(ctx, c.principalStore, saUID)
	if err != nil {
		return nil, err
	}

	// Ensure principal has required permissions on parent (ensures that parent exists)
	if err = apiauth.CheckServiceAccount(ctx, c.authorizer, session, c.spaceStore, c.repoStore,
		sa.ParentType, sa.ParentID, sa.UID, enum.PermissionServiceAccountView); err != nil {
		return nil, err
	}

	return c.tokenStore.List(ctx, sa.ID, enum.TokenTypeSAT)
}
