package info

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rai-project/machine"

	"github.com/mohae/joefriday/cpu/cpuinfo"
	"github.com/mohae/joefriday/mem/membasic"
	"github.com/mohae/joefriday/net/netdev"
	sysos "github.com/mohae/joefriday/system/os"
	"github.com/mohae/joefriday/system/version"
)

type info struct {
	machine.Machine
}

var (
	Info *machine.Machine
)

func New() (*machine.Machine, error) {
	info := &info{}
	err := info.get()
	if err != nil {
		return nil, err
	}
	return &info.Machine, nil
}

func (s *info) get() error {
	var err error
	s.Hostname, err = os.Hostname()
	if err != nil {
		return errors.Wrap(err, "hostname")
	}

	err = s.version()
	if err != nil {
		return err
	}

	err = s.os()
	if err != nil {
		return err
	}

	err = s.memory()
	if err != nil {
		return err
	}

	err = s.netdev()
	if err != nil {
		return err
	}

	err = s.cpus()
	if err != nil {
		return err
	}

	return nil
}

func (s *info) version() error {
	//Get Kernel info
	k, err := version.Get()
	if err != nil {
		return errors.Wrap(err, "kernel info")
	}
	s.KernelOS = k.OS
	s.KernelVersion = k.Version
	s.KernelArch = k.Arch
	s.KernelType = k.Type
	s.KernelCompileDate = k.CompileDate
	return nil
}

func (s *info) os() error {
	// Get release info
	o, err := sysos.Get()
	if err != nil {
		return errors.Wrap(err, "os release info")
	}
	s.OSName = o.Name
	s.OSID = o.ID
	s.OSIDLike = o.IDLike
	s.OSVersion = o.Version
	return nil
}

func (s *info) memory() error {
	m, err := membasic.Get()
	if err != nil {
		return errors.Wrap(err, "mem info")
	}
	s.MemTotal = m.MemTotal
	s.SwapTotal = m.SwapTotal
	return nil
}

func (s *info) netdev() error {
	// Get network devices
	inf, err := netdev.Get()
	if err != nil {
		return errors.Wrap(err, "network devices info")
	}
	s.NetDev = make([]string, len(inf.Device))
	for i := 0; i < len(inf.Device); i++ {
		s.NetDev[i] = inf.Device[i].Name
	}
	return nil
}

func (s *info) cpus() error {
	// Get processors
	cs, err := cpuinfo.Get()
	if err != nil {
		return errors.Wrap(err, "cpu info")
	}
	s.CPU = make([]*machine.CPU, len(cs.CPU))
	for i := 0; i < len(cs.CPU); i++ {
		var cpu machine.CPU
		cpu.CoreID = int32(cs.CPU[i].CoreID)
		cpu.Siblings = int32(cs.CPU[i].Siblings)
		cpu.VendorID = cs.CPU[i].VendorID
		cpu.CPUFamily = cs.CPU[i].CPUFamily
		cpu.Model = cs.CPU[i].Model
		cpu.ModelName = cs.CPU[i].ModelName
		cpu.Stepping = cs.CPU[i].Stepping
		cpu.Microcode = cs.CPU[i].Microcode
		cpu.CPUMHz = cs.CPU[i].CPUMHz
		cpu.BogoMIPS = cs.CPU[i].BogoMIPS
		cpu.CacheSize = cs.CPU[i].CacheSize
		cpu.CPUCores = int32(cs.CPU[i].CPUCores)
		cpu.Flags = make([]string, len(cs.CPU[i].Flags))
		copy(cpu.Flags, cs.CPU[i].Flags)
		cpu.Bugs = make([]string, len(cs.CPU[i].Bugs))
		copy(cpu.Bugs, cs.CPU[i].Bugs)
		s.CPU[i] = &cpu
	}
	s.Sockets = int32(cs.Sockets)
	return nil
}

func init() {
	if i, err := New(); err == nil {
		Info = i
	} else {
		Info = new(machine.Machine)
	}
}
