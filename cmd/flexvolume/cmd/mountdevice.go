/*
Copyright 2017 The Rook Authors. All rights reserved.

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

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	mountDeviceCmd = &cobra.Command{
		Use:   "mountdevice",
		Short: "Mounts the attached device to a mount directory",
		RunE:  mountDevice,
	}
)

func init() {
	RootCmd.AddCommand(mountDeviceCmd)
}

func mountDevice(cmd *cobra.Command, args []string) error {
	return new(NotSupportedError)
}
