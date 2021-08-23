package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCanUnmarshallResponseToStruct(t *testing.T) {
	input := `{
    "preparationGroups": [
        {
            "hed": "Mango jelly",
            "steps": [
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Blend lime zest, lime juice, mango nectar, half of mangoes, and a small pinch of salt in a blender until very smooth."
                        ]
                    ]
                }
            ],
            "microSteps": []
        },
        {
            "hed": "Flan",
            "steps": [
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Place a rack in middle of oven; preheat to 350°. Cook sugar and 1 Tbsp. water in a small saucepan over medium heat, stirring until sugar is dissolved, then cook, undisturbed, swirling pan occasionally and brushing down sides with a wet pastry brush as needed, until mixture turns deep amber, 7–9 minutes. Scrape caramel into another 9x5\" loaf pan."
                        ]
                    ]
                },
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Whisk egg yolks, coconut milk, condensed milk, vanilla, and salt in a medium bowl to combine. Pour through a fine-mesh sieve into loaf pan with caramel. Cover with foil. Set inside a roasting pan and place on oven rack. Carefully pour hot water into roasting pan to come halfway up sides of loaf pan (a tea kettle makes this easy if you have one)."
                        ]
                    ]
                },
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Bake flan until just set in the center (it should still wobble a little when gently jiggled) and barely golden on top, 55–60 minutes. Let cool (still covered), then chill until cold, at least 4 hours. "
                        ],
                        [
                            "p",
                            [
                                "strong",
                                "Do ahead:"
                            ],
                            " Flan can be made 2 days ahead. Keep chilled."
                        ]
                    ]
                }
            ],
            "microSteps": []
        },
        {
            "hed": "Corn and Assembly",
            "steps": [
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Spread out corn on a rimmed baking sheet. Sprinkle sugar and salt over, then drizzle oil over and toss to combine. Roast until crisped slightly but not yet browned and a little sticky, 12–15 minutes. Let cool."
                        ]
                    ]
                },
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Meanwhile, bring sago pearls and ½ cup water to a simmer in a small saucepan and cook until translucent and firm-tender, 8–10 minutes. Drain, then combine in a small bowl with coconut milk, stirring well to separate pearls. Let cool."
                        ]
                    ]
                },
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Toast coconut in a small skillet over medium heat until edges start to brown and crisp, about 5 minutes (be careful not to overcook)."
                        ]
                    ]
                },
                {
                    "description": [
                        "div",
                        [
                            "p",
                            "Cut flan into 1\"pieces and divide some among tall glasses or wide bowls. Cut out mango jelly into desired shapes and arrange on top, then spoon in ice cream, condensed milk, shaved ice, strawberries, macapuno, palm seeds, corn, sago pearls, toasted coconut, red beans, and pinipig as desired."
                        ]
                    ]
                }
            ],
            "microSteps": []
        }
    ]
}`

	var data AutoGenerated

	decErr := json.Unmarshal([]byte(input), &data)

	if decErr != nil {
		t.Errorf(decErr.Error())
	}

	for _, desc := range data.PreparationGroups[0].Steps[0].Description {
		_, isString := desc.(string)
		if !isString {

			for _, value := range desc.([]interface{}) {
				if len(value.(string)) > 3 {
					fmt.Println(value.(string))
				}
			}
		}
	}
}

type AutoGenerated struct {
	PreparationGroups []struct {
		Hed   string `json:"hed"`
		Steps []struct {
			Description []interface{} `json:"description"`
		} `json:"steps"`
		MicroSteps []interface{} `json:"microSteps"`
	} `json:"preparationGroups"`
}