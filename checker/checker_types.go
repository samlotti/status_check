package checker

type RuleResult struct {
	Passed bool
	Name   string
	Extra  string
}

type IRule interface {
	GetTestName() string
	CheckRule() *RuleResult
}

type BaseCheck struct {
	Name string
}

func (b *BaseCheck) GetTestName() string {
	return b.Name
}
