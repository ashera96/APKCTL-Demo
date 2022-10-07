package integration

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/BLasan/APKCTL-Demo/integration/adminservices"
	"github.com/BLasan/APKCTL-Demo/integration/base"
	testutils "github.com/BLasan/APKCTL-Demo/integration/testUtils"
)

var testCaseUsers = []testutils.TestCaseUsers{
	{
		Description:   "APKCTL user admin Super Tenant",
		ApiCreator:    testutils.Credentials{Username: creator.UserName, Password: creator.Password},
		ApiPublisher:  testutils.Credentials{Username: publisher.UserName, Password: publisher.Password},
		ApiSubscriber: testutils.Credentials{Username: subscriber.UserName, Password: subscriber.Password},
		Admin:         testutils.Credentials{Username: adminservices.AdminUsername, Password: adminservices.AdminPassword},
		CtlUser:       testutils.Credentials{Username: adminservices.AdminUsername, Password: adminservices.AdminPassword},
	},
	{
		Description:   "APKCTL user admin Tenant",
		ApiCreator:    testutils.Credentials{Username: creator.UserName + "@" + TENANT1, Password: creator.Password},
		ApiPublisher:  testutils.Credentials{Username: publisher.UserName + "@" + TENANT1, Password: publisher.Password},
		ApiSubscriber: testutils.Credentials{Username: subscriber.UserName + "@" + TENANT1, Password: subscriber.Password},
		Admin:         testutils.Credentials{Username: adminservices.AdminUsername + "@" + TENANT1, Password: adminservices.AdminPassword},
		CtlUser:       testutils.Credentials{Username: adminservices.AdminUsername + "@" + TENANT1, Password: adminservices.AdminPassword},
	},
	{
		Description:   "APKCTL user devops Super Tenant",
		ApiCreator:    testutils.Credentials{Username: creator.UserName, Password: creator.Password},
		ApiPublisher:  testutils.Credentials{Username: publisher.UserName, Password: publisher.Password},
		ApiSubscriber: testutils.Credentials{Username: subscriber.UserName, Password: subscriber.Password},
		Admin:         testutils.Credentials{Username: adminservices.AdminUsername, Password: adminservices.AdminPassword},
		CtlUser:       testutils.Credentials{Username: devops.UserName, Password: devops.Password},
	},
	{
		Description:   "APKCTL user devops Tenant",
		ApiCreator:    testutils.Credentials{Username: creator.UserName + "@" + TENANT1, Password: creator.Password},
		ApiPublisher:  testutils.Credentials{Username: publisher.UserName + "@" + TENANT1, Password: publisher.Password},
		ApiSubscriber: testutils.Credentials{Username: subscriber.UserName + "@" + TENANT1, Password: subscriber.Password},
		Admin:         testutils.Credentials{Username: adminservices.AdminUsername + "@" + TENANT1, Password: adminservices.AdminPassword},
		CtlUser:       testutils.Credentials{Username: devops.UserName + "@" + TENANT1, Password: devops.Password},
	},
}

var (
	creator    = Users["creator"][0]
	subscriber = Users["subscriber"][0]
	publisher  = Users["publisher"][0]
	devops     = Users["devops"][0]
)

var Users = map[string][]adminservices.User{
	"creator":    {{UserName: adminservices.CreatorUsername, Password: adminservices.Password, Roles: []string{"Internal/creator"}}},
	"publisher":  {{UserName: adminservices.PublisherUsername, Password: adminservices.Password, Roles: []string{"Internal/publisher"}}},
	"subscriber": {{UserName: adminservices.SubscriberUsername, Password: adminservices.Password, Roles: []string{"Internal/subscriber"}}},
	"devops":     {{UserName: adminservices.DevopsUsername, Password: adminservices.Password, Roles: []string{"Internal/devops"}}},
}

const (
	superAdminUser     = adminservices.AdminUsername
	superAdminPassword = adminservices.AdminPassword

	userMgtService   = "RemoteUserStoreManagerService"
	tenantMgtService = "TenantMgtAdminService"

	DEFAULT_TENANT_DOMAIN = adminservices.DefaultTenantDomain
	TENANT1               = adminservices.Tenant1
)

func TestMain(m *testing.M) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	base.RelativeBinaryPath = dir + "/../"
	base.RelativeTargetDirPath = dir

	deployBackendService()

	m.Run()

	removeBackendService()
}

func deployBackendService() {

	args := []string{"apply", "-f", filepath.Join(base.RelativeTargetDirPath, "testData/BackendService.yaml")}
	fmt.Println("Args: ", args)
	_, err := base.ExecuteKubernetesCommands(args...)
	if err != nil {
		fmt.Println("Error while deploying the Backend Service")
	}

	// fmt.Println("Output: ", out)
}

func removeBackendService() {
	args := []string{"delete", "-f", filepath.Join(base.RelativeTargetDirPath, "testData/BackendService.yaml")}
	_, err := base.ExecuteKubernetesCommands(args...)
	if err != nil {
		fmt.Println("Error while removing the Backeend Service")
	}
}
