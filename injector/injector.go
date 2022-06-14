package injector

import (
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var (
	inject_new = []byte("electrom")
	inject_old = []byte("electron")
)

//go:embed inject
var injectFs embed.FS

func Inject(prog string, inject_name string) error {
	data, err := injectFs.ReadFile(path.Join("inject", inject_name+".js"))
	if err == nil {
		return inject(prog, data)
	} else {
		data, err := ioutil.ReadFile(inject_name)
		if err != nil {
			return err
		}
		return inject(prog, data)
	}
}

func inject(prog string, inject_content []byte) error {
	if _, err := os.Stat(prog + ".bak"); err == nil {
		fmt.Println("Backup file already exists, skipping backup")
	} else {
		err := os.Rename(prog, prog+".bak")
		if err != nil {
			return err
		}
	}
	prog_content, err := ioutil.ReadFile(prog + ".bak")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", prog)
		}
		return err
	}
	inject_count := bytes.Count(prog_content, inject_old)
	if inject_count != 1 {
		return fmt.Errorf("%s does not contain valid injection point", prog)
	}

	inject_path := path.Join(path.Dir(prog), "..", "node_modules", string(inject_new)+".js")
	inject_content = []byte(fmt.Sprintf("%s;\nmodule.exports = require(\"electron\");", inject_content))
	err = ioutil.WriteFile(inject_path, inject_content, 0644)
	if err != nil {
		return err
	}

	inject_payload := inject_new

	prog_content = []byte(bytes.Replace(prog_content, inject_old, inject_payload, 1))

	err = ioutil.WriteFile(prog, prog_content, 0644)
	if err != nil {
		return err
	}
	return nil
}
