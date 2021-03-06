/*
Copyright © 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package aid assist Cobra commands individually. Example: a command under `cmd/foo.go`, has its respective `aid/foo.go`.
// This allows a more cleaner and readable code.
package aid

import (
	"github.com/spf13/cobra"
)

// SetApplyFlags define flags for the cobra command
func SetApplyFlags(cmd *cobra.Command) {
	usage := `The Apply ID`
	cmd.Flags().String("id", "", usage)
}
