package main

import "fmt"

// on import fmt
func main() {
	//on declare nos variables
	var (
		b   bool
		s   string
		i   int
		u   uint   //c'est unsigned, pas de signe. le nombre peux pas etre negatif
		i8  int8   // de -128 à 127
		u8  uint8  //de 0 à 255
		i16 int16  // -32768 à 32767
		u16 uint16 // 0 à 65535
		i32 int32  // -2147483648 à 2147483647
		u32 uint32 //0 à 4294967295
		i64 int64  // -9223372036854775808 à 9223372036854775807
		u64 uint64 // 0 à 18446744073709551615
		//int est par defaut int64 et uint est par defaut int64 (si on est sur 32 bits alors sa va etre int32 et uint32)
		f32 float32 //ya pas de type float tout court
		//f64 float64 quand ya trop de nombre apres la virgle (rare)
	)
	//on assigne nos variables
	b = true //ou false
	s = "Bonjour"
	i = -15
	u = 10
	i8 = 125
	u8 = 250
	i16 = 32765
	u16 = 65530
	i32 = 100000000  //cent million
	u32 = 4000000000 //4 milliard (4 mille millions)
	i64 = 1000000000000000000
	u64 = 10000000000000000000
	f32 = 3.1415

	//on traite nos variables(cette fois avec println sans formatage)
	fmt.Println(b, s, i, u, i8, u8, i16, u16, i32, u32, i64, u64, f32)
}
