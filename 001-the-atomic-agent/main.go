package main

import (
	"fmt"
)

type Decision struct {
	Tool  string
	Final string
}
type Model interface {
	Decide(objective string, observations []string) Decision
}

type FakeModel struct{}

func (m FakeModel) Decide(objective string, observations []string) Decision {
	if len(observations) == 0 {
		return Decision{Tool: "get_pod_status"}
	} else if len(observations) == 1 {
		return Decision{Tool: "inspect_pod_logs"}
	}

	return Decision{
		Final: "The payment pod is unhealthy.\nStatus: " +
			observations[0] + "\nLogs: " + observations[1],
	}
}

func getPodStatus() string {
	return "CrashLoopBackOff"
}

func inspectPodLogs() string {
	return "Error: panic with an invalid memory address"
}

func executeTool(name string) (string, error) {
	switch name {
	case "get_pod_status":
		return getPodStatus(), nil
	case "inspect_pod_logs":
		return inspectPodLogs(), nil
	default:
		return "", fmt.Errorf("unknown tool: %s", name)
	}
}

func runAgent(objective string, model Model) {
	var observations []string

	const maxIterations = 5

	for i := 0; i < maxIterations; i++ {
		decision := model.Decide(objective, observations)
		if decision.Final != "" {
			fmt.Println("Agent:", decision.Final)
			return
		}

		result, err := executeTool(decision.Tool)
		if err != nil {
			fmt.Println("Execution error:", err)
			return
		}

		observations = append(observations, result)
	}
	fmt.Println("Agent stopped: iteration limit reached")
}

func main() {
	runAgent(
		"Find why the payments are failing",
		FakeModel{},
	)
}
