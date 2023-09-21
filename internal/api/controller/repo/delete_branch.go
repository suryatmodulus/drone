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

package repo

import (
	"context"
	"fmt"

	"github.com/harness/gitness/gitrpc"
	"github.com/harness/gitness/internal/api/usererror"
	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/types/enum"
)

// DeleteBranch deletes a repo branch.
func (c *Controller) DeleteBranch(ctx context.Context,
	session *auth.Session,
	repoRef string,
	branchName string,
) error {
	repo, err := c.getRepoCheckAccess(ctx, session, repoRef, enum.PermissionRepoPush, false)
	if err != nil {
		return err
	}

	// make sure user isn't deleting the default branch
	// ASSUMPTION: lower layer calls explicit branch api
	// and 'refs/heads/branch1' would fail if 'branch1' exists.
	// TODO: Add functional test to ensure the scenario is covered!
	if branchName == repo.DefaultBranch {
		return usererror.ErrDefaultBranchCantBeDeleted
	}

	writeParams, err := CreateRPCWriteParams(ctx, c.urlProvider, session, repo)
	if err != nil {
		return fmt.Errorf("failed to create RPC write params: %w", err)
	}

	err = c.gitRPCClient.DeleteBranch(ctx, &gitrpc.DeleteBranchParams{
		WriteParams: writeParams,
		BranchName:  branchName,
	})
	if err != nil {
		return err
	}

	return nil
}
