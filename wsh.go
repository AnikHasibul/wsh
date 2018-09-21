/*
* This program is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 2 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program; if not, see <http://www.gnu.org/licenses/>.
*
* Copyright (C) Anik Hasibul (@AnikHasibul), 2018
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
)

func init() {
	fmt.Println("")
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Getting default shell
	shellLink := usr.HomeDir + "/.wshell"
	if _, err := os.Stat(shellLink); os.IsNotExist(err) {
		fmt.Print("Setting up the working shell!\nEnter the new value, or press ENTER for the default\n\tLogin Shell [bash]: ")
		var inp = "bash"
		fmt.Scanf("%s", &inp)
		paths := strings.Split(os.Getenv("PATH"), ":")
		var exist bool
		for _, v := range paths {

			if _, err := os.Stat(v + "/" + inp); err == nil {
				exist = true
			}
			if exist == true {
				err := os.Symlink(v+"/"+inp, shellLink)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return

			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	// creating output pipeline
	pr, pw := io.Pipe()
	// getting current user directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Getting default shell
	shellpath := usr.HomeDir + "/.wshell"
	// Setting up exec options
	cmd := exec.Command(shellpath, "-l")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = pw
	wg.Add(1)
	// Running cmd!
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

	}(&wg)

	// piping stderr
	go streamStderr(pr)
	// Wait for getting shell done!
	wg.Wait()
}

func streamStderr(reader io.Reader) {
	// reading from reader
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		warning := scanner.Text()
		// ignore linker error
		if !strings.Contains(
			warning,
			"WARNING: linker:",
		) {
			fmt.Fprintln(os.Stderr, warning)
		}
	}
}
