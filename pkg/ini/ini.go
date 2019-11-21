// 读取ini配置
package ini

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ini struct {
	conf map[string]map[string]string
}

// 返回String的配置值
func (c *Ini) String(section, key string) string {
	key = strings.Join(strings.Split(key, "."), "_")
	value := c.conf[section][key]
	return value
}

// 返回Int的配置值
func (c *Ini) Int(section, key string) (int, error) {
	return strconv.Atoi(c.String(section, key))
}

// 返回Int64配置值
func (c *Ini) Int64(section, key string) (int64, error) {
	return strconv.ParseInt(c.String(section, key), 10, 64)
}

// 返回Float64配置值
func (c *Ini) Float64(section, key string) (float64, error) {
	return strconv.ParseFloat(c.String(section, key), 64)
}

// 返回bool配置值
func (c *Ini) Bool(section, key string) (bool, error) {
	return strconv.ParseBool(c.String(section, key))
}

// 初始化一个ini句柄
func NewIni(filePath string) (*Ini, error) {

	cf := &Ini{
		conf: make(map[string]map[string]string, 5),
	}

	f, err := newFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	section := ""
	buf := bufio.NewReader(f)

	for {
		lineInfo, err := buf.ReadString('\n')
		if err != nil && err != errors.New("EOF") && lineInfo == "" {
			break
		}

		lineInfo = strings.TrimSpace(lineInfo)

		if lineInfo == "" {
			continue
		}

		if idx := strings.Index(lineInfo, "["); idx != -1 {
			if lineInfo[len(lineInfo)-1:] != "]" {
				return nil, errors.New("Error:failed to parse this section:\"" + lineInfo + "\"")
			}
			section = lineInfo[1 : len(lineInfo)-1]
			cf.conf[section] = make(map[string]string)
		} else {
			replacer := strings.NewReplacer(" ", "")
			lineInfo = replacer.Replace(lineInfo)
			spl := strings.Split(lineInfo, "=")

			// 注释的内容
			if lineInfo[0:1] == ";" {
				continue
			}

			if len(spl) < 2 {
				return nil, errors.New("Error:failed to parse key value:\"" + lineInfo + "\"")
			}
			key := strings.Replace(spl[0], ".", "_", -1)
			cf.conf[section][key] = spl[1]
		}
	}

	return cf, nil
}

// 打开一个文件句柄
func newFile(filePath string) (*os.File, error) {
	exist, err := pathExists(filePath)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("Error:File not exists:" + filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// 检查文件或文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
