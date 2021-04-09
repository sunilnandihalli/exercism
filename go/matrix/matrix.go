package matrix
import "strings"
import "strconv"
import "errors"


type Matrix interface {
     Rows() [][]int
     Cols() [][]int
     Set(int,int,int) bool
}

type mymat struct {
     m [][]int
     num_rows int
     num_cols int
}

func New(s string) (m Matrix,e error) {
     rows:=strings.Split(s,"\n")
     num_rows := len(rows)
     if num_rows==0 {
     	m = nil
	e = errors.New("empty input")
	return
     }
     
     num_cols := len(strings.Split(strings.Trim(rows[0]," \n")," "))
     if num_cols == 0 {
     	m  = nil
	e = errors.New("empty rows")
     }
     vals := make([][]int,num_rows)
     for i,row :=range rows {
     	 vals[i] = make([]int,num_cols)
	 row_vals := strings.Split(strings.Trim(row," \n")," ")
	 if len(row_vals)!=num_cols {
	    e = errors.New("uneven row_sizes");
	    m = nil
	    return
	 }
	 for j,v :=range row_vals {
	 val,err := strconv.Atoi(v)
	     if err!=nil {
	     	e = err
		m = nil
		return
	     }
	     vals[i][j] = val
	 }
     }
     m = mymat{m:vals,num_rows:num_rows,num_cols:num_cols}
     e = nil
     return
}

func (m mymat) Rows() [][]int {
     ret := make([][]int, m.num_rows)
     for i,_ := range ret {
     	 ret[i] = make([]int,m.num_cols)
	 for j:=0;j<m.num_cols;j++ {
	     ret[i][j] = m.m[i][j]
	 }
     }
     return ret
}

func (m mymat) Cols() [][]int {
     ret := make([][]int, m.num_cols)
     for i,_ := range ret {
     	 ret[i] = make([]int,m.num_rows)
	 for j:=0;j<m.num_rows;j++ {
	     ret[i][j] = m.m[j][i]
	 }
     }
     return ret
}

func (m mymat) Set(i, j, v int) bool {
     if i>= m.num_rows || j>= m.num_cols || i<0 || j<0 {
     	return false
     }
     m.m[i][j] = v
     return true
}