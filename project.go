package tyrellcorp

import (
	"bytes"
	fmt "fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// project directories
	cmdDir    = "cmd"
	configDir = "config"
	protoDir  = "proto"
	grpcDir   = "grpc"
	restDir   = "rest"

	// project assets
	versionFile  = ".version"
	makefileFile = "Makefile"
	licenseFile  = "LICENSE"
	readmeFile   = "README.md"

	recordProtoFile  = "record.proto"
	serviceProtoFile = "service.proto"

	mainGoFile        = "main.go"
	serverGoFile      = "server.go"
	serviceGoFile     = "service.go"
	serviceTestGoFile = "service_test.go"
	compositionGoFile = "composition.go"
	errorGoFile       = "error.go"

	// project templates
	versionTpl  = "templates/project/version.tpl"
	makefileTpl = "templates/project/makefile.tpl"
	licenseTpl  = "templates/project/license.tpl"
	readmeTpl   = "templates/project/readme.tpl"

	modelProtoTpl   = "templates/proto/model.proto.tpl"
	recordProtoTpl  = "templates/proto/record.proto.tpl"
	serviceProtoTpl = "templates/proto/service.proto.tpl"

	grpcServerTpl  = "templates/go/grpc_server.go.tpl"
	restServerTpl  = "templates/go/rest_server.go.tpl"
	compositionTpl = "templates/go/config_composition.go.tpl"
	mainTpl        = "templates/go/cmd_main.go.tpl"
	modelCRUDTpl   = "templates/go/model_crud.go.tpl"
	modelTestTpl   = "templates/go/model_test.go.tpl"
	serviceTpl     = "templates/go/service.go.tpl"
	serviceTestTpl = "templates/go/service_test.go.tpl"
	errorTpl       = "templates/go/error.go.tpl"
)

func getDir(repo, pkg string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", strings.ToLower(repo), strings.ToLower(pkg))
}

func CreateProject(spec *Spec) error {
	dir := getDir(spec.Repository, spec.Package)

	// generate main project directory
	log.Debugf("generating working directory %s", dir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrapf(err, "unable to generate project directory %s", dir)
	}

	// generate sub directories
	for _, d := range []string{cmdDir, configDir, protoDir, grpcDir, restDir} {
		child := filepath.Join(dir, d)
		log.Debugf("generating child director %s", child)
		if err := os.MkdirAll(child, 0755); err != nil {
			return errors.Wrapf(err, "unable to generate project subdirectory %s", child)
		}
	}

	return nil
}

func parseTpl(f string, d interface{}) ([]byte, error) {
	if _, err := os.Stat(f); err != nil {
		return nil, errors.Wrapf(err, "required template file %s not found", f)
	}

	fm := template.FuncMap{
		"ToLower": strings.ToLower,
		"ToUpper": strings.ToUpper,
		"Trim":    strings.TrimSpace,
	}

	tpl, err := template.New(filepath.Base(f)).Funcs(fm).ParseFiles(f)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse required template file %s", f)
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	if err := tpl.Execute(buf, d); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GenerateAssets(spec *Spec) error {
	dir := getDir(spec.Repository, spec.Package)

	assets := []struct {
		tpl  string
		data interface{}
		path string
	}{
		{
			makefileTpl,
			spec,
			filepath.Join(dir, makefileFile),
		},
		{
			versionTpl,
			struct{ Version string }{"1.0.0"},
			filepath.Join(dir, versionFile),
		},
		{
			recordProtoTpl,
			spec,
			filepath.Join(dir, protoDir, recordProtoFile),
		},
		{
			modelProtoTpl,
			spec,
			filepath.Join(dir, protoDir, fmt.Sprintf("%s.proto", strings.ToLower(spec.Name))),
		},
		{
			serviceProtoTpl,
			spec,
			filepath.Join(dir, protoDir, fmt.Sprintf("%s%s", strings.ToLower(spec.Name), serviceProtoFile)),
		},
		{
			grpcServerTpl,
			spec,
			filepath.Join(dir, grpcDir, serverGoFile),
		},
		{
			restServerTpl,
			spec,
			filepath.Join(dir, restDir, serverGoFile),
		},
		{
			compositionTpl,
			spec,
			filepath.Join(dir, configDir, compositionGoFile),
		},

		{
			mainTpl,
			spec,
			filepath.Join(dir, cmdDir, mainGoFile),
		},
		{
			modelCRUDTpl,
			spec,
			filepath.Join(dir, fmt.Sprintf("%s.go", strings.ToLower(spec.Name))),
		},
		{
			modelTestTpl,
			spec,
			filepath.Join(dir, fmt.Sprintf("%s_test.go", strings.ToLower(spec.Name))),
		},
		{
			serviceTpl,
			spec,
			filepath.Join(dir, fmt.Sprintf("%s%s", strings.ToLower(spec.Name), serviceGoFile)),
		},
		{
			serviceTestTpl,
			spec,
			filepath.Join(dir, serviceTestGoFile),
		},
		{
			errorTpl,
			spec,
			filepath.Join(dir, errorGoFile),
		},
	}

	for _, asset := range assets {
		tpl, err := parseTpl(asset.tpl, asset.data)
		if err != nil {
			return errors.Wrapf(err, "unable to parse template file %s", asset.tpl)
		}

		if err := ioutil.WriteFile(asset.path, tpl, 0644); err != nil {
			return errors.Wrapf(err, "unable to write %s to disk", asset.tpl)
		}
	}

	log.Debugf("running protoc for %s", spec.Name)

	// need to run the makefile in order to generate the additional go code from the .proto files
	cmd := exec.Command("make", "proto")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "unable to process proto files for %s", spec.Name)
	}

	return nil
}
