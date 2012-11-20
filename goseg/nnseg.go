package goseg

import (
	"./gonn"
	"fmt"
	"os"
	"time"
)

var nn * gonn.PCNNetwork

func getFeature(pattern []int) map[string]float64{
	X := make(map[string]float64)
	for pos,x := range pattern{
		key := fmt.Sprintf("C%d==%d",pos,x)
		X[key]=1.0
		if pos<len(pattern)-1{
			key := fmt.Sprintf("C%d&%d==%d&%d",pos,pos+1,x,pattern[pos+1])
			X[key]=1.0
		}
	}
	key := fmt.Sprintf("C1&3==%d&%d",pattern[1],pattern[3])
	X[key]=1.0
	return X
}


func getWindow(ary []rune, i int) map[string]float64{
	pattern := make([]int,0)
	for j:=i-2;j<=i+2;j++{
		if j<0 || j>=len(ary){
			pattern = append(pattern,int(' '))
		}else{
			pattern = append(pattern,int(ary[j]))
		}
	}

	X := getFeature(pattern)

	return X
}

func getTags(N int) []int{
	result := make([]int,N)
	if N==1{
		result[0] = 3 //'s'
	}else{
		result[0] = 0 //'b'
		for i:=1;i<N-1;i++{
			result[i] = 1 //'m'
		}
		result[N-1] = 2 //'e'
	}
	return result
}


func CutNN(unicode_ary []rune) []string{
	result := make([]string,0)
	score_mat := make([][]float64,len(unicode_ary))
	for i,_ := range score_mat{
		score_mat[i] = make([]float64,4)
	}
	for i,_ := range unicode_ary{
		window := getWindow(unicode_ary,i)
		//fmt.Println(window)
		tmp := nn.ForwardMapForPredicate(window)
		//fmt.Println(i,tmp)
		copy(score_mat[i],tmp)
	}
	//fmt.Println(score_mat)
	best_score := make([]float64,len(unicode_ary))
	best_len := make([]int,len(unicode_ary))

	for i:=len(unicode_ary)-1;i>-1;i--{
		best_score_i := 0.0
		best_len_i := 0
		for l:=1;l<=10;l++{
			if i+l-1 >= len(unicode_ary){
				break
			}
			tags := getTags(l)
			score := 0.0
			for j:=0;j<l;j++{
				score += score_mat[i+j][tags[j]]
			}
			if i+l<len(unicode_ary){
				score += best_score[i+l]
			}
			//fmt.Println(i,l,score)
			if score > best_score_i{
				best_score_i = score
				best_len_i = l
			}
		}
		best_score[i] = best_score_i
		best_len[i] = best_len_i
	}
	
	i := 0
	for{
		if i>=len(unicode_ary){
			break
		}
		step := best_len[i]
		result = append(result,string(unicode_ary[i:i+step]))
		i += step
	}
	return result
}


func init(){
	t1 := time.Now()
	fmt.Fprintln(os.Stderr,"loading NN model...")
	nn = gonn.LoadPCN("SimpleChinese.model")
	fmt.Fprintf(os.Stderr, "NN model loaded. cost %v\n", time.Now().Sub(t1))
}


