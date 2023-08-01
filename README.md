## Zadatak 1:

Napisati API upotrebom standardne biblioteke ili u frameworku po izboru koji ima sledece mogucnosti:
    * upload slike u folder
        * ako slika postoji odbijte upload razumnom porukom
        * naziv slike je SHA256(slike) tako da cemo onemoguciti duplikate
    * listanje uploadovanih slika gde je izlaz JSON output sa nazivima slika.
    * sortiraj nazive A-Z rastuce
    * brisanje odabrane slike, brise se po nazivu slike
    * download odabrane slike, preuzima se po nazivu slike

Pitanja na koja treba dati odgovore:
    * koji su potencijalni izazovi ovakvog mikroservisa u cilju skaliranja?
    * kako bi obradio velike fajlove od po 1-2GB?

Dodatno (nije obavezno):
    * Go unit testovi
    * Dockerizovati mikroservis


## Zadatak 2:

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


## Zadatak 5:

Napravi main.go fajl i u njemu funkciju koja radi sledece:
_ sa komandne linije uz pomoc flag-ova primi ulazne parametre
_ validira ih
_ ulaz je nekakav string "[ 1, 2, 3, 4, 5 ]" koji treba da pretvoris u niz int-ova
_ treba da uradis deduplikaciju
_ sortiraj izlaz
_ prikazi izlaz na stdout (konzoli)

Pitanja na koja treba dati odgovore: \* Koja je razlika izmedju 'var' i ':=' definisanja varijabli?

Dodatno (nije obavezno):
_ uradi deduplikaciju 'in place'
_ Go unit testovi
