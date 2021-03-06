package main

type position int

const (
	farLeft  position = 0
	left     position = 1
	centre   position = 2
	right    position = 3
	farRight position = 4
)

type women int

const (
	ladyWinslow    women = 0
	doctorMarcolla women = 1
	countessContee women = 2
	madamNatsiou   women = 3
	baronessFinch  women = 4
)

var womanNames = map[women]string{
	ladyWinslow:    "Lady Winslow",
	doctorMarcolla: "Doctor Marcolla",
	countessContee: "Countess Contee",
	madamNatsiou:   "Madam Natsiou",
	baronessFinch:  "Baroness Finch",
}

type colours int

const (
	purple colours = 0
	white  colours = 1
	red    colours = 2
	blue   colours = 3
	green  colours = 4
)

type heirloom int

const (
	ring        heirloom = 0
	birdPendant heirloom = 1
	diamond     heirloom = 2
	warMedal    heirloom = 3
	snuffTin    heirloom = 4
)

var heirloomNames = map[heirloom]string{
	ring:        "Prized Ring",
	birdPendant: "Bird Pendant",
	diamond:     "Diamond",
	warMedal:    "War Medal",
	snuffTin:    "Snuff Tin",
}

type drink int

const (
	beer     drink = 0
	whiskey  drink = 1
	rum      drink = 2
	absinthe drink = 3
	wine     drink = 4
)

type hometown int

const (
	dunwall  hometown = 0
	dabokva  hometown = 1
	baleton  hometown = 2
	fraeport hometown = 3
	karnaca  hometown = 4
)

type description struct {
	position position
	woman    women
	wearing  colours
	from     hometown
	drinking drink
	owns     heirloom
}

type rule struct {
	subject, target interface{}
}

type isTrue rule

type notTrue rule

type nextTo rule

type leftOf rule

type rightOf rule
