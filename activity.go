package data2tensorflogo

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func init() {
	activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{})

func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{}

	fmt.Println("GETS HERE !!!!!!!!!!!!!!!!!!!!!!!")
	ctx.Logger().Debugf("Mappings: %+v", s.Mappings)
	fmt.Printf("Mappings: %+v\n", s.Mappings)

	// fmt.Println("!!!!!GETS HERE !!!!!!!!!!!!!!!!!!!!!!!")
	// act.mapper, err = ctx.MapperFactory().NewMapper(s.Mappings)
	// if err != nil {
	// 	return nil, err
	// }

	return act, nil
}

// Activity is an Activity that is used to reply/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type Activity struct {
	mapper mapper.Mapper
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	fmt.Println("GO GO GADGET CODE")
	actionCtx := ctx.ActivityHost()

	// if a.mapper == nil {
	// 	//No mapping
	// 	return true, nil
	// }

	inputScope := actionCtx.Scope() //host data

	input := &Input{}
	ctx.GetInputObject(input)

	fmt.Println(inputScope)
	fmt.Println("input", input)

	// // Pre-Process image
	src := input.Data //.(image.Image)
	// bounds := src.Bounds()
	// w, h := bounds.Max.X, bounds.Max.Y
	// fmt.Println(w, h)
	// // fmt.Println(img.Height)
	// imgsizeX := 256
	// imgsizeY := 256
	// src = imaging.Resize(src, imgsizeX, imgsizeY, imaging.Lanczos)
	// // src = imaging.Grayscale(src)
	// // src = grayscale.Convert(src, grayscale.ToGrayLuminance)

	// //Converting Image to array
	// var flatimg []float32
	// for x := 0; x < imgsizeX; x++ {
	// 	for y := 0; y < imgsizeY; y++ {
	// 		imageColor := src.At(x, y)
	// 		// fmt.Println(imageColor)
	// 		rr, _, _, _ := imageColor.RGBA()
	// 		// if rr != 65535 {
	// 		//fmt.Println(rr, bb, gg)
	// 		// }
	// 		gray := float32(rr) / 65535.
	// 		flatimg = append(flatimg, gray)

	// 	}
	// }

	fmt.Println("gets to the end")

	// Array to TF tensor
	flatimgout, err := tf.NewTensor(src)
	if err != nil {
		return false, err
	}

	output := &Output{Tensor: flatimgout}

	ctx.SetOutputObject(output)

	actionCtx.Scope().SetValue("output", flatimgout)

	return true, nil
}
