package internal

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Panicf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
func Execute(format string, args ...any) error {
	cmd := fmt.Sprintf(format, args...)
	println(cmd)
	executor := exec.Command("sh", "-c", cmd)

	// 创建一个管道用于读取标准输出
	stdoutPipe, err := executor.StdoutPipe()
	if err != nil {
		Panicf("Failed to create stdout pipe: %v", err)
	}

	// 创建一个管道用于读取标准错误
	stderrPipe, err := executor.StderrPipe()
	if err != nil {
		Panicf("Failed to create stderr pipe: %v", err)
	}

	// 启动命令
	if err := executor.Start(); err != nil {
		Panicf("Failed to start command: %v", err)
	}

	// 读取标准输出
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading stdout: %v", err)
		}
	}()

	// 读取标准错误
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading stderr: %v", err)
		}
	}()

	// 等待命令执行完成
	return executor.Wait()
}
