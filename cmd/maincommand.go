package cmd

import (
	"errors"
	"fmt"
	"go/types"
	"io/ioutil"
	"os"

	"github.com/ptcar2009/avro-generator/pkg/avrotypes"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
)

type flags struct {
	outputFile *string
	pkg        *string
}

var f = flags{
	outputFile: new(string),
	pkg:        new(string),
}

var MainCommand = &cobra.Command{
	Use:  "avro-generator",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if *f.outputFile == "." {
			*f.outputFile = fmt.Sprintf("%s.avsc", args[0])
		}
		sourceTypePackage := *f.pkg
		sourceTypeName := args[0]

		// 2. Inspect package and use type checker to infer imported types
		pkg, err := loadPackage(sourceTypePackage)
		if err != nil {
			return err
		}

		// 3. Lookup the given source type name in the package declarations
		obj := pkg.Types.Scope().Lookup(sourceTypeName)
		if obj == nil {
			return fmt.Errorf("%s not found in declared types of %s",
				sourceTypeName, pkg)
		}

		// 4. We check if it is a declared type
		if _, ok := obj.(*types.TypeName); !ok {
			return fmt.Errorf("%v is not a named type", obj)
		}

		err = ioutil.WriteFile(*f.outputFile, []byte(avrotypes.ASTNodeToAvro(obj.Name(), obj.Type().Underlying())), os.ModePerm)
		return err
	},
}

func loadPackage(path string) (*packages.Package, error) {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return nil, fmt.Errorf("loading packages for inspection: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		return nil, errors.New("error in reading package")
	}

	return pkgs[0], nil
}

func init() {

	f.outputFile = MainCommand.Flags().StringP(
		"output",
		"o",
		".",
		"set the output file",
	)
	f.pkg = MainCommand.Flags().StringP(
		"package",
		"p",
		"",
		"set the package",
	)
}
