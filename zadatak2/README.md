Napisati API upotrebom standardne biblioteke ili u frameworku po izboru koji ima sledece mogucnosti:
    * upload ruta
        * prihvata JSON ulaz koji je ova struktura
        type Input struct {
            Operation string
            Data      []int64
        }

    * Moguce Operacije su:
        * deduplicate - eliminisi iz niza duplikate, primer:
            [ 1, 1, 2, 3, 4, 5 ] => [ 1, 2, 3, 4, 5 ]
       
        * getPairs - nadji parove u nizu i grupisi ih, primer:
            [ 1, 1, 1, 2, 3, 4, 5, 5, 5, 5, 6, 7, 8, 9 ] ->
            map[int64]int{
                1: 3,
                5: 4
            }
   
    * Output je definisan kao:
        type Output struct {
            ID        string        // ID requesta koje ces sam definisati da bude unikatno
            Operation string
            Data      []int64
        }

    * za nevalidnu operaciju, prikazi gresku

Dodatno (nije obavezno):
    * Uradi deduplicate 'in place'
    * Go unit testovi
    * Dockerizovati mikroservis