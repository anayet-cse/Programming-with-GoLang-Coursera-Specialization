package main 
  
import ( 
    "fmt"
    "sort"
    "sync"
)

func ArraySort(wg *sync.WaitGroup,array []int,starting_index int,ending_index int,thread string){
	var sub []int
	for i:=starting_index;i<=ending_index;i++{
		sub=append(sub,array[i])
	}
	sort.Ints(sub)
	j:=0
	for i:=starting_index;i<=ending_index;i++{
		array[i]=sub[j]
		j++
	}
	fmt.Println("Array Sorted by thread "+thread+": ",sub)
	//fmt.Println(sub)
	wg.Done()
} 

func main(){
	var n int
	fmt.Println("Enter size of array")
	fmt.Scanln(&n)
	fmt.Println("Enter the integers(Press Enter after each integer)")
	var array []int
	for i:=0;i<n;i++{
		var x int
		fmt.Scanln(&x)
		array=append(array,x)
	}
	//Computing size for 4 approx size subarray
	n1:=n/2
	n2:=n-n1
	n3:=n1/2
	n4:=n1-n3
	n5:=n2/2
	//n6:=n2-n5
	var wg sync.WaitGroup
	wg.Add(4)
	go ArraySort(&wg,array,0,n3-1,"1")
	go ArraySort(&wg,array,n3,n3+n4-1,"2")
	go ArraySort(&wg,array,n3+n4,n3+n4+n5-1,"3")
	go ArraySort(&wg,array,n3+n4+n5,n-1,"4")
	wg.Wait()
	fmt.Println("Final Sorting of whole array:")
	sort.Ints(array)
	fmt.Println(array) 
}