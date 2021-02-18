package ini

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type config struct {
	filepath string
	conflist []map[string]map[string]string
}

func SetConfig(filepath string) *config {
	return &config{filepath: filepath}
}

func (c *config) GetValue(section, name string) string {
	c.ReadList()
	conf := c.ReadList()
	for _, v := range conf {
		for k, v := range v {
			if k == section {
				return v[name]
			}
		}
	}
	return ""
}

func (c *config) ReadList() []map[string]map[string]string {
	file, err := os.Open(c.filepath)
	if err != nil {
		log.Fatal("err1:", err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				log.Fatal("err2:", err)
			}
			if len(line) == 0 {
				break
			}
		}
		if len(line) == 0 {
			break
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#":
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1:])
			data[section][strings.TrimSpace(line[0:i])] = value
			if c.uniquappend(section) == true {
				c.conflist = append(c.conflist, data)
			}
		}
	}
	return c.conflist
}

func (c *config) uniquappend(section string) bool {
	for _, v := range c.conflist {
		for k := range v {
			if k == section {
				return false
			}
		}
	}
	return true
}
