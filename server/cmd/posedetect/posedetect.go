package posedetect

import (
	"fmt"
	"image"
	"image/color"
	"server/cmd/config"
	"time"

	"gocv.io/x/gocv"
)

func Detect(path string) string {
	protoFile := "./dataset/" + config.NewConfig().ProtoPoseDetectFileName
    weightsFile := "./dataset/" + config.NewConfig().WeightsPoseDetectFileName
    nPoints  := 18
    posePairs := [][]int{ {1,0}, {1,2},{1,5},{2,3},{3,4},{5,6},{6,7},{1,8},{8,9},{9,10},{1,11},{11,12},{12,13},{0,14},{0,15},{14,16},{15,17} }


	iMat := gocv.IMRead(path, gocv.IMReadColor)

	oMat := iMat.Clone()
	
 	frameWidth := oMat.Size()[1]
 	frameHeight := oMat.Size()[0]
	threshold := float32(0.1)

	net := gocv.ReadNetFromCaffe(protoFile, weightsFile)


	inpBlob := gocv.BlobFromImage(iMat, 1.0 / 255, image.Point{X:368, Y:368}, gocv.Scalar{Val1: 0,Val2: 0, Val3: 0,Val4: 0}, false, false)
	defer inpBlob.Close()

	net.SetInput(inpBlob, "")

	output := net.Forward("")
	defer output.Close()

	defer oMat.Close() 

	prob := output.Size()

	H, W := prob[2], prob[3]


	points := []image.Point{}

	for i := 0; i < nPoints; i++ {
		probMat, _ := output.FromPtr(H, W, gocv.MatTypeCV32FC1, 0, i)
		_, prob, _, point := gocv.MinMaxLoc(probMat)

		x := (frameWidth * point.X) / W
		y := (frameHeight * point.Y) / H

		if prob > threshold { 
			gocv.Circle(&oMat, image.Point{X: int(x), Y: int(y)}, 8, color.RGBA{R: 0, G: 255, B: 255}, -1)
			gocv.PutText(&oMat, fmt.Sprint(i),image.Point{X: int(x),Y: int(y) } , gocv.FontHersheySimplex, 1, color.RGBA{R: 0, B: 0, G: 255}, 2)


			points = append(points, image.Point{X: x, Y: y}) 
		} else {
			points = append(points, image.Point{X: -1, Y: -1})
		}
	}

	for i := 0; i < len(posePairs); i++ {
		pair := posePairs[i]
	
		partA := pair[0]
		partB := pair[1]

		mat1 := points[partA]
		mat2 := points[partB]
		
		if mat1.X != -1 && mat1.Y != -1 && mat2.X != -1 && mat2.Y != -1 {
			gocv.Line(&oMat, mat1, mat2, color.RGBA{R: 0, G: 255, B: 255}, 2)
			gocv.Circle(&oMat, mat1, 8, color.RGBA{R: 0, B: 0, G: 255}, -1)
		}
	}

	// gocv.Rotate(iMat, &oMat, gocv.Rotate180Clockwise)
	// gocv.Flip(iMat, &oMat, 1)
	// gocv.Blur(oMat, &oMat, image.Point{X: 40, Y: 40})
	// gocv.BoxFilter(oMat, &oMat, 5, image.Point{X: 30, Y: 30})

	// gocv.GaussianBlur(iMat, &oMat, image.Point{X: 5, Y: 9}, 50.0, 50.0, gocv.BorderReflect)

	fileName := fmt.Sprint(time.Now().UnixMicro()) + ".jpg"
	fmt.Print(fileName)

	gocv.IMWrite("uploads/" + fileName, oMat)

	return fileName
}