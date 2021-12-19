package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	in := make(chan []string)
	go getInputFromBuffer(in)
	wg.Wait()
}

func getInputFromBuffer(ch chan []string) {
	rd := bufio.NewReader(os.Stdin)
	buffer := make(chan string)
	for {
		cmd, err := rd.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		cmd = strings.TrimSuffix(cmd, "\n")
		str := strings.Split(cmd, " | ")
		for i, v := range str {
			arr := strings.Split(v, " ")
			if i != 0 {
				arr = append(arr, <-buffer)
			}
			go handler(ch, buffer)
			ch <- arr
		}
		s := <-buffer
		fmt.Printf(s)
	}
}

func handler(in chan []string, out chan string) {
	for {
		arr := <-in
		switch arr[0] {
		case "cd":
			err := os.Chdir(arr[1])
			if err != nil {
				out <- err.Error()
			} else {
				out <- ""
			}
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				out <- err.Error()
			} else {
				out <- pwd + "\n"
			}
		case "echo":
			out <- arr[1] + "\n"
		case "ps":
			cmd := exec.Command("ps")
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
			out <- ""
		case "kill":
			i, _ := strconv.Atoi(arr[1])
			proc, err := os.FindProcess(i)
			proc.Kill()
			if err != nil {
				out <- err.Error()
			} else {
				out <- ""
			}
		case "nc":
			err := netcat(arr[1], arr[2])
			if err != nil {
				out <- err.Error()

			} else {
				out <- "ok"
			}
		default:
			out <- arr[0] + ": command not found\n"
		}
	}
}

func netcat(host, port string) error {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return err
	}
	tcp := conn.(*net.TCPConn)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		io.Copy(tcp, os.Stdin)
		fmt.Fprintln(os.Stderr, "conn 2 os.Stdin done")
		tcp.CloseWrite()
		wg.Done()
	}()
	go func() {
		io.Copy(os.Stdout, tcp)
		fmt.Fprintln(os.Stderr, "os.Stdout 2 conn done")
		tcp.CloseRead()
		wg.Done()
	}()
	wg.Wait()
	return err
}
