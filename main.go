package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var name string
	app := cli.NewApp()
	app.Name = "echo-cli"
	app.Version = "0.0.1"
	app.Usage = "This echo whatever you provide to the app as args"
	app.ArgsUsage = "echo app [arg 1] [args2] ..."
	app.Authors = []cli.Author{
		{
			Name:  "Richard",
			Email: "ricdon41@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "Input name",
			Hidden:      false,
			Value:       "Richard",
			Destination: &name,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", name)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "details",
			Usage: "echo the full details of the user",
			Action: func(c *cli.Context) error {
				fmt.Printf("Echoing the details")
				return nil
			},
		},
		{
			Name:  "write",
			Usage: "write user information to the file",
			Action: func(c *cli.Context) error {
				file := get_or_create_file(".", "user_name.txt")
				defer file.Close()
				var data string
				if len(c.Args().Get(0)) != 0 {
					data = c.Args().Get(0)
					if write_file(file, data) {
						fmt.Println("The write was successful")
					}
				} else {
					fmt.Println("No data to write")
				}
				return nil
			},
		},
		{
			Name:  "read",
			Usage: "write user information to the file",
			Action: func(c *cli.Context) error {
				file := get_or_create_file(".", "user_name.txt")

				if print_file(file) {
					fmt.Println("The read was successful")
				}
				file.Close()
				return nil
			},
		},
	}

	app.Run(os.Args)
}

// get file or create file
func get_or_create_file(path string, name string) *os.File {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", path, name), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// write to a file
func write_file(file *os.File, data string) bool {
	_, err := file.WriteString(fmt.Sprintf("%s\n", data))
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func print_file(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return true
}
