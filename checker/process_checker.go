package checker

///
/// Checks to see if an executable is in the process list.
///

import (
	"fmt"
	"github.com/mitchellh/go-ps"
)

type ProcessCheck struct {
	BaseCheck
	Executable string
}

func (p *ProcessCheck) CheckRule() *RuleResult {
	processList, err := ps.Processes()
	if err != nil {
		return &RuleResult{
			Passed: false,
			Name:   p.Name,
			Extra:  fmt.Sprintf("error getting processes: %s", err),
		}
	}

	for _, process := range processList {
		if p.Executable == process.Executable() {
			return &RuleResult{
				Passed: true,
				Name:   p.Name,
				Extra:  fmt.Sprintf("pid: %d", process.Pid()),
			}
		}
	}

	return &RuleResult{
		Passed: false,
		Name:   p.Name,
		Extra:  "did not find the executable running",
	}
}

var pc IRule = &ProcessCheck{}

func ProcessParser(sections []string) IRule {
	return &ProcessCheck{
		BaseCheck:  BaseCheck{sections[1]},
		Executable: sections[2],
	}
}
