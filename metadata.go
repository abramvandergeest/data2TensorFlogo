package data2tensorflogo

type Settings struct {
	Filler map[string]interface{} `md:"filler"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (o *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

func (o *Input) FromMap(values map[string]interface{}) error {
	o.Data = values["data"]
	return nil
}

type Output struct {
	Tensor interface{} `md:"tensor"`
}

func (r *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tensor": r.Tensor,
	}
}

func (r *Output) FromMap(values map[string]interface{}) error {

	r.Tensor, _ = values["data"]

	return nil
}
