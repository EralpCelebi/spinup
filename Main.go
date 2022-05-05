/*
		DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
				Version 2, December 2004

Copyright 2022 Eralp Çelebi

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

		DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

0. You just DO WHAT THE FUCK YOU WANT TO.

Eralp Çelebi
*/

package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type Configuration struct {
	Image 	*string;
	Files		*string;
	Host		*string;
	Target	*string;
	Command *string;
	Mount		*bool;
	Folder	*bool;
}

func Matrix(Status Configuration) {
	if os.Getpid() == 1 {
		Enter(Status);	
	} else {
		Setup(Status);
	}
}

func Setup(Status Configuration) {
	if !*Status.Folder {
		Try(os.Mkdir(*Status.Target, 0777));
		
		Hack := exec.Command("mount", "-t", *Status.Files, *Status.Image, *Status.Target);
		Check(Hack.Run());
	} else  {
		_, Error := os.Stat(*Status.Target);
		Check(Error);
	} 
	
	if *Status.Mount {
		Check(syscall.Mount("/sys", filepath.Join(*Status.Target, "sys"),  "", syscall.MS_BIND, ""));
		Check(syscall.Mount("/dev", filepath.Join(*Status.Target, "dev"),  "", syscall.MS_BIND, ""));
	}

	Executable, _ := os.Executable()

	Command := exec.Command(Executable, os.Args[1:]...);
	
	Command.Stdin 	= os.Stdin;
	Command.Stdout 	= os.Stdout;
	Command.Stderr 	= os.Stderr;
	
	Command.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
								syscall.CLONE_NEWPID |
								syscall.CLONE_NEWIPC |
								syscall.CLONE_NEWNS,
	};

	Command.Run();
	
	if *Status.Mount {
		Try(syscall.Unmount(filepath.Join(*Status.Target, "sys"), 0));
		Try(syscall.Unmount(filepath.Join(*Status.Target, "dev"), 0));
	}
	
	Try(syscall.Unmount(*Status.Target, 0));
	Try(os.Remove(*Status.Target));
}

func Enter(Status Configuration) {
	Check(syscall.Chroot(*Status.Target));
	Check(syscall.Chdir("/"));

	Check(syscall.Mount("proc", "proc", "proc", 0, ""));

	Command := exec.Command(*Status.Command);

	Command.Stdin 	= os.Stdin;
	Command.Stdout 	= os.Stdout;
	Command.Stderr 	= os.Stderr;

	Check(Command.Run());
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1);
	}

	Status := Configuration {
		Image: 		flag.String("image", "root.img", "Specifies the target image."),
		Files:		flag.String("files", "ext4", "Specifies the filesystem the image uses."),
		Host: 		flag.String("host", "spinup", "Specifies the hostname to be used."),
		Target: 	flag.String("target", "/tmp/spinup", "Specifies the work directory."),
		Command:	flag.String("command", "/bin/bash", "Specifies the command to run."),
		Mount: 		flag.Bool("mount", false, "Mounts system directories like /dev, /sys"),
		Folder: 	flag.Bool("folder", false, "Use the target folder as the rootfs."),
	};

	flag.Parse();

	Matrix(Status);
}
