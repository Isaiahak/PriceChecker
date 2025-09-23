package main

import(
	
)

type BaseItem String

type Rarity String

const (
	Rare Rarity = "Normal" | "Magic" | "Rare"
	Unqiue Rarity = "Unique"
	Currency Rarity = "Currency"
	Gem Rarity = "Gem"
)

// check if i have missed any base types
const (
	Amulet BaseItem = "Amulets" 
	Ring BaseItem = "Rings"
  	Boots BaseItem = "Boots"
 	Gloves BaseItem = "Gloves"
 	BodyArmour BaseItem = "Body Armours"
 	Bow  BaseItem = "Bows"
 	OneHandedMace BaseItem = "One Handed Maces"
 	TwoHandedMace BaseItem = "Two Handed Maces"
 	OneHandedSpear BaseItem = "One Handed Spear"
 	TwoHandedSpear BaseItem = "Two Handed Spear"
 	TwoHandedStaff BaseItem = "Two Handed Staves"
 	Staff BaseItem = "Staves"
 	Sceptre BaseItem = "Sceptres"
 	CrossBow BaseItem = "CrossBows"
 	Currency BaseItem = "Stackable Currency"
 	SupportGem BaseItem = "Support Gem"
 	SkillGem BaseItem = "Skill Gem"
 	Waystone BaseItem = "Waystones"
)




func ParseItem(item string) {
	// that is the string containing all of the item data

	itemInfo := strings.SplitN(item,"\n",2)

	// should contain the item type value [1] value [0] key
	itemType := strings.Split(itemInfo[0],":")[1] 

	// should contain the item type value [1] value [0] key
	itemRarity := strings.Split(itemInfo[1],":")[1] 
 
	switch Rarity(itemRarity) {
		// parse all items with normal, magic, and rare rarity
		case Rare:
			switch BaseItem(itemType) {
				// everthing within here could be considered a base item
				case Amulet:
					ParseRare(item)
				case Ring:
					ParseRare(item)
				case Gloves:
					ParseRare(item)
				case Boots:
					ParseRare(item)
				case BodyArmour:
					ParseRare(item)
				case Bow:
					ParseRare(item)
				case OneHandedMace:
					ParseRare(item)
				case TwoHandedMace:
					ParseRare(item)
				case TwoHandedStaff:
					ParseRare(item)
				case Staff:
					ParseRare(item)
				case TwoHandedSpear:
					ParseRare(item)
				case OneHandedSpear:
					ParseRare(item)
				case Sceptre:
					ParseRare(item)
				case CrossBow:
					ParseRare(item)
				// everthing within here could be considered a base item
				case Waystone:
				default:
			}
		// parse all unique items
		case Unique:
			switch BaseItem(itemType){
				case Amulet:
					ParseUnique(item)
				case Ring:
					ParseUnique(item)
				case Gloves:
					ParseUnique(item)
				case Boots:
					ParseUnique(item)
				case BodyArmour:
					ParseUnique(item)
				case Bow:
					ParseUnique(item)
				case OneHandedMace:
					ParseUnique(item)
				case TwoHandedMace:
					ParseUnique(item)
				case TwoHandedStaff:
					ParseUnique(item)
				case Staff:
					ParseUnique(item)
				case TwoHandedSpear:
					ParseUnique(item)
				case OneHandedSpear:
					ParseUnique(item)
				case Sceptre:
					ParseUnique(item)
				case CrossBow:
					ParseUnique(item)
				default:
			}
		// parse all gems
		case Gem:
			switch BaseItem(itemType){
				case SupportGem:
					ParseSupportGem(item)
				case SkillGem:
					ParseSkillGem(item)
				default:
			}
		// parse all currency
		case Currency:
			switch BaseItem(itemType){
			case Currency:
				ParseCurrency(item)
			default:

			}
		// if there isn't an item type or is invalid 	
		default: 

	}
}



go func ParseRare() {

}

// skill gems will be searched based on:
// name after the Rarity new line
// Level: #
// Quality: #%
go func ParseSkillGem() {

}

// support gems will be searched based on:
// name after the Rarity new line
go func ParseSupportGem() {

}

// uniques items wil lbe searched based on:
// name after Rarity new Line
// sockets count S until new line
// implicit stat
go func ParseUnique() {

}

go func ParseCurrency() {

}


