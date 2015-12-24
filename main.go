package main

/*

beberapa kata dari pembedaharaan kata di Bahasa Indonesia merupakan
kata serapan dari bahasa Minang, namun kata-kata tersebut mengalami perubahan.
perubahan tersebut antara lain, perubahan pengucapan/pelafalan, bentuk kata.

- iang umum nya dipakai pada akhiran kata bahasa Minang seperti "ladiang", "kariang"
 	pada perubahannya di bahasa Indonesia, kata-kata tersebut berubah bentuk menjadi
	"lading" dan "kering", berikut beberapa daftar kata serupa:

	Bahasa Indonesia  	Bahasa Minang

	kucing							kuciang
	kambing 						kambiang
	keling 							kaliang
	piring							piriang
	pening							paniang
	lading 							ladiang
	kering							kariang

program ini melakukan generasi bahasa Minang terhadap
kata-kata serapan minang yang memiliki pola yang sama dengan kata-kata di atas.

PS: kesempurnaan hanya milik Tuhan!
*/
import (
	"fmt"
	"strings"
)

func RubahKeBahasaMinang(kata string)string{
	return strings.Replace(kata, "ng", "ang" , 100) // set parameter to antisipate longer input
}

func main(){
	inp := "kaling kambing kancing lading kucing kering"
	// hasil output: kaliang kambiang kanciang ladiang kuciang keriang
	fmt.Println(RubahKeBahasaMinang(inp))
}
