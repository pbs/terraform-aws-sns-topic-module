package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testSNSTopic(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	awsAccountID := getAWSAccountID(t)
	awsRegion := getAWSRegion(t)

	expectedName := fmt.Sprintf("example-tf-sns-topic-%s", variant)

	name := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, expectedName, name)

	expectedARN := fmt.Sprintf("arn:aws:sns:%s:%s:%s", awsRegion, awsAccountID, expectedName)

	arn := terraform.Output(t, terraformOptions, "arn")
	assert.Equal(t, expectedARN, arn)
}
