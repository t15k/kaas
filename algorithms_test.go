package main

import (
	"math"
	"testing"
)

func TestUnDef(t *testing.T) {
	if unDef(math.NaN()) != true || unDef(math.Inf(-1)) != true || unDef(math.Inf(1)) != true {
		t.Fatal("unDef() should have returned true but returned false")
	}
	if unDef(1) != false {
		t.Fatal("unDef() was provided with a valid number but returned true")
	}
}

func TestRound(t *testing.T) {
	if round(0.0, 1) != 0.0 {
		t.Fatal("round(0.0, 1) should be 0.0 but got", round(0.0, 1))
	}
	if round(0.12345, 2) != 0.12 {
		t.Fatal("round(0.12345, 2) should be 0.12 but got", round(0.12345, 2))
	}
	if round(1.3456, 2) != 1.35 {
		t.Fatal("round(1.3456, 2) should be 1.35 but got", round(1.3456, 2))
	}
}

func TestMean(t *testing.T) {
	ts := []float64{5.639145579739498, 5.938634058698721, 4.1078877195188355, 0.8514291847212441, 5.5867610325912835, 9.370561508108866, 8.041583970902263, 6.661064864421492, 4.244556640773159, 5.74105756677287, 6.464063629139909, 6.712758257547336, 1.820294533703306, 8.897756631040881, 0.44206115293154324}
	if mean(ts) != 5.3679744220407475 {
		t.Fatal("Mean should be 5.3679744220407475 but is", mean(ts))
	}
	emptySlice := []float64{}
	if mean(emptySlice) != 0.0 {
		t.Fatal("mean(0) should be 0 but was", mean(emptySlice))
	}

}

func TestMedian(t *testing.T) {
	series := []float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.01}
	if median(series) != 5.05 {
		t.Fatal("wrong median", median(series))
	}
	emptySlice := []float64{}
	if median(emptySlice) != 0.0 {
		t.Fatal("median of [0.0] should be 0 but was", median(emptySlice))
	}
}

func TestVar(t *testing.T) {
	ts := []float64{0.10324773173787438, 8.690175411623171, 3.0127094333698503, 7.229031170368377, 1.1390009044076699, 7.025926079441213, 4.159771728162745, 6.360868308198036, 6.062490139216057, 1.6340976374917215, 2.13104308511487, 9.619084712684327, 3.7323056512984296, 8.956329024456558, 4.985841767863985}
	if round(variance(ts), 6) != round(9.0975915672960674, 6) {
		t.Fatal("Variance is 9.0975915672960674 but was calculated as", variance(ts))
	}

}

func TestCov(t *testing.T) {
	ts1 := []float64{4.195242529398381, 6.80733043083745, 0.7455287999771032, 7.70070089351621, 3.089962180134008, 6.288179887517467, 4.596409177423682, 2.1212924647965092, 7.636150996552125, 1.016192510562095, 7.070371560695419, 8.053412735504203, 3.8479711237694514, 9.525871454119626, 5.6180940623256905}
	ts2 := []float64{6.297991472572121, 5.770316637293231, 8.698992152328595, 1.8987421150485562, 3.2593820772925652, 3.1966296468963504, 1.8250611430569141, 8.698938964272042, 9.059364862384314, 5.272066372193422, 7.915633798761328, 3.321459973870513, 9.732051917378959, 7.303905540791, 4.085083512027085}
	if cov(ts1, ts2) != -1.4068001425690473 {
		t.Fatal("Covariance is -1.4068001425690473 but was calculated as", cov(ts1, ts2))
	}
	tsA := []float64{4.195242529398381, 6.80733043083745, 0.7455287999771032}
	tsB := []float64{0.0}
	emptySlice := []float64{}
	if cov(tsA, tsB) != 0.0 {
		t.Fatal("cov() was provide with slices of different length and should have returned 0.0 but returned", cov(tsA, tsB))
	}
	if cov(tsA, emptySlice) != 0.0 {
		t.Fatal("cov() was provided with an empty slice and should have returned 0.0 but returned", cov(tsA, emptySlice))
	}
}

func TestTailAvg(t *testing.T) {
	ts := []float64{4.195242529398381, 6.80733043083745, 0.7455287999771032, 7.70070089351621, 3.089962180134008, 6.288179887517467}
	if round(tailAvg(ts), 6) != round(5.6929476537225625, 6) {
		t.Fatal("tailAvg() should have returned 5.6929476537225625 but returned", tailAvg(ts))
	}
	if tailAvg([]float64{3.0}) != 3.0 {
		t.Fatal("tailAvg(3.0) should be 3.0 but returned", tailAvg([]float64{3.0}))
	}
	if tailAvg([]float64{}) != 0.0 {
		t.Fatal("tailAvg(0.0) should be 0.0 but returned", tailAvg([]float64{}))
	}
}

func TestStd(t *testing.T) {
	ts := []float64{0.6652356971378492, 2.828082160729557, 4.492799589097807, 6.4885349866234066, 8.323505050316992, 4.235336161652312, 2.6864488789516905, 9.315871316707883, 6.196127077653522, 1.0475738614605756, 3.6130700415059644, 8.580966992844761, 9.787803840922486, 7.319726729729728, 2.5759349867985595}
	if round(std(ts), 6) != round(3.0107567623831053, 6) {
		t.Fatal("Standard deviation is 3.0107567623831053 but was calculated as", std(ts))
	}

}

func TestLinearRegressionLSE(t *testing.T) {
	var ts []Measurement
	for i := 0; i < 10; i++ {
		a := Measurement{
			timestamp: int64(i),
			value:     float64(i)*3.1 - 2.1,
		}
		ts = append(ts, a)
	}
	c, m := linearRegressionLSE(ts)

	if round(m, 1) != 3.1 || round(c, 1) != -2.1 {
		t.Fatal("wrong linearregressionlse", "m", m, "c", c)
	}
}

func TestEwma(t *testing.T) {
	series := []float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.01}
	rst := []float64{0.09999999999999978, 0.6554455445544544, 1.214520977649978, 1.7772255876832508, 2.3435583786886025, 2.9135180706168184, 3.48710309969332, 4.064311618855566, 4.645141498269393, 5.121538107701817}
	rt := ewma(series, 50)
	for i, v := range rt {
		if v != rst[i] {
			t.Fatal("ewma error", t)
		}
	}
	retLen := len(ewma([]float64{}, 10))
	if retLen != 0 {
		t.Fatal("ewma() was provided with empty series and should have returned empty series, but returned", ewma([]float64{}, 10))
	}
}

func TestEwmStd(t *testing.T) {
	series := []float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.01}
	rst := []float64{4.9526750297502914e-09, 0.5527160659008843, 0.902537317532201, 1.2357653238068602, 1.5629953497356235, 1.8872927402148911, 2.209889422762198, 2.531374353771067, 2.8520607676954124, 3.0195071357543375}
	rt := ewmStd(series, 50)
	for i, v := range rt {
		if v != rst[i] {
			t.Fatal("ewma std error", t)
		}
	}
}

func TestHistogram(t *testing.T) {
	series := []float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.01}
	hist := []int{1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 2}
	bin := []float64{0.1, 0.694, 1.288, 1.8820000000000001, 2.476, 3.07, 3.664, 4.257999999999999, 4.851999999999999, 5.446, 6.039999999999999, 6.6339999999999995, 7.228, 7.821999999999999, 8.415999999999999, 9.01}
	h, b := histogram(series, 15)
	for i, v := range h {
		if v != hist[i] {
			t.Fatal("hist error", t)
		}
	}
	for i, v := range b {
		if v != bin[i] {
			t.Fatal("hist error", t)
		}
	}
	emptyHist, emptyHistEdges := histogram([]float64{}, 1)
	if len(emptyHist) != 0 || len(emptyHistEdges) != 0 {
		t.Fatal("histogram was provided with empty series and should have returend two empty slices but returned", emptyHist, emptyHistEdges)
	}

}

func TestKS2Samp(t *testing.T) {
	reference := []float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.01, 1.2, 2, 4, 6, 9, 1, 22, 11, 19, 18.9, 11, 14}
	probe := []float64{0.4, 0.1, 1.3, 2.4, 6.5, 3.6, 5.7, 6.8, 8.9, 9, 9.1, 11.2, 1.2, 1.3, 14, 4, 5, 0.123, 9, 7, 8.1, 9.9, 2.1}
	ksD, ksPValue := kS2Samp(reference, probe)
	if round(ksD, 6) != round(0.18577075098814222, 6) || ksPValue != 0.789955481957006 {
		t.Fatal("testks2samp error\t", ksD, ksPValue, t)
	}
}

func TestMedianAbsoluteDeviation(t *testing.T) {
	goodSeries := []float64{0.6652356971378492, 2.828082160729557, 4.492799589097807, 6.4885349866234066, 8.323505050316992, 4.235336161652312, 2.6864488789516905, 9.315871316707883, 6.196127077653522, 1.0475738614605756, 3.6130700415059644, 8.580966992844761, 9.787803840922486, 7.319726729729728, 2.5759349867985595}
	anomSeries := []float64{0.6652356971378492, 2.828082160729557, 4.492799589097807, 6.4885349866234066, 8.323505050316992, 4.235336161652312, 2.6864488789516905, 9.315871316707883, 6.196127077653522, 1.0475738614605756, 3.6130700415059644, 8.580966992844761, 9.787803840922486, 7.319726729729728, 500.5759349867985595}
	if medianAbsoluteDeviation(goodSeries) != false {
		t.Fatal("medianAbsoluteDeviation() returned true for good series")
	}
	if medianAbsoluteDeviation(anomSeries) != true {
		t.Fatal("medianAbsoluteDeviation() returned false for bad series")
	}
	if medianAbsoluteDeviation([]float64{}) != false || medianAbsoluteDeviation([]float64{0.0}) != false {
		t.Fatal("medianAbsoluteDeviation() was incorrect for empty or slice with only 0.0 calculation")
	}
}

func TestSimpleStddevFromMovingAverage(t *testing.T) {
	tsAnom := []float64{2.981327510952622, 3.1352498611087554, 5.082869663872875, 6.618291099712494, 2.2608586618361413, 2.4522340531396924, 1.0148059366821838, 9.219352536115258, 2.153918824176978, 6.475097733614631, 0.6411545069161773, 7.652087952609515, 6.285300598985705, 0.28238375542215643, 1.5854977285624505, 2.375281414351995, 6.814109597000528, 6.96357476665019, 4.727754996793142, 1.118482131471743, 7.660645519367183, 7.7212910430357375, 7.578089213066831, 7.665175737483606, 9.268902846067077, 9.665652781345235, 7.8771181419967355, 0.5166381780239959, 0.6471254304615881, 1.669393381801093, 7.477733011772495, 7.455780680178977, 2.061197844872779, 7.826621975872231, 9.511205398681653, 2.354250680483746, 9.049518493859598, 9.622123796656325, 8.007861466713557, 5.430623799519938, 1.8381616240646426, 0.9328210092651534, 4.0911323710451, 4.75099822844837, 1.3326143721882389, 4.318490584455798, 7.517310467011012, 7.04056011225794, 4.574055602064595, 8.462497972817147, 0.14308108484967996, 2.64421409184193, 4.329087261780812, 1.305751882474555, 9.324932570977516, 1.5340505850988573, 7.861765988207504, 3.515003972006415, 1.2117875334678707, 1.408833655562104, 9.905754134627012, 5.8319688144920185, 2.5482369545436443, 5.4600466813010105, 9.341127265913212, 8.453858158706158, 3.204501612449955, 6.502273946158131, 5.489442374801488, 0.3314990469030066, 3.0346949000616776, 2.244153891218428, 2.5568366448202307, 8.85880574200714, 5.168669854669171, 5.965777942490709, 1.2110230091452923, 7.128851540774549, 2.938729168551598, 2.726213899652307, 0.31501219589320395, 3.723776517401748, 9.108478759168142, 2.038578200373238, 5.923780268323395, 5.615480620443757, 8.716642455624063, 0.47370203635324404, 6.783734820108117, 4.168345044679997, 6.857407055551164, 2.365374210837686, 8.382383385809273, 7.345611298483753, 8.495616341319042, 3.3503863887555054, 9.40398543878947, 8.755458549584516, 0.25563479422747504, 400.973355700794282}
	tsNormal := []float64{2.981327510952622, 3.1352498611087554, 5.082869663872875, 6.618291099712494, 2.2608586618361413, 2.4522340531396924, 1.0148059366821838, 9.219352536115258, 2.153918824176978, 6.475097733614631, 0.6411545069161773, 7.652087952609515, 6.285300598985705, 0.28238375542215643, 1.5854977285624505, 2.375281414351995, 6.814109597000528, 6.96357476665019, 4.727754996793142, 1.118482131471743, 7.660645519367183, 7.7212910430357375, 7.578089213066831, 7.665175737483606, 9.268902846067077, 9.665652781345235, 7.8771181419967355, 0.5166381780239959, 0.6471254304615881, 1.669393381801093, 7.477733011772495, 7.455780680178977, 2.061197844872779, 7.826621975872231, 9.511205398681653, 2.354250680483746, 9.049518493859598, 9.622123796656325, 8.007861466713557, 5.430623799519938, 1.8381616240646426, 0.9328210092651534, 4.0911323710451, 4.75099822844837, 1.3326143721882389, 4.318490584455798, 7.517310467011012, 7.04056011225794, 4.574055602064595, 8.462497972817147, 0.14308108484967996, 2.64421409184193, 4.329087261780812, 1.305751882474555, 9.324932570977516, 1.5340505850988573, 7.861765988207504, 3.515003972006415, 1.2117875334678707, 1.408833655562104, 9.905754134627012, 5.8319688144920185, 2.5482369545436443, 5.4600466813010105, 9.341127265913212, 8.453858158706158, 3.204501612449955, 6.502273946158131, 5.489442374801488, 0.3314990469030066, 3.0346949000616776, 2.244153891218428, 2.5568366448202307, 8.85880574200714, 5.168669854669171, 5.965777942490709, 1.2110230091452923, 7.128851540774549, 2.938729168551598, 2.726213899652307, 0.31501219589320395, 3.723776517401748, 9.108478759168142, 2.038578200373238, 5.923780268323395, 5.615480620443757, 8.716642455624063, 0.47370203635324404, 6.783734820108117, 4.168345044679997, 6.857407055551164, 2.365374210837686, 8.382383385809273, 7.345611298483753, 8.495616341319042, 3.3503863887555054, 9.40398543878947, 8.755458549584516, 0.25563479422747504, 4.973355700794282}
	if simpleStddevFromMovingAverage(tsAnom) != true {
		t.Fatal("simpleStddevFromMovingAverage() should return true but returned false")
	}
	if simpleStddevFromMovingAverage(tsNormal) != false {
		t.Fatal("simpleStddevFromMovingAverage() should return false but returned true")
	}
}

func TestStddevFromMovingAverage(t *testing.T) {
	tsAnom := []float64{2.981327510952622, 3.1352498611087554, 5.082869663872875, 6.618291099712494, 2.2608586618361413, 2.4522340531396924, 1.0148059366821838, 9.219352536115258, 2.153918824176978, 6.475097733614631, 0.6411545069161773, 7.652087952609515, 6.285300598985705, 0.28238375542215643, 1.5854977285624505, 2.375281414351995, 6.814109597000528, 6.96357476665019, 4.727754996793142, 1.118482131471743, 7.660645519367183, 7.7212910430357375, 7.578089213066831, 7.665175737483606, 9.268902846067077, 9.665652781345235, 7.8771181419967355, 0.5166381780239959, 0.6471254304615881, 1.669393381801093, 7.477733011772495, 7.455780680178977, 2.061197844872779, 7.826621975872231, 9.511205398681653, 2.354250680483746, 9.049518493859598, 9.622123796656325, 8.007861466713557, 5.430623799519938, 1.8381616240646426, 0.9328210092651534, 4.0911323710451, 4.75099822844837, 1.3326143721882389, 4.318490584455798, 7.517310467011012, 7.04056011225794, 4.574055602064595, 8.462497972817147, 0.14308108484967996, 2.64421409184193, 4.329087261780812, 1.305751882474555, 9.324932570977516, 1.5340505850988573, 7.861765988207504, 3.515003972006415, 1.2117875334678707, 1.408833655562104, 9.905754134627012, 5.8319688144920185, 2.5482369545436443, 5.4600466813010105, 9.341127265913212, 8.453858158706158, 3.204501612449955, 6.502273946158131, 5.489442374801488, 0.3314990469030066, 3.0346949000616776, 2.244153891218428, 2.5568366448202307, 8.85880574200714, 5.168669854669171, 5.965777942490709, 1.2110230091452923, 7.128851540774549, 2.938729168551598, 2.726213899652307, 0.31501219589320395, 3.723776517401748, 9.108478759168142, 2.038578200373238, 5.923780268323395, 5.615480620443757, 8.716642455624063, 0.47370203635324404, 6.783734820108117, 4.168345044679997, 6.857407055551164, 2.365374210837686, 8.382383385809273, 7.345611298483753, 8.495616341319042, 3.3503863887555054, 9.40398543878947, 8.755458549584516, 0.25563479422747504, 400.973355700794282}
	tsNormal := []float64{2.981327510952622, 3.1352498611087554, 5.082869663872875, 6.618291099712494, 2.2608586618361413, 2.4522340531396924, 1.0148059366821838, 9.219352536115258, 2.153918824176978, 6.475097733614631, 0.6411545069161773, 7.652087952609515, 6.285300598985705, 0.28238375542215643, 1.5854977285624505, 2.375281414351995, 6.814109597000528, 6.96357476665019, 4.727754996793142, 1.118482131471743, 7.660645519367183, 7.7212910430357375, 7.578089213066831, 7.665175737483606, 9.268902846067077, 9.665652781345235, 7.8771181419967355, 0.5166381780239959, 0.6471254304615881, 1.669393381801093, 7.477733011772495, 7.455780680178977, 2.061197844872779, 7.826621975872231, 9.511205398681653, 2.354250680483746, 9.049518493859598, 9.622123796656325, 8.007861466713557, 5.430623799519938, 1.8381616240646426, 0.9328210092651534, 4.0911323710451, 4.75099822844837, 1.3326143721882389, 4.318490584455798, 7.517310467011012, 7.04056011225794, 4.574055602064595, 8.462497972817147, 0.14308108484967996, 2.64421409184193, 4.329087261780812, 1.305751882474555, 9.324932570977516, 1.5340505850988573, 7.861765988207504, 3.515003972006415, 1.2117875334678707, 1.408833655562104, 9.905754134627012, 5.8319688144920185, 2.5482369545436443, 5.4600466813010105, 9.341127265913212, 8.453858158706158, 3.204501612449955, 6.502273946158131, 5.489442374801488, 0.3314990469030066, 3.0346949000616776, 2.244153891218428, 2.5568366448202307, 8.85880574200714, 5.168669854669171, 5.965777942490709, 1.2110230091452923, 7.128851540774549, 2.938729168551598, 2.726213899652307, 0.31501219589320395, 3.723776517401748, 9.108478759168142, 2.038578200373238, 5.923780268323395, 5.615480620443757, 8.716642455624063, 0.47370203635324404, 6.783734820108117, 4.168345044679997, 6.857407055551164, 2.365374210837686, 8.382383385809273, 7.345611298483753, 8.495616341319042, 3.3503863887555054, 9.40398543878947, 8.755458549584516, 0.25563479422747504, 4.973355700794282}
	if stddevFromMovingAverage(tsAnom) != true {
		t.Fatal("simpleStddevFromMovingAverage() should return true but returned false")
	}
	if stddevFromMovingAverage(tsNormal) != false {
		t.Fatal("simpleStddevFromMovingAverage() should return false but returned true")
	}
}

func TestMeanSubtractionCumulation(t *testing.T) {
	tsAnom := []float64{8.359239145572921, 4.3382786304085705, 0.7268435093236214, 9.75731297595692, 8.629253088217913, 2.7368693662546075, 2.0098388082853935, 2.1853108829852586, 6.039251161723268, 2.0906302584742322, 4.259970760914222, 0.3695083869607618, 0.05900961227263579, 2.5594287166993315, 30.198482161483356}
	tsNorm := []float64{8.359239145572921, 4.3382786304085705, 0.7268435093236214, 9.75731297595692, 8.629253088217913, 2.7368693662546075, 2.0098388082853935, 2.1853108829852586, 6.039251161723268, 2.0906302584742322, 4.259970760914222, 0.3695083869607618, 0.05900961227263579, 2.5594287166993315, 3.198482161483356}
	if meanSubtractionCumulation(tsNorm) != false {
		t.Fatal("should be false")
	}
	if meanSubtractionCumulation(tsAnom) != true {
		t.Fatal("should be true")
	}
}

func TestLeastSquares(t *testing.T) {
	tsNorm := []float64{8.9410498340368, 9.222689023121966, 6.299456244144405, 7.159242199559776, 2.5617962396565854, 8.725564230852438, 7.483008310130174, 0.4523045173851181, 7.02552347813639, 6.561397821476246, 3.2100225901878665, 1.749620987917102, 8.740918605579015, 7.7702583415980975, 4.940381124308742}
	tsAnom := []float64{8.9410498340368, 9.222689023121966, 6.299456244144405, 7.159242199559776, 2.5617962396565854, 8.725564230852438, 7.483008310130174, 0.4523045173851181, 7.02552347813639, 6.561397821476246, 3.2100225901878665, 1.749620987917102, 8.740918605579015, 7.7702583415980975, 8.359239145572921, 4.3382786304085705, 0.7268435093236214, 9.75731297595692, 8.629253088217913, 2.7368693662546075, 2.0098388082853935, 2.1853108829852586, 6.039251161723268, 2.0906302584742322, 4.259970760914222, 0.3695083869607618, 0.05900961227263579, 2.5594287166993315, 3.198482161483356, 40000.940381124308742}
	var measurementsNorm Measurements
	var measurementsAnom Measurements
	for i, v := range tsNorm {
		measurementsNorm = append(measurementsNorm, Measurement{v, int64(i)})
	}
	for i, v := range tsAnom {
		measurementsAnom = append(measurementsAnom, Measurement{v, int64(i)})
	}

	if leastSquares(measurementsNorm) != false {
		t.Fatal("should be false")
	}
	if leastSquares(measurementsAnom) != true {
		t.Fatal("should be true")
	}
}
