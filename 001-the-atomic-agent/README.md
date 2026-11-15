# 001: The atomic agent

This is the first step in the repository's journey to understand agent systems
from first principles. It keeps the core loop visible:

1. Give the agent an objective.
2. Ask a model what to do next.
3. Execute the requested tool.
4. Feed the result back as an observation.
5. Repeat until the model returns a final answer or the run reaches its limit.

Here, `FakeModel` stands in for a language model, and the two Kubernetes-style
tools return fixed values. This keeps attention on the control flow before
adding a real model, external services, or more complex orchestration.

## Run

Requires Go 1.20 or newer. From the repository root:

```bash
go run ./001-the-atomic-agent
```

The example asks why the payments are failing. The agent checks a pod's status,
inspects its logs, and reports what it observed.
