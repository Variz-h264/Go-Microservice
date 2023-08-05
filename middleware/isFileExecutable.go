package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
)

func FileExecutable(filename string) (bool, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	// ตรวจสอบสิทธิ์ execute ของไฟล์
	return fileInfo.Mode().Perm()&0111 != 0, nil
}

func Ready() {
	dir := "service"

	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Filename", "Status"})

	for _, file := range files {
		filename := filepath.Base(file)
		status := "Not ready to execute"
		isExecutable, err := FileExecutable(file)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if isExecutable {
			status = "Ready to execute"
		}

		table.Append([]string{filename, status})
	}

	table.Render()
}
