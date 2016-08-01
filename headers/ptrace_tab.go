// Location: include/uapi/linux/ptrace.h

package headers

var PtraceLookup = map[int]string{
	0:      "PTRACE_TRACEME",
	1:      "PTRACE_PEEKTEXT",
	2:      "PTRACE_PEEKDATA",
	3:      "PTRACE_PEEKUSER",
	4:      "PTRACE_POKETEXT",
	5:      "PTRACE_POKEDATA",
	6:      "PTRACE_POKEUSER",
	7:      "PTRACE_CONT",
	8:      "PTRACE_KILL",
	9:      "PTRACE_SINGLESTEP",
	12:     "PTRACE_GETREGS",
	13:     "PTRACE_SETREGS",
	14:     "PTRACE_GETFPREGS",
	15:     "PTRACE_SETFPREGS",
	16:     "PTRACE_ATTACH",
	17:     "PTRACE_DETACH",
	18:     "PTRACE_GETFPXREGS",
	19:     "PTRACE_SETFPXREGS",
	24:     "PTRACE_SYSCALL",
	0x4200: "PTRACE_SETOPTIONS",
	0x4201: "PTRACE_GETEVENTMSG",
	0x4202: "PTRACE_GETSIGINFO",
	0x4203: "PTRACE_SETSIGINFO",
	0x4204: "PTRACE_GETREGSET",
	0x4205: "PTRACE_SETREGSET",
	0x4206: "PTRACE_SEIZE",
	0x4207: "PTRACE_INTERRUPT",
	0x4208: "PTRACE_LISTEN",
	0x4209: "PTRACE_PEEKSIGINFO",
	0x420a: "PTRACE_GETSIGMASK",
	0x420b: "PTRACE_SETSIGMASK",
}