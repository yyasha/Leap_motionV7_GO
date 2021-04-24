package leap

type Frame struct {
	CurrentFrameRate float64       `json:"currentFrameRate"`
	Devices          []interface{} `json:"devices"`
	Hands            []struct {
		ArmBasis      [][]float64 `json:"armBasis"`
		ArmWidth      float64     `json:"armWidth"`
		Confidence    float32         `json:"confidence"`
		Direction     []float64   `json:"direction"`
		Elbow         []float64   `json:"elbow"`
		GrabAngle     float64     `json:"grabAngle"`
		GrabStrength  float32         `json:"grabStrength"`
		ID            int         `json:"id"`
		PalmNormal    []float64   `json:"palmNormal"`
		PalmPosition  []float64   `json:"palmPosition"`
		PalmVelocity  []float64   `json:"palmVelocity"`
		PalmWidth     float64     `json:"palmWidth"`
		PinchDistance float64     `json:"pinchDistance"`
		PinchStrength float32         `json:"pinchStrength"`
		TimeVisible   float64     `json:"timeVisible"`
		Type          string      `json:"type"`
		Wrist         []float64   `json:"wrist"`
	} `json:"hands"`
	ID         int `json:"id"`
	Pointables []struct {
		Bases        [][][]float64 `json:"bases"`
		BtipPosition []float64     `json:"btipPosition"`
		CarpPosition []float64     `json:"carpPosition"`
		DipPosition  []float64     `json:"dipPosition"`
		Direction    []float64     `json:"direction"`
		Extended     bool          `json:"extended"`
		HandID       int           `json:"handId"`
		ID           int           `json:"id"`
		Length       float64       `json:"length"`
		McpPosition  []float64     `json:"mcpPosition"`
		PipPosition  []float64     `json:"pipPosition"`
		TimeVisible  float64       `json:"timeVisible"`
		TipPosition  []float64     `json:"tipPosition"`
		Type         int           `json:"type"`
		Width        float64       `json:"width"`
	} `json:"pointables"`
	Timestamp int64 `json:"timestamp"`
}