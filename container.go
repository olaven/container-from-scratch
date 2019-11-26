/*
	Although this code is techicaly written by me (@olaven), 
	it is _heavily_ inspired by the following talk
	by Liz Rice: https://www.youtube.com/watch?v=Utf-A4rODH8
*/

package main 
import (
	"os" 
	"os/exec"
	"syscall"
)

func main() {

	if (len(os.Args) < 3) {

		panic("Invalid argument count."); 
	}

	switch (os.Args[1]) {
		case "run": 
			run() 
		case "child": 
			child()	
	default: 
		panic("Invalid argument.")
	}
}

func run() {

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout 
	cmd.Stderr = os.Stderr 
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout 
	cmd.Stderr = os.Stderr 

	must(syscall.Chroot("/home/rootfs"))
	must(syscall.Chdir("/"))

	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run()) 
}


func must(err error) {

	if err != nil {
		panic(err)
	}
} 
