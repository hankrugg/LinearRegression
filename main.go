package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"strings"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Multiple Regression")

	content := widget.NewLabel("Select a data set to continue")

	var text string

	// Initialize with an empty canvas.Image
	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillOriginal

	res := widget.NewLabel("Results will be displayed here when a data set is selected")

	combo := widget.NewSelect([]string{"50_Startups.csv", "insurance.csv", "Salary_dataset.csv"}, func(value string) {
		text = processCSV(value) // Assign to outer variable
		res.SetText(text)        // Update label with processed data

		// Load the image corresponding to the selected dataset
		image.File = "predictions_vs_actual.png"
		image.Refresh()
	})

	scroll := container.NewVScroll(res)
	scroll.SetMinSize(fyne.NewSize(200, 200))

	myWindow.SetContent(container.NewVBox(content, combo, scroll, image))
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func processCSV(fileName string) string {
	data := LoadData(fileName)
	coef := train(data)

	// test the results on the training data
	results := test(coef, data)
	numRow, _ := results.Dims()

	p := plot.New()
	p.Title.Text = "Predictions (light blue) vs Actual Results (dark blue)"

	max := numRow
	if max > 20 {
		max = 20
	}

	predCharVals := make(plotter.Values, max)
	actuCharVals := make(plotter.Values, max)
	for i := 0; i < max; i++ {
		predCharVals[i] = results.At(i, 0)
		actuCharVals[i] = results.At(i, 1)
	}

	predChart, err := plotter.NewBarChart(predCharVals, vg.Points(10))
	if err != nil {
		log.Fatalf("could not create bar chart: %+v", err)
	}

	actuChart, err := plotter.NewBarChart(actuCharVals, vg.Points(10))
	if err != nil {
		log.Fatalf("could not create bar chart: %+v", err)
	}

	predChart.Color = color.RGBA{B: 255, A: 255}
	predChart.Offset = vg.Points(-5)

	actuChart.Color = color.RGBA{B: 100, A: 255}
	actuChart.Offset = vg.Points(5)

	p.Add(predChart, actuChart)

	// Save the plot to an image file
	err = p.Save(500, 250, "predictions_vs_actual.png")
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}

	// Prepare the string with results to print on the screen
	var resultString strings.Builder
	resultString.WriteString("Scatter plot saved as predictions_vs_actual.png\n")
	resultString.WriteString("Results:\n")
	for i := 0; i < numRow; i++ {
		resultString.WriteString(fmt.Sprintf("Prediction is %f. Actual result is %f\n", results.At(i, 0), results.At(i, 1)))
	}

	return resultString.String()
}
