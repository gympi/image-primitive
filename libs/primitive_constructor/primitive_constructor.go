package primitive_constructor

import (
	"log"
	"os/exec"

	"github.com/gympi/image-primitive/libs/string_utils"
)

// primitive console params
// https://github.com/fogleman/primitive
type PrimitiveApiParams struct {
	i string //input file
	o string //output file
	n int    //number of shapes
	m int    //mode: 0=combo, 1=triangle, 2=rect, 3=ellipse, 4=circle, 5=rotatedrect, 6=beziers, 7=rotatedellipse, 8=polygon
}

func PrimitiveConstructor(in string) string {
	binary, lookErr := exec.LookPath("./vender/primitive")
	if lookErr != nil {
		log.Println(binary)
		panic(lookErr)
	}

	primitive_filepath := "public/static/shared/primitive/" + string_utils.RandString(10) + ".svg"
	args := []string{"-i", in, "-o", primitive_filepath, "-n", "25"}

	cmd := "./vender/primitive"

	cmd_command := exec.Command(cmd, args...)
	cmd_command.Start()

	if err := cmd_command.Wait(); err != nil {
		log.Println(err)
	} else {
		log.Println("Importing complete")
	}

	return primitive_filepath
}
