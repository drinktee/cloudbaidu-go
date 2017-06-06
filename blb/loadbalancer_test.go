package blb

import (
	"testing"

	"fmt"

	"github.com/drinktee/bce-sdk-go/util"
)

// var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

// //var bceConfig = bce.NewConfig(credentials)
// var bceConfig = &bce.Config{
// 	Credentials: credentials,
// 	Checksum:    true,
// 	Region:      os.Getenv("BOS_REGION"),
// }
// var bccConfig = NewConfig(bceConfig)
// var blbClient = NewBLBClient(bccConfig)

// func TestCredentials(t *testing.T) {
// 	fmt.Println(credentials.AccessKeyID)
// 	fmt.Println(credentials.SecretAccessKey)
// }
func TestCreateLoadBalance(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateLoadBalancerArgs{
		Name: "golang-sdk-blb2",
	}
	blb, err := blbClient.CreateLoadBalancer(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("TestCreateLoadBalance", err.Error(), "nil"))

	} else {
		fmt.Println(blb.Address)
	}
}

func TestDescribeLoadBalancers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DescribeLoadBalancersArgs{
		LoadBalancerName: "blb-test",
	}
	list, err := blbClient.DescribeLoadBalancers(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}

	for _, blb := range list {
		fmt.Println(blb.PublicIp)
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	blbClient.SetDebug(true)
	args := &UpdateLoadBalancerArgs{
		LoadBalancerId: "lb-f9b682b2",
		Name:           "golang-name2",
	}
	err := blbClient.UpdateLoadBalancer(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteLoadBalancer(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DeleteLoadBalancerArgs{
		LoadBalancerId: "lb-f9b682b2",
	}
	err := blbClient.DeleteLoadBalancer(args)
	if err != nil {
		t.Error(err)
	}
}
