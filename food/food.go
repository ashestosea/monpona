package food

//go:generate go-enum --marshal --names --values --nocase --noprefix
/*ENUM(
RainbowGrass,
Ooznip,
Berry,
Groundnut,
Carnivorous
)
*/
type Type byte
