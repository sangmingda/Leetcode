package _733_图像渲染

import (
	"reflect"
	"testing"
)

//  * 示例 1:
//  * 
//  * 
//  * 输入: 
//  * image = [[1,1,1],[1,1,0],[1,0,1]]
//  * sr = 1, sc = 1, newColor = 2
//  * 输出: [[2,2,2],[2,2,0],[2,0,1]]
//  * 解析: 
//  * 在图像的正中间，(坐标(sr,sc)=(1,1)),
//  * 在路径上所有符合条件的像素点的颜色都被更改成2。
//  * 注意，右下角的像素没有更改为2，
//  * 因为它不是在上下左右四个方向上与初始点相连的像素点。
func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
			[][]int{
				[]int{1,1,1},[]int{1,1,0},[]int{1,0,1},
			},
			1,
			1,
			2,
			},
			[][]int{
	[]int{2,2,2},[]int{2,2,0},[]int{2,0,1},
},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
