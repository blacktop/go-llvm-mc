package main

import (
	"bufio"
	"context"
	"debug/macho"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func Dump(r *bufio.Reader) string {

	var err error

	if r == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(r.Size() * 4)

	leBuff := make([]byte, 4)

	for {
		err = binary.Read(r, binary.LittleEndian, &leBuff)
		if err == io.EOF {
			break
		}

		buf.WriteString(fmt.Sprintf("%#02x %#02x %#02x %#02x ", leBuff[0], leBuff[1], leBuff[2], leBuff[3]))
	}

	return buf.String()
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "timeout",
			Value:  45,
			Usage:  "timeout (in seconds)",
			EnvVar: "LLVMMC_TIMEOUT",
		},
		cli.IntFlag{
			Name:  "count, c",
			Value: 500,
			Usage: "number of instructions to print",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "disassemble",
			Aliases: []string{"d"},
			Usage:   "disassemble binary",
			Action: func(c *cli.Context) error {

				if c.Args().Present() {
					path, err := filepath.Abs(c.Args().First())
					if err != nil {
						return err
					}

					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.GlobalInt("timeout"))*time.Second)
					defer cancel()

					// disass := exec.CommandContext(ctx, "llvm-mc", "-disassemble", "-arch=arm64", "-mattr=v8.3a", "-show-encoding")
					disass := exec.CommandContext(ctx, "llvm-mc", "-disassemble", "-arch=arm64", "-mattr=v8.5a")
					stdin, err := disass.StdinPipe()
					if err != nil {
						return err
					}

					f, err := macho.Open(path)
					if err != nil {
						return err
					}

					for _, sec := range f.Sections {
						// fmt.Println(sec.Name, sec.Seg)
						if strings.EqualFold(sec.Name, "__text") && strings.EqualFold(sec.Seg, "__TEXT_EXEC") {
							go func() {
								instCount := c.GlobalInt("count")
								defer stdin.Close()
								r := bufio.NewReader(sec.Open())
								leBuff := make([]byte, 4)
								for {
									err = binary.Read(r, binary.LittleEndian, &leBuff)
									if err == io.EOF {
										break
									}
									instr := fmt.Sprintf("%#02x %#02x %#02x %#02x ", leBuff[0], leBuff[1], leBuff[2], leBuff[3])
									io.WriteString(stdin, instr)
									instCount--
									if instCount < 0 {
										break
									}
								}
							}()
							break
						}
					}

					disass.Wait()

					output, err := disass.Output()
					// output, err := disass.CombinedOutput()
					if err != nil {
						return errors.Wrap(err, string(output))
					}

					// check for exec context timeout
					if ctx != nil {
						if ctx.Err() == context.DeadlineExceeded {
							return fmt.Errorf("disassemble timed out")
						}
					}

					fmt.Println(string(output))
				} else {
					log.Fatal(fmt.Errorf("please supply a mach-o to disassemble"))
				}
				return nil
			},
		},
		{
			Name:    "assemble",
			Aliases: []string{"a"},
			Usage:   "convert ASM to binary",
			Action: func(c *cli.Context) error {
				fmt.Println("not implimented yet")
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
