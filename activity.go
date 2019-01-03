package data2tensorflogo

import (
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/harrydb/go/img/grayscale"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func init() {
	activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{})

// New - Creating the new activity
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{}

	return act, nil
}

// Activity - Creating the Activity object
type Activity struct {
	mapper mapper.Mapper
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// So I chose to use tf.NewTensor instead of something like this:
//    &tensorflow.Feature{&tensorflow.Feature_BytesList{&tensorflow.BytesList{values}}}
// because I feel it is more general.  THe Inference activity actually does the conversion to features for you.

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	// actionCtx := ctx.ActivityHost()

	input := &Input{}
	ctx.GetInputObject(input)

	var result *tf.Tensor

	switch v := input.Data.(type) {
	case int, []int, [][]int, [][][]int, [][][][]int, [][][][][]int:
		err = fmt.Errorf("%v needs to be converted to another format, may I suggest int32", v)
		return false, err

	case image.Image:
		// Converts an Image into a grayscale int array tensor that TF can then convert later

		src := input.Data.(image.Image)
		bounds := src.Bounds()
		w, h := bounds.Max.X, bounds.Max.Y
		// fmt.Println(w, h)
		imgsizeX := 256
		imgsizeY := 256
		src = imaging.Resize(src, imgsizeX, imgsizeY, imaging.Lanczos)

		src = grayscale.Convert(src, grayscale.ToGrayLuminance)

		//Converting Image to array
		var flatimg []int32
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				imageColor := src.At(x, y)
				rr, _, _, _ := imageColor.RGBA()
				gray := int32(rr) / 65535.
				flatimg = append(flatimg, gray)

			}
		}

		result, err = tf.NewTensor(flatimg)
		if err != nil {
			return false, err
		}
	default:
		result, err = tf.NewTensor(input.Data)
		if err != nil {
			return false, err
		}
	}

	output := &Output{Tensor: result}
	ctx.SetOutputObject(output)

	return true, nil
}
