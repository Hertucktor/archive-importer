package mongodb

type MultipleCards struct {
	Cards []Card `json:"cards"`
}
type Card struct {
	Name          string         `json:"name,omitempty" bson:"name,omitempty"`
	ManaCost      string         `json:"manaCost,omitempty" bson:"manaCost,omitempty"`
	Cmc           float64        `json:"cmc,omitempty" bson:"cmc,omitempty"`
	Colors        []string       `json:"colors,omitempty" bson:"colors,omitempty"`
	ColorIdentity []string       `json:"colorIdentity,omitempty" bson:"colorIdentity,omitempty"`
	Type          string         `json:"type,omitempty" bson:"type,omitempty"`
	Supertypes    []string       `json:"superTypes,omitempty" bson:"superTypes,omitempty"`
	Types         []string       `json:"types,omitempty" bson:"types,omitempty"`
	Subtypes      []string       `json:"subTypes,omitempty" bson:"subTypes,omitempty"`
	Rarity        string         `json:"rarity,omitempty" bson:"rarity,omitempty"`
	Set           string         `json:"set,omitempty" bson:"set,omitempty"`
	SetName       string         `json:"setName,omitempty" bson:"setName,omitempty"`
	Text          string         `json:"text,omitempty" bson:"text,omitempty"`
	Flavor        string         `json:"flavor,omitempty" bson:"flavor,omitempty"`
	Artist        string         `json:"artist,omitempty" bson:"artist,omitempty"`
	Number        string         `json:"number,omitempty" bson:"number,omitempty"`
	Power         string         `json:"power,omitempty" bson:"power,omitempty"`
	Toughness     string         `json:"toughness,omitempty" bson:"toughness,omitempty"`
	Layout        string         `json:"layout,omitempty" bson:"layout,omitempty"`
	Multiverseid  string         `json:"multiverseID,omitempty" bson:"multiverseID,omitempty"`
	ImageURL      string         `json:"imageURL,omitempty" bson:"imageURL,omitempty"`
	Rulings       []Rulings      `json:"rulings,omitempty" bson:"rulings,omitempty"`
	ForeignNames  []ForeignNames `json:"foreignNames,omitempty" bson:"foreignNames,omitempty"`
	Printings     []string       `json:"printings,omitempty" bson:"printings,omitempty"`
	OriginalText  string         `json:"originalText,omitempty" bson:"originalText,omitempty"`
	OriginalType  string         `json:"originalType,omitempty" bson:"originalType,omitempty"`
	Legalities    []Legalities   `json:"legalities,omitempty" bson:"legalities,omitempty"`
	ID            string         `json:"id,omitempty" bson:"id,omitempty"`
	Variations    []string       `json:"variations,omitempty" bson:"variations,omitempty"`
	Quantity      int            `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Created       string         `json:"created,omitempty" bson:"created,omitempty"`
	Modified      string         `json:"modified,omitempty" bson:"modified,omitempty"`
}
type Rulings struct {
	Date string `json:"date,omitempty" bson:"date,omitempty"`
	Text string `json:"text,omitempty" bson:"text,omitempty"`
}
type ForeignNames struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Text         string `json:"text,omitempty" bson:"text,omitempty"`
	Flavor       string `json:"flavor,omitempty" bson:"flavor,omitempty"`
	ImageURL     string `json:"imageURL,omitempty" bson:"imageURL,omitempty"`
	Language     string `json:"language,omitempty" bson:"language,omitempty"`
	Multiverseid int    `json:"multiverseID,omitempty" bson:"multiverseID,omitempty"`
}
type Legalities struct {
	Format   string `json:"format,omitempty" bson:"format,omitempty"`
	Legality string `json:"legality,omitempty" bson:"legality,omitempty"`
}
