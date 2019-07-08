package addDimension

import (
	"fmt"

	"github.com/project-flogo/core/activity"
)

func init() {
	activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	d := input.Data
	fmt.Println("BLAH@@@@@@@", d)
	fmt.Printf("%T\n", d)
	var outd [][]float64
	var interm []float64
	for _, val := range d.([]interface{}) {
		interm = append(interm, val.(float64))
	}
	outd = append(outd, interm)

	output := &Output{Output: outd}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
