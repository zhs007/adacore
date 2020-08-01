package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/zhs007/adacore"
)

type dayData struct {
	data     string
	group    string
	brand    string
	playNums int
	users    int
	amount   float64
	won      float64
	gw       float64
	env      string
}

type dayStats struct {
	data      string
	groupNums int
	brandNums int
	playNums  int
	users     int
	amount    float64
	won       float64
	gw        float64
}

// Dataset - dataset
type Dataset struct {
	Data      []string
	GroupNums []int
	BrandNums []int
	PlayNums  []int
	Users     []int
	Amount    []float64
	Won       []float64
	GW        []float64
}

// RateOfChg - rate of change
type RateOfChg struct {
	Data      []string  `yaml:"data"`
	GroupNums []float32 `yaml:"groupnums"`
	BrandNums []float32 `yaml:"brandnums"`
	PlayNums  []float32 `yaml:"playnums"`
	Users     []float32 `yaml:"users"`
	Amount    []float32 `yaml:"amount"`
	Won       []float32 `yaml:"won"`
	GW        []float32 `yaml:"gw"`
}

// EnvPer - env per
type EnvPer struct {
	Env []string
	Val []float64
}

func addAmount(ep *EnvPer, env string, amount float64) {
	for i, v := range ep.Env {
		if v == env {
			ep.Val[i] = ep.Val[i] + amount

			return
		}
	}

	ep.Env = append(ep.Env, env)
	ep.Val = append(ep.Val, amount)
}

func loadSheet(f *excelize.File, sheetname string) ([]*dayData, error) {
	// sheetname := f.GetSheetName(index)
	arr, err := f.GetRows(sheetname)
	if err != nil {
		return nil, err
	}

	lst := []*dayData{}
	for y := 1; y < len(arr); y++ {
		cdd := &dayData{
			data:  strings.TrimSpace(arr[y][2]),
			group: arr[y][3],
			brand: arr[y][4],
			env:   strings.TrimSpace(arr[y][13]),
		}

		pn, err := strconv.Atoi(arr[y][8])
		if err != nil {
			return nil, err
		}

		cdd.playNums = pn

		u, err := strconv.Atoi(arr[y][9])
		if err != nil {
			return nil, err
		}

		cdd.users = u

		amount, err := strconv.ParseFloat(arr[y][10], 64)
		if err != nil {
			return nil, err
		}

		cdd.amount = amount

		won, err := strconv.ParseFloat(arr[y][11], 64)
		if err != nil {
			return nil, err
		}

		cdd.won = won

		gw, err := strconv.ParseFloat(arr[y][12], 64)
		if err != nil {
			return nil, err
		}

		cdd.gw = gw

		lst = append(lst, cdd)
	}

	return lst, nil
}

func excelChart(fn string) ([]*dayStats, []*dayData, error) {
	f, err := excelize.OpenFile(fn)
	if err != nil {
		return nil, nil, err
	}

	lst := []*dayData{}
	lststats := []*dayStats{}

	ms := f.GetSheetMap()
	for _, k := range ms {
		clst, err := loadSheet(f, k)
		if err != nil {
			return nil, nil, err
		}

		cs := &dayStats{
			data:      clst[0].data,
			brandNums: len(clst),
		}

		for _, v1 := range clst {
			if v1.env == "mt" {
				continue
			}

			cs.playNums = cs.playNums + v1.playNums
			cs.users = cs.users + v1.users
			cs.amount = cs.amount + v1.amount
			cs.won = cs.won + v1.won
			cs.gw = cs.gw + v1.gw
		}

		lst = append(lst, clst...)
		lststats = append(lststats, cs)
	}

	return lststats, lst, nil
}

func analyzeEnvPer(lst []*dayData, date string) (*EnvPer, error) {
	ep := &EnvPer{}

	for _, v := range lst {
		if v.data == date {
			addAmount(ep, v.env, v.amount)
		}
	}

	return ep, nil
}

func main() {
	err := adacore.InitTemplates()
	if err != nil {
		fmt.Printf("InitTemplates error %v", err)

		return
	}

	lsts, lst, err := excelChart("./excelchart.xlsx")
	if err != nil {
		fmt.Printf("excelChart err %v", err)
	}

	ep, err := analyzeEnvPer(lst, "2020-02-07")

	sort.Slice(lsts, func(i, j int) bool {
		return lsts[i].data < lsts[j].data
	})

	ds := &Dataset{}
	roc := &RateOfChg{}
	for _, v := range lsts {
		ds.Data = append(ds.Data, v.data)
		ds.GroupNums = append(ds.GroupNums, v.groupNums)
		ds.BrandNums = append(ds.BrandNums, v.brandNums)
		ds.PlayNums = append(ds.PlayNums, v.playNums)
		ds.Users = append(ds.Users, v.users)
		ds.Amount = append(ds.Amount, v.amount)
		ds.Won = append(ds.Won, v.won)
		ds.GW = append(ds.GW, v.gw)
	}

	roc.Data = ds.Data[1:]
	roc.BrandNums = adacore.NewRateOfChgInt(ds.BrandNums)
	roc.PlayNums = adacore.NewRateOfChgInt(ds.PlayNums)
	roc.Users = adacore.NewRateOfChgInt(ds.Users)
	roc.Amount = adacore.NewRateOfChgFloat64(ds.Amount)
	roc.Won = adacore.NewRateOfChgFloat64(ds.Won)
	roc.GW = adacore.NewRateOfChgFloat64(ds.GW)

	md := adacore.NewMakrdown("Ada Core")

	md.AppendDataset("ds001", ds)
	md.AppendDataset("ds002", roc)

	md.AppendChartBar(&adacore.ChartBar{
		ID:          "bar001",
		DatasetName: "ds001",
		Title:       "amount bar",
		LegendData:  []string{"amount"},
		XType:       "category",
		XData:       "data",
		XShowAll:    true,
		YType:       "value",
		YData: []adacore.ChartBasicData{
			adacore.ChartBasicData{
				Name: "amount",
				Data: "amount",
			},
		},
	})

	md.AppendChartBar(&adacore.ChartBar{
		ID:          "bar005",
		DatasetName: "ds001",
		Title:       "users bar",
		LegendData:  []string{"uesrs"},
		XType:       "category",
		XData:       "data",
		XShowAll:    true,
		YType:       "value",
		YData: []adacore.ChartBasicData{
			adacore.ChartBasicData{
				Name: "users",
				Data: "users",
			},
		},
	})

	md.AppendChartBar(&adacore.ChartBar{
		ID:          "bar006",
		DatasetName: "ds001",
		Title:       "users bar",
		LegendData:  []string{"playnums"},
		XType:       "category",
		XData:       "data",
		XShowAll:    true,
		YType:       "value",
		YData: []adacore.ChartBasicData{
			adacore.ChartBasicData{
				Name: "playnums",
				Data: "playnums",
			},
		},
	})

	md.AppendChartBar(&adacore.ChartBar{
		ID:          "bar002",
		DatasetName: "ds001",
		Title:       "gw bar",
		LegendData:  []string{"gw"},
		XType:       "category",
		XData:       "data",
		XShowAll:    true,
		YType:       "value",
		YData: []adacore.ChartBasicData{
			adacore.ChartBasicData{
				Name: "gw",
				Data: "gw",
			},
		},
	})

	md.AppendChartBar(&adacore.ChartBar{
		ID:          "bar003",
		DatasetName: "ds002",
		Title:       "per bar",
		LegendData:  []string{"brandnums", "playnums", "users", "amount", "gw"},
		XType:       "category",
		XData:       "data",
		XShowAll:    true,
		YType:       "value",
		YData: []adacore.ChartBasicData{
			adacore.ChartBasicData{
				Name: "brandnums",
				Data: "brandnums",
			},
			adacore.ChartBasicData{
				Name: "playnums",
				Data: "playnums",
			},
			adacore.ChartBasicData{
				Name: "users",
				Data: "users",
			},
			adacore.ChartBasicData{
				Name: "amount",
				Data: "amount",
			},
			adacore.ChartBasicData{
				Name: "gw",
				Data: "gw",
			},
		},
	})

	_, err = md.AppendDataset("envperds", ep)
	if err != nil {
		fmt.Printf("AppendDataset error %v", err)

		return
	}

	_, err = md.AppendChartPie(&adacore.ChartPie{
		ID:          "pie001",
		DatasetName: "envperds",
		Title:       "Pie",
		SubText:     "test pie chart",
		A:           "pie name",
		BVal:        "env",
		CVal:        "val",
		Sort:        adacore.ChartSortReverse,
	})
	if err != nil {
		fmt.Printf("AppendChartPie error %v", err)

		return
	}

	err = ioutil.WriteFile("./output.md", []byte(md.GetMarkdownString(nil)), 0644)
	if err != nil {
		fmt.Printf("WriteFile err %v", err)
	}
}
