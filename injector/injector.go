package injector

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	inject_new = "function validateString(){};"
	inject_old = regexp.MustCompile(`function validateString[^}]+}`)
)

//go:embed inject
var injectFs embed.FS

func Inject(prog string, inject_name string) error {
	data, err := injectFs.ReadFile(path.Join("inject", inject_name+".js"))
	if err == nil {
		return inject(prog, inject_name+".js", data)
	} else {
		fmt.Printf("%s\n", err)
		data, err := ioutil.ReadFile(inject_name)
		if err != nil {
			return err
		}
		return inject(prog, inject_name, data)
	}
}

func inject(prog string, inject_name string, inject_content []byte) error {
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
	inject_point := inject_old.Find(prog_content)
	if inject_point == nil {
		return fmt.Errorf("%s does not contain valid injection point", prog)
	}

	inject_path := path.Join(path.Dir(prog), "..", "node_modules", path.Base(inject_name))
	err = ioutil.WriteFile(inject_path, inject_content, 0644)
	if err != nil {
		return err
	}

	inject_payload := fmt.Sprintf("%s;mod.require('%s')", inject_new, path.Base(inject_name))

	if len(inject_payload) > len(inject_point) {
		return fmt.Errorf("Inject payload is too long")
	}

	inject_payload = strings.Repeat(" ", len(inject_point)-len(inject_payload)) + inject_payload

	prog_content = []byte(bytes.Replace(prog_content, inject_point, []byte(inject_payload), 1))

	err = ioutil.WriteFile(prog, prog_content, 0644)
	if err != nil {
		return err
	}
	return nil
}
