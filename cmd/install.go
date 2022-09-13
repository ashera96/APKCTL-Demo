/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package cmd

import (
	"github.com/BLasan/APKCTL-Demo/impl"
	"github.com/BLasan/APKCTL-Demo/utils"
	"github.com/spf13/cobra"
)

const InstallPlatformCmdLiteral = "install platform"
const InstallPlatformCmdShortDesc = "Install APIM Control Plane component(s) and Data Plane component(s)"
const InstallPlatformCmdLongDesc = `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`
const InstallPlatformCmdExamples = utils.ProjectName + ` ` + InstallPlatformCmdLiteral

// InstallPlatformCmd represents the APKCTL platform installation command
var InstallPlatformCmd = &cobra.Command{
	Use:     InstallPlatformCmdLiteral,
	Short:   InstallPlatformCmdShortDesc,
	Long:    InstallPlatformCmdLongDesc,
	Example: InstallPlatformCmdExamples,
	Run: func(cmd *cobra.Command, args []string) {
		handleInstallPlatform()
	},
}

func handleInstallPlatform() {
	impl.InstallPlatform()
}

func init() {

}
