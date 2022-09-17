package run

//go:generate mockgen -destination ../mock/run/firstRunner.go -package mock_run demo.golang.test/run FirstRunner
//go:generate mockgen -destination ../mock/run/secondRunner.go -package mock_run demo.golang.test/run SecondRunner
//go:generate mockgen -destination ../mock/run/thirdRunner.go -package mock_run demo.golang.test/run ThirdRunner

type FirstRunner interface {
	GetName() string
	FirstSpeedUp(speed int) int
}

type DefaultFirstRunner struct {
	Name   string
	Speed  int
	Runner SecondRunner
}

func (the *DefaultFirstRunner) GetName() string {
	return the.Name
}
func (the *DefaultFirstRunner) FirstSpeedUp(speed int) (result int) {
	the.Runner.GetName()
	the.Runner.GetName()
	result = the.Runner.SecondSpeedUp(speed) * the.Speed
	the.Runner.GetName()
	the.Runner.GetName()
	return
}

type SecondRunner interface {
	GetName() string
	SecondSpeedUp(speed int) int
}

type DefaultSecondRunner struct {
	Name   string
	Speed  int
	Runner ThirdRunner
}

func (the *DefaultSecondRunner) GetName() string {
	return the.Name
}
func (the *DefaultSecondRunner) SecondSpeedUp(speed int) (result int) {
	the.Runner.GetName()
	the.Runner.GetName()
	the.Runner.GetName()
	result = the.Runner.ThirdSpeedUp(speed) * the.Speed
	the.Runner.GetName()
	the.Runner.GetName()
	the.Runner.GetName()
	return
}

type ThirdRunner interface {
	GetName() string
	ThirdSpeedUp(speed int) int
}

type DefaultThirdRunner struct {
	Name  string
	Speed int
}

func (the *DefaultThirdRunner) GetName() string {
	return the.Name
}
func (the *DefaultThirdRunner) ThirdSpeedUp(speed int) (result int) {
	result = speed * the.Speed
	return
}
