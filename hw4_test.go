package hillcipher

import (
	"testing"
	"reflect"

	"github.com/UofSC-Fall-2022-Math-587-001/homework4/linearalg"
)

func Test1(t *testing.T) {
	K1 := linearalg.Matrix2x2{1,3,2,2}
	K2 := linearalg.Vector2{5,4}
	H := HillCipherKey{K1,K2}
	message := linearalg.Vector2{2,1}
	got := H.encrypt(message,7)
	want := linearalg.Vector2{3,3}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test2(t *testing.T) {
	K1 := linearalg.Matrix2x2{1,3,2,2}
	K2 := linearalg.Vector2{5,4}
	H := HillCipherKey{K1,K2}
	ciphertext := linearalg.Vector2{3,3}
	got := H.decrypt(ciphertext,7)
	want := linearalg.Vector2{2,1}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test3(t *testing.T) {
	K1 := linearalg.Matrix2x2{2,1,3,3}
	K2 := linearalg.Vector2{1,2}
	H := HillCipherKey{K1,K2}
	message := linearalg.Vector2{1,1}
	got := H.encrypt(message,5)
	want := linearalg.Vector2{4,3}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test4(t *testing.T) {
	K1 := linearalg.Matrix2x2{2,1,3,3}
	K2 := linearalg.Vector2{1,2}
	H := HillCipherKey{K1,K2}
	ciphertext := linearalg.Vector2{4,3}
	got := H.decrypt(ciphertext,5)
	want := linearalg.Vector2{1,1}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("%d vs %d", got, want)
	}
}
