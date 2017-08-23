// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sfn_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSFN_CreateActivity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.CreateActivityInput{
		Name: aws.String("Name"), // Required
	}
	resp, err := svc.CreateActivity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_CreateStateMachine() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.CreateStateMachineInput{
		Definition: aws.String("Definition"), // Required
		Name:       aws.String("Name"),       // Required
		RoleArn:    aws.String("Arn"),        // Required
	}
	resp, err := svc.CreateStateMachine(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_DeleteActivity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.DeleteActivityInput{
		ActivityArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteActivity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_DeleteStateMachine() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.DeleteStateMachineInput{
		StateMachineArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteStateMachine(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_DescribeActivity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.DescribeActivityInput{
		ActivityArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeActivity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_DescribeExecution() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.DescribeExecutionInput{
		ExecutionArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeExecution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_DescribeStateMachine() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.DescribeStateMachineInput{
		StateMachineArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeStateMachine(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_GetActivityTask() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.GetActivityTaskInput{
		ActivityArn: aws.String("Arn"), // Required
		WorkerName:  aws.String("Name"),
	}
	resp, err := svc.GetActivityTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_GetExecutionHistory() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.GetExecutionHistoryInput{
		ExecutionArn: aws.String("Arn"), // Required
		MaxResults:   aws.Int64(1),
		NextToken:    aws.String("PageToken"),
		ReverseOrder: aws.Bool(true),
	}
	resp, err := svc.GetExecutionHistory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_ListActivities() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.ListActivitiesInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PageToken"),
	}
	resp, err := svc.ListActivities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_ListExecutions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.ListExecutionsInput{
		StateMachineArn: aws.String("Arn"), // Required
		MaxResults:      aws.Int64(1),
		NextToken:       aws.String("PageToken"),
		StatusFilter:    aws.String("ExecutionStatus"),
	}
	resp, err := svc.ListExecutions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_ListStateMachines() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.ListStateMachinesInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PageToken"),
	}
	resp, err := svc.ListStateMachines(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_SendTaskFailure() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.SendTaskFailureInput{
		TaskToken: aws.String("TaskToken"), // Required
		Cause:     aws.String("Cause"),
		Error:     aws.String("Error"),
	}
	resp, err := svc.SendTaskFailure(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_SendTaskHeartbeat() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.SendTaskHeartbeatInput{
		TaskToken: aws.String("TaskToken"), // Required
	}
	resp, err := svc.SendTaskHeartbeat(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_SendTaskSuccess() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.SendTaskSuccessInput{
		Output:    aws.String("Data"),      // Required
		TaskToken: aws.String("TaskToken"), // Required
	}
	resp, err := svc.SendTaskSuccess(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_StartExecution() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.StartExecutionInput{
		StateMachineArn: aws.String("Arn"), // Required
		Input:           aws.String("Data"),
		Name:            aws.String("Name"),
	}
	resp, err := svc.StartExecution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSFN_StopExecution() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sfn.New(sess)

	params := &sfn.StopExecutionInput{
		ExecutionArn: aws.String("Arn"), // Required
		Cause:        aws.String("Cause"),
		Error:        aws.String("Error"),
	}
	resp, err := svc.StopExecution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}