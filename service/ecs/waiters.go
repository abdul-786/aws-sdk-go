// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ecs

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

var waiterServicesInactive *waiter.Config

func (c *ECS) WaitUntilServicesInactive(input *DescribeServicesInput) error {
	if waiterServicesInactive == nil {
		waiterServicesInactive = &waiter.Config{
			Operation:   "DescribeServices",
			Delay:       15,
			MaxAttempts: 40,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "failures[].reason",
					Expected: "MISSING",
				},
				{
					State:    "success",
					Matcher:  "pathAny",
					Argument: "services[].status",
					Expected: "INACTIVE",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterServicesInactive,
	}
	return w.Wait()
}

var waiterServicesStable *waiter.Config

func (c *ECS) WaitUntilServicesStable(input *DescribeServicesInput) error {
	if waiterServicesStable == nil {
		waiterServicesStable = &waiter.Config{
			Operation:   "DescribeServices",
			Delay:       15,
			MaxAttempts: 40,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "failures[].reason",
					Expected: "MISSING",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "services[].status",
					Expected: "DRAINING",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "services[].status",
					Expected: "INACTIVE",
				},
				{
					State:    "success",
					Matcher:  "path",
					Argument: "services | [@[?length(deployments)!=`1`], @[?desiredCount!=runningCount]][] | length(@) == `0`",
					Expected: true,
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterServicesStable,
	}
	return w.Wait()
}

var waiterTasksRunning *waiter.Config

func (c *ECS) WaitUntilTasksRunning(input *DescribeTasksInput) error {
	if waiterTasksRunning == nil {
		waiterTasksRunning = &waiter.Config{
			Operation:   "DescribeTasks",
			Delay:       6,
			MaxAttempts: 100,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "tasks[].lastStatus",
					Expected: "STOPPED",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "failures[].reason",
					Expected: "MISSING",
				},
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "tasks[].lastStatus",
					Expected: "RUNNING",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterTasksRunning,
	}
	return w.Wait()
}

var waiterTasksStopped *waiter.Config

func (c *ECS) WaitUntilTasksStopped(input *DescribeTasksInput) error {
	if waiterTasksStopped == nil {
		waiterTasksStopped = &waiter.Config{
			Operation:   "DescribeTasks",
			Delay:       6,
			MaxAttempts: 100,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "tasks[].lastStatus",
					Expected: "STOPPED",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterTasksStopped,
	}
	return w.Wait()
}
