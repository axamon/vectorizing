package vectorizing

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/axamon/gobrain"
	"github.com/axamon/persist"

	"github.com/cdipaolo/goml/base"
)

const alfabeto string = "abcdefghijklmnopqrstuvwxyz"
const alfabetoNoVocali string = "bcdfghjklmnpqrstvwxyz"
const ascii string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@:-$."

var isLetter *regexp.Regexp

// CreateNetwork creates a new network
func CreateNetwork() (ff *gobrain.FeedForward, err error) {
	ff = &gobrain.FeedForward{}
	return ff, err
}

// TrainNetwork trains the network
func TrainNetwork(ff gobrain.FeedForward, inputs, hiddens, outputs, iterations int, patterns [][][]float64) (ffout gobrain.FeedForward, err error) {

	/*
	   // Esempio di patterns
	   	patterns := [][][]float64{
	   		{{0, 0}, {0}},
	   		{{0, 1}, {1}},
	   		{{1, 0}, {1}},
	   		{{1, 1}, {0}},
	   	}
	*/

	// Istanzia la rete neurale
	ffout = ff

	// Struttura gli accessi, gli strati nascosti e le uscite
	ffout.Init(inputs, hiddens, outputs)

	// Istruisce la rete
	ffout.Train(patterns, iterations, 0.6, 0.4, true)

	/*
		// Crea l'input da passare in ingresso alla rete
		input := []float64{1, 1}

		// Inva l'input alla rete e salva il risultato
		result = ff.Update(input)

		// Scrive il il risultato a video
		// fmt.Println(result) // debug
	*/
	return ffout, err

}

// SaveNetwork saves network on file
func SaveNetwork(file string, ff gobrain.FeedForward) (err error) {

	/*
	   // Esempio di patterns
	   	patterns := [][][]float64{
	   		{{0, 0}, {0}},
	   		{{0, 1}, {1}},
	   		{{1, 0}, {1}},
	   		{{1, 1}, {0}},
	   	}
	*/

	// Salva la rete su file
	err = persist.Save(file, ff)
	if err != nil {
		log.Printf("Saving Network to file %s not possible %s", file, err.Error())
	}
	/*
		// Crea l'input da passare in ingresso alla rete
		input := []float64{1, 1}

		// Inva l'input alla rete e salva il risultato
		result = ff.Update(input)

		// Scrive il il risultato a video
		// fmt.Println(result) // debug
	*/
	return err

}

// LoadNetwork loads network from file
func LoadNetwork(file string) (ff *gobrain.FeedForward, err error) {
	ff = &gobrain.FeedForward{}

	err = persist.Load(file, &ff)
	if err != nil {
		log.Printf("Loading network from file %s not possible %s", file, err.Error())
	}
	return ff, err
}

func Vectorize(segment string) (vector []float64) {

	// Elimina eventuali spazi tra le parole
	solosegment := strings.Replace(segment, " ", "", -1)

	//solosegment = strings.ToLower(solosegment)

	//var vector [26]float64
	l := len(solosegment)
	for index, letter := range alfabeto {
		tot := strings.Count(solosegment, string(letter))
		vector = append(vector, float64(tot))
		vector[index] = float64(tot) / float64(l)
	}

	//fmt.Println(vector) // Debug

	return vector

}

const alfabetonovocali string = "bcdfghjklmnpqrstvwxyz"

func VectorizeNoVocali(segment string) (vector [21]float64) {

	// Elimina eventuali spazi tra le parole
	solosegment := strings.Replace(segment, " ", "", -1)

	solosegment = strings.ToLower(solosegment)

	// var vector [26]float64
	l := len(solosegment)
	for index, letter := range alfabetoNoVocali {
		tot := strings.Count(solosegment, string(letter))
		vector[index] = float64(tot) / float64(l)
	}

	//fmt.Println(vector) // Debug

	return vector

}

func AddVectorsAndNormalize(a, b [26]float64) (sumNormalized [26]float64, err error) {

	if len(a) != len(b) {
		err = fmt.Errorf("due vettori con len differente")
		return
	}

	for i := 0; i < len(a); i++ {
		sumNormalized[i] = a[i] + b[i]

	}

	var tot float64
	for _, v := range sumNormalized {
		tot += v
	}

	for i, v := range sumNormalized {
		sumNormalized[i] = v / tot
	}

	return sumNormalized, nil
}

func RecuperaTraccia(segment string) (traccia string) {

	// Elimina eventuali spazi tra le parole
	solosegment := strings.Replace(segment, " ", "", -1)

	solosegment = strings.ToLower(solosegment)

	// Trasforma la stringa in una lista di caratteri
	chars := []rune(solosegment)

	var vuota []string
	for _, char := range chars {
		switch isLetter.MatchString(string(char)) {
		case false:
			continue
		default:
			vuota = append(vuota, string(char))

		}
	}

	var lunghezza int
	lunghezza = len(vuota)
	// fmt.Println(lunghezza) // debug

	// Ordino la serie di stringhe alfabeticamente
	sort.Strings(vuota)

	traccia = strings.Join(vuota, "")

	freq := strings.Count(traccia, "n")
	// fmt.Println(freq) // debug

	v := float64(float64(freq) / float64(lunghezza))
	fmt.Printf("valore: %0.2f\n", v)

	return traccia
}

var persistenceFile = "/tmp/.goml/store1"

func init() {
	// create the /tmp/.goml/ dir for persistance testing
	// if it doesn't already exist!
	err := os.MkdirAll("/tmp/.goml", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

}

func impara(tokens []string, category Diagnosi) {

	stream := make(chan base.TextDatapoint, 3000)
	errors := make(chan error)

	model := NewNaiveBayes(stream, 7, base.OnlyWords)

	model.RestoreFromFile(persistenceFile)

	go model.OnlineLearn(errors)

	stream <- base.TextDatapoint{
		X: strings.Join(tokens, " "),
		Y: uint8(category),
	}

	close(stream)

	for {
		err, more := <-errors
		if more {
			fmt.Printf("Error passed: %v", err)
		} else {
			// training is done!
			break
		}
	}

	model.PersistToFile(persistenceFile)

}

func Predici(frase string) {
	stream := make(chan base.TextDatapoint, 40)

	model := NewNaiveBayes(stream, 7, base.OnlyWordsAndNumbers)
	model.RestoreFromFile(persistenceFile)

	predizione := model.Predict(frase)

	fmt.Printf("", predizione)
}
