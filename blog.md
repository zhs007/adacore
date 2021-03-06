# AdaCore Development Log

### 2019-10-05

``excelize``是一个excel库，360的，getrows返回的数据是一个完整的数据，前面有空行或空列，也会有数据。  
所以这个坐标是绝对的坐标。  

如果表格不是从右上开始的，需要获取到起点，可以参考这段代码。

``` golang
func isEmptyRow(arr [][]string, x int) bool {
	if x >= 0 && x < len(arr[0]) {
		for y := 0; y < len(arr); y++ {
			cr := strings.TrimSpace(arr[y][x])
			if cr != "" {
				return false
			}
		}

		return true
	}

	return true
}

func isEmptyColumn(arr [][]string, y int) bool {
	if y >= 0 && y < len(arr) {
		for x := 0; x < len(arr[0]); x++ {
			cr := strings.TrimSpace(arr[y][x])
			if cr != "" {
				return false
			}
		}

		return true
	}

	return true
}

// GetStartXY - get start x & y
func GetStartXY(arr [][]string) (int, int) {
	cx := 0
	for x := 0; x < len(arr[0]); x++ {
		if !isEmptyRow(arr, x) {
			cx = x

			break
		}
	}

	cy := 0
	for y := 0; y < len(arr); y++ {
		if !isEmptyColumn(arr, y) {
			cy = y

			break
		}
	}

	return cx, cy
}
```

``CoordinatesToCellName``这个接口，传入的坐标是从1开始的。

``GetComments``返回的是全部表格的数据，可以参考这个函数使用。

``` golang
// GetComments - get comments with sheetName
func GetComments(f *excelize.File, sheetName string) map[string]string {
	mc := f.GetComments()
	cursheetcomments, isok := mc[sheetName]
	if !isok {
		return nil
	}

	mapComments := make(map[string]string)
	for _, v := range cursheetcomments {
		mapComments[v.Ref] = v.Text
	}

	return mapComments
}
```

``AddComment``接口插入数据时，类似如下：

``` golang
err := f.AddComment("Sheet1", "A30", `{"author":"Excelize: ","text":"This is a comment."}`)
```

这时，如果不传入 author，会默认帮你生成一个空的 Author字符串，如果传入空格，则不会生成Author，但使用excel打开文件时会提示报错，但文件还是可以打开。

### 2019-10-04

今天开始处理``excel``文件了。  
第一版不考虑问问题，而是尽可能的表格化数据，因为可能用户自己都不知道希望怎么表现数据，我们尽可能的帮他把数据可视化，然后再让他决定如何展现数据。

- 需要保证head都有值，如果没有值，默认给成 ``__colume?__`` 这样的。
- 需要将时间的列筛选出来，这里可能有多列时间数据，可能还需要按照一个粒度合并数据。
- 需要将类别的列筛选出来，这部分应该是最麻烦的，类别数据很难从文字上区分，我建议取一个长度以内的，有重复内容的字段作为类别备选，而且，优先选取前面的列。
- 需要将数值类型的列筛选出来。
- 数据上，最终应该就是3种，一种是序列图，可能是时间序列也可能是类别序列；一种是饼图，还有一种是树状图。
- ID列，每一列都是唯一的，如果有一列非数字的ID列，就应该忽略数字ID列。
- 数字ID列，从0或者从1开始排列的唯一标识列。
- 百分数，应该一定输出饼图。

数据列类型：

- 数字ID列，一般为主键列，从0或者从1开始排列的唯一标识列，如果该表格中存在ID列，则忽略数字ID列。
- ID列，每行都唯一的字符串列。
- 类型列，有重复的字符串列。
- 描述列，较长的字符串列，一般不做统计分析用。
- 时间列，时间字符串列，或时间戳数字列（秒或者毫秒）。
- 数值列，每行数据都是数值的列。

先按自己的思路做分析，不对或者不准确都没关系，后面再来进一步优化处理。

### 2019-10-03

今天将Ada部署到外网了，只有``telegram``，这里可以添加：[Ada](https://t.me/@ada_heyalgo_bot)。

现在只是一个很基础的版本，可以帮你渲染markdown文件，是按Ada要求的格式渲染，所以一些基本的图表都可以支持。

后面会有更多功能添加进来的。