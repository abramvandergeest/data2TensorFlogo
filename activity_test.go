package data2tensorflogo

import (
	"fmt"
	"image"
	_ "image/jpeg" //I am decoding images here
	_ "image/png"  //I am decoding images here
	"os"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestImage(t *testing.T) {

	settings := &Settings{}
	iCtx := test.NewActivityInitContext(settings, nil)

	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	filename := "/Users/avanderg@tibco.com/datasets/box_images/Box/boxes/google-image(0001).jpeg"

	// Opening the image
	imgfile, err := os.Open(filename)
	if err != nil {
		iCtx.Logger().Info("Unable to download item %q, %v", filename, err)
	}
	defer imgfile.Close()

	src, _, err := image.Decode(imgfile)
	if err != nil {
		iCtx.Logger().Info("Error Decoding file: %v", err)
	}

	tc.SetInput("data", src)

	_, err = act.Eval(tc)
	if err != nil {
		fmt.Println(err)
	}

	output := tc.GetOutput("tensor")

	ten, err2 := tf.NewTensor([]int32{2, 3, 4})
	if err2 != nil {
		fmt.Println(err2)
	}
	assert.IsType(t, output, ten)
	assert.Nil(t, err)
}

func TestData(t *testing.T) {

	settings := &Settings{}
	iCtx := test.NewActivityInitContext(settings, nil)

	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	src := [][]float32{{1}, {2}, {3}, {4}}

	tc.SetInput("data", src)

	_, err = act.Eval(tc)
	if err != nil {
		fmt.Println(err)
	}

	output := tc.GetOutput("tensor")

	ten, err2 := tf.NewTensor([]int32{2, 3, 4})
	if err2 != nil {
		fmt.Println(err2)
	}
	assert.IsType(t, output, ten)
	assert.Nil(t, err)
}

func newActivityHost() *test.TestActivityHost {
	input := map[string]data.TypedValue{"Input1": data.NewTypedValue(data.TypeString, "")}
	output := map[string]data.TypedValue{"Output1": data.NewTypedValue(data.TypeString, ""), "Output2": data.NewTypedValue(data.TypeInt, "")}

	ac := &test.TestActivityHost{
		HostId:     "1",
		HostRef:    "github.com/TIBCOSoftware/flogo-contrib/action/flow",
		IoMetadata: &metadata.IOMetadata{Input: input, Output: output},
		HostData:   data.NewSimpleScope(nil, nil),
	}

	return ac
}
