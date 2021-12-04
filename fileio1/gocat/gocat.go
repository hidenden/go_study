package gocat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Cat interface {
	read() error
}

func Read(c Cat) error {
	return c.read()
}

type cat struct {
	fileName string
	output   io.WriteCloser
}

type catReadLine struct {
	cat
}

type catReadBytes struct {
	cat
}

type catReadString struct {
	cat
}

type catScanner struct {
	cat
}

func NewCatReadLine(f string, o io.WriteCloser) Cat {
	c := catReadLine{}
	c.fileName = f
	c.output = o
	return &c
}

func NewCatReadBytes(f string, o io.WriteCloser) Cat {
	c := catReadBytes{}
	c.fileName = f
	c.output = o
	return &c
}

func NewCatReadString(f string, o io.WriteCloser) Cat {
	c := catReadString{}
	c.fileName = f
	c.output = o
	return &c
}

func NewCatScanner(f string, o io.WriteCloser) Cat {
	c := catScanner{}
	c.fileName = f
	c.output = o
	return &c
}

// 原始的なReadLineを使うパターン
func (c *catReadLine) read() error {
	defer c.output.Close()
	fp, err := os.Open(c.fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		line, prefix, err := reader.ReadLine()
		fmt.Fprint(c.output, string(line))
		if !prefix {
			fmt.Fprint(c.output, "\n")
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

// ReadBytes('\n')を使うパターン
func (c *catReadBytes) read() error {

	defer c.output.Close()
	fp, err := os.Open(c.fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadBytes('\n')
		fmt.Fprint(c.output, string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

// ReadString('\n')を使うパターン
func (c *catReadString) read() error {
	defer c.output.Close()
	fp, err := os.Open(c.fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		fmt.Fprint(c.output, line)
	}
	if err != io.EOF {
		return err
	}
	return nil
}

// Scannerを使うパターン
func (c *catScanner) read() error {
	defer c.output.Close()
	fp, err := os.Open(c.fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	for s.Scan() {
		line := s.Text()
		fmt.Fprintln(c.output, line)
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}
