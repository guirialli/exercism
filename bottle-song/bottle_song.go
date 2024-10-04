package bottlesong

var bottler []string = []string{
	"One green bottle hanging on the wall,",
	"One green bottle hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be no green bottles hanging on the wall.",

	"Two green bottles hanging on the wall,",
	"Two green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be one green bottle hanging on the wall.",

	"Three green bottles hanging on the wall,",
	"Three green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be two green bottles hanging on the wall.",

	"Four green bottles hanging on the wall,",
	"Four green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be three green bottles hanging on the wall.",

	"Five green bottles hanging on the wall,",
	"Five green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be four green bottles hanging on the wall.",

	"Six green bottles hanging on the wall,",
	"Six green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be five green bottles hanging on the wall.",

	"Seven green bottles hanging on the wall,",
	"Seven green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be six green bottles hanging on the wall.",

	"Eight green bottles hanging on the wall,",
	"Eight green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be seven green bottles hanging on the wall.",

	"Nine green bottles hanging on the wall,",
	"Nine green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be eight green bottles hanging on the wall.",

	"Ten green bottles hanging on the wall,",
	"Ten green bottles hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be nine green bottles hanging on the wall.",
}

func Recite(startBottles, takeDown int) []string {
	verses := make([]string, 0)
	for takeDown > 0 {
		for i := 0; i < 4; i++ {
			verses = append(verses, bottler[(startBottles-1)*4+i])
		}
		startBottles--
		takeDown--
	}
	return verses

}
