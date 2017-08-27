package ngt

type ObjectType int

const (
	ObjectTypeNone ObjectType = iota
	ObjectTypeUint8
	ObjectTypeFloat
)

func (ot ObjectType) String() string {
	switch ot {
	case ObjectTypeNone:
		return "None"
	case ObjectTypeUint8:
		return "Integer"
	case ObjectTypeFloat:
		return "Float"
	default:
		return "Unknown"
	}
}

type DistanceType int

const (
	DistanceTypeNone DistanceType = iota - 1
	DistanceTypeL1
	DistanceTypeL2
	DistanceTypeHamming
	DistanceTypeAngle
	DistanceTypeCosine
)

func (dt DistanceType) String() string {
	switch dt {
	case DistanceTypeNone:
		return "None"
	case DistanceTypeL1:
		return "L1"
	case DistanceTypeL2:
		return "L2"
	case DistanceTypeHamming:
		return "Hamming"
	case DistanceTypeAngle:
		return "Angle"
	case DistanceTypeCosine:
		return "Cosine"
	default:
		return "Unknown"
	}
}
