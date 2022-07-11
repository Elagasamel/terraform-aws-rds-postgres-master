package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
)

// These const are declared according to what is found at "../../examples/minimal/main.tf"
const (
	expectedPw       = "examplePassword"
	expectedUsername = "exampleUsername"
	expectedDBName   = "example0"
)

// initTestCases initializes a list of RdsTestCase
func initTestCases() []RdsTestCase {
	return []RdsTestCase{
		{
			testName:         "minimal",
			expectApplyError: false,
			vars: map[string]interface{}{
				"vpc_cidr":            "172.18.0.0/18",
				"database_subnets":    []string{"172.18.0.0/24", "172.18.1.0/24"},
				"egress_cidr_blocks":  []string{"0.0.0.0/0"},
				"ingress_cidr_blocks": []string{"0.0.0.0/0"},
				"name_prefix":         "",
			},
		},
	}
}

// TestTerraformCreateRDS runs all test cases
func TestTerraformCreateRDS(t *testing.T) {

	testCases := initTestCases()

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			// These will create a tempTestFolder for each bucketTestCase.
			tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "test_examples/minimal")

			// this stage will generate a random `awsRegion` and a `uniqueId` to be used in tests.
			test_structure.RunTestStage(t, "pick_new_randoms", func() {
				usRegions := []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2"}
				// This function will first check for the Env Var TERRATEST_REGION and return its value if != ""
				awsRegion := aws.GetRandomStableRegion(t, usRegions, nil)

				test_structure.SaveString(t, tempTestFolder, "region", awsRegion)
				test_structure.SaveString(t, tempTestFolder, "unique_id", strings.ToLower(random.UniqueId()))
			})

			defer test_structure.RunTestStage(t, "teardown", func() {
				teraformOptions := test_structure.LoadTerraformOptions(t, tempTestFolder)
				terraform.Destroy(t, teraformOptions)
			})

			test_structure.RunTestStage(t, "setup_options", func() {
				awsRegion := test_structure.LoadString(t, tempTestFolder, "region")
				uniqueID := test_structure.LoadString(t, tempTestFolder, "unique_id")

				testCase.vars["name_prefix"] = uniqueID

				terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
					TerraformDir: tempTestFolder,
					Vars:         testCase.vars,
					EnvVars: map[string]string{
						"AWS_REGION": awsRegion,
					},
				})

				test_structure.SaveTerraformOptions(t, tempTestFolder, terraformOptions)
			})

			test_structure.RunTestStage(t, "create_rds", func() {
				terraformOptions := test_structure.LoadTerraformOptions(t, tempTestFolder)
				_, err := terraform.InitAndApplyE(t, terraformOptions)

				if testCase.expectApplyError {
					require.Error(t, err)
					// If it failed as expected, we should skip the rest (validate function).
					t.SkipNow()
				}
			})

			test_structure.RunTestStage(t, "validate", func() {
				awsRegion := test_structure.LoadString(t, tempTestFolder, "region")
				terraformOptions := test_structure.LoadTerraformOptions(t, tempTestFolder)
				validateModuleOutputs(t,
					terraformOptions,
					awsRegion,
					int64(5432),
					expectedUsername,
					expectedDBName,
				)
			})
		})
	}
}
