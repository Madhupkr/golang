// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snippets

// [START iam_rename_service_account]
import (
	"context"
	"fmt"
	"io"

	iam "google.golang.org/api/iam/v1"
)

// renameServiceAccount renames a service account.
func renameServiceAccount(w io.Writer, email, newDisplayName string) (*iam.ServiceAccount, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %w", err)
	}

	// First, get a ServiceAccount using List() or Get().
	resource := "projects/-/serviceAccounts/" + email
	serviceAccount, err := service.Projects.ServiceAccounts.Get(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Get: %w", err)
	}
	// Then you can update the display name.
	serviceAccount.DisplayName = newDisplayName
	serviceAccount, err = service.Projects.ServiceAccounts.Update(resource, serviceAccount).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Update: %w", err)
	}

	fmt.Fprintf(w, "Updated service account: %v", serviceAccount.Email)
	return serviceAccount, nil
}

// [END iam_rename_service_account]
