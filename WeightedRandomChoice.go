package WeightedRandomChoice

import (
	"math/rand"
	"sort"
	"time"
)

type WeightedRandomChoice struct {
	Elements          [] string
	Weights           [] int
	TotalWeight       int
	calibratedWeights bool
	precision         int
	calibrateValue    int
}

func (wrs *WeightedRandomChoice) AddElement(element string, weight int) {
	weight *= wrs.calibrateValue
	i := sort.Search(len(wrs.Weights), func(i int) bool { return wrs.Weights[i] > weight })
	wrs.Weights = append(wrs.Weights, 0)
	wrs.Elements = append(wrs.Elements, "")
	copy(wrs.Weights[i+1:], wrs.Weights[i:])
	copy(wrs.Elements[i+1:], wrs.Elements[i:])
	wrs.Weights[i] = weight
	wrs.Elements[i] = element
	wrs.TotalWeight += weight
}

func (wrs *WeightedRandomChoice) AddElements(elements map[string]int) {
	for element, weight := range elements{
		wrs.AddElement(element, weight)
	}
}

func (wrs *WeightedRandomChoice) GetRandomChoice() string {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(wrs.TotalWeight)
	if !wrs.calibratedWeights {
		wrs.calibrateWeights()
	}
	for key, weight := range wrs.Weights {
		value -= weight
		if value <= 0 {
			return wrs.Elements[key]
		}
	}
	return ""
}

func (wrs *WeightedRandomChoice) calibrateWeights() {
	if wrs.TotalWeight/wrs.precision < 1 {
		wrs.calibrateValue = wrs.precision / wrs.TotalWeight
		wrs.TotalWeight = 0
		for key := range wrs.Weights {
			wrs.Weights[key] *= wrs.calibrateValue
			wrs.TotalWeight += wrs.Weights[key]
		}
		wrs.calibratedWeights = true
	}
}

func New(arguments ...int) WeightedRandomChoice {
	var precision = 1000
	if len(arguments) > 0 {
		precision = arguments[0]
	}
	return WeightedRandomChoice{
		precision:         precision,
		calibratedWeights: false,
		calibrateValue:    1,
	}
}
