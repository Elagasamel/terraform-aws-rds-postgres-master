package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RdsTestCase is a struct for defining tests for Elastic Map Reduce Module
type RdsTestCase struct {
	testName         string
	vars             map[string]interface{}
	expectApplyError bool
}

// validateModuleOutputs validates the values of outputs with expected values
func validateModuleOutputs(t *testing.T, terraformOptions *terraform.Options, awsRegion string, expectedPort int64, expectedUser string, expectedDBName string) {

	oRDS := terraform.OutputMap(t, terraformOptions, "rds")

	oRDSInstanceID := oRDS["rds_postgres_id"]
	oRDSSGIDs := oRDS["rds_security_group_ids"]
	oRDShostname := oRDS["rds_hostname"]
	oRDSport, _ := strconv.ParseInt(oRDS["rds_db_port"], 10, 64)
	oRDSuser := oRDS["rds_username"]
	oDBName := oRDS["rds_dbname"]

	// Fails test if Instance ID is nil
	require.NotNil(t, oRDSInstanceID)

	// Information in RDS API can take more than 20 mins to be available. We retry for 40mins before failing
	rdsObj := retry.DoWithRetryInterface(t, "Waiting RDS API to be available", 20, 2*time.Minute, func() (interface{}, error) {
		return aws.GetRdsInstanceDetailsE(t, oRDSInstanceID, awsRegion)
	}).(*rds.DBInstance)

	t.Run("validate_address", func(t *testing.T) {
		// Verify that the address is not null and equal to output
		address := aws.GetAddressOfRdsInstance(t, oRDSInstanceID, awsRegion)
		require.NotNil(t, address)
		assert.Equal(t, oRDShostname, address)
		assert.Equal(t, oRDShostname, *rdsObj.Endpoint.Address)
	})

	t.Run("validate_port", func(t *testing.T) {
		// Verify that the DB instance is listening on the expected port and equal to output
		port := aws.GetPortOfRdsInstance(t, oRDSInstanceID, awsRegion)
		assert.Equal(t, expectedPort, port)
		assert.Equal(t, int64(oRDSport), port)
	})

	t.Run("validate_SG", func(t *testing.T) {
		// Verify Sec Group IDs output is not nil
		assert.NotNil(t, oRDSSGIDs)
	})

	t.Run("validate_user", func(t *testing.T) {
		// Verify that user is the same as expected and equal to output
		assert.Equal(t, expectedUser, *rdsObj.MasterUsername)
		assert.NotNil(t, oRDSuser)
	})

	t.Run("validate_dbname", func(t *testing.T) {
		// Verify that dbName is the same as expected and equal to output
		assert.Equal(t, expectedDBName, *rdsObj.DBName)
		assert.Equal(t, oDBName, *rdsObj.DBName)
	})

}
