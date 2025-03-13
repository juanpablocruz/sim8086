package instrumentation

import "github.com/juanpablocruz/aware/pkg/instruction"

type ClockInterval struct {
	Min uint32
	Max uint32
}

type Timing struct {
	Base      ClockInterval
	Transfers uint32
	EAClocks  uint32
}

type TimingState struct {
	Assume8088               bool
	AssumeBranchTaken        bool
	AssumeAddressUnanaligned bool
	AssumeRepCount           uint32
	AssumeShiftCount         uint32
}

func (ts *TimingState) EstimateInstructionClocks(in instruction.Instruction) Timing {
	res := Timing{}

	return res
}

func (ts *TimingState) ExpectedClocksFrom(in instruction.Instruction, t Timing) ClockInterval {
	res := ClockInterval{}

	return res
}
