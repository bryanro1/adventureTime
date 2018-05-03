package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type choices struct {
	cmd 			string // player choice
	description 		string
	nextNode 		*storyNode
	nextChoice 		*choices
}

type storyNode struct {
	text 			string
	choices 		*choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}

	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd,":",currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd){
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}
var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func build() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text:`
	You have grown up in a traveling troupe and your entire family
	has been killed.  You have survived in the wild but you are now
	in Tarbean, an enormous city.  You hitched a ride with a farmer
	and you are now in a large market.

	You are given options at every turn, will you make it out of this
	city alive?
	`} //sets anything I dont use to nil
	alley1 := storyNode{text:"One of the many winding alleys"}
	alley2 := storyNode{text:"The alley you were attacked in"}
	farmer := storyNode{text:"I can try to help you..."}
	market := storyNode{text:"A busy market"}
	market2 := storyNode{text:"A busy market with no one willing to help you"}
	stopped := storyNode{text:"'Where do you think you're going?'"}
	answer := storyNode{text:"It's a group of young boys, all look almost as rough as you"}
	badAnswer := storyNode{text:"They grab you, throw you on the ground, and take everything"}
	hungry := storyNode{text:"You're hungry and tired, you need to find food and somewhere to rest"}
	staircase1 := storyNode{text:"A dingy looking staircase where children are coming from with bread"}
	staircase2 := storyNode{text:"The staircase between Trapis' basement and the alleys"}
	basement := storyNode{text:"The kind man Trapis is here"}
	food := storyNode{text:"Trapis offers you food for some work, a fair trade"}
	story:= storyNode{text:"'Ahhhhhhhhhhhhhhhhhhh well let me tell you about Tehlu.........."}
	docks := storyNode{text:"You've heard that a bad sort hangs around here"}
	temple := storyNode{text:"A place of worship with holy men walking in and out"}
	map1 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------   XX   -------        |
		SSS -----------   XX   -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	map11 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------   XX   -------        |
		SSS -----------   XX   -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	map2 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 |X          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	map3 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 |X          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	map4 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  |X    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	map5 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||XX||||
								------
								|    |
								|    |
								------
	`}
	map6 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|  X |
								|    |
								------
	`}
	map7 := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  |	_	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |X |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	mapdocks := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  | _	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS  X --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	mapbar := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  | _	      ||          |       |
		SSS	   =X|	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	maptemple := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |			|
		SSS	__		   ---  ----       _______
		SSS	  | _	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        ------- X      |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}
	mapnorth := storyNode{text:`
		SSSSS          _________
		SSSS		  |			|
		SSS 		  |	  X		|
		SSS	__		   ---  ----       _______
		SSS	  | _	      ||          |       |
		SSS	   = |	      ||          |       |
		S	  |	-		--  --        |       |
		SSS   |        |      |       |       |
		SSS    --------        -------        |
		SSS -----------        -------        |
		SSS            |      |       |       |
		SS              --  --        |       |
		S                 ||          |       |
		S                 | =======   |       |
					      ||   |  |   |       |
						  ||    --     -------
						  ||_____
						   ------|
							||||||||||
								------
								|    |
								|    |
								------
	`}

	north1 := storyNode{text:"Before you get anywhere a guard stops you"}
	badEnding1 := storyNode{text:"The guard grabs and throws you and then darkness..."}

	start.addChoice("South", "Into the dark alley to the south of the market", &alley1)
	start.addChoice("Talk", "He just wants to help", &farmer)

	farmer.addChoice("Ignore", "Your heart is heavy but you just can't accept it", &market)

	market.addChoice("South", "Into the dark alley to the south of the market", &alley1)
	market.addChoice("Map", "Map of Tarbean", &map1)
	map1.addChoice("Close", "Close your map", &market)

	market2.addChoice("North", "You've heard that the rich people live there... I've hear to not go there", &north1)
	market2.addChoice("South", "Into the dark alley to the south of the market", &alley2)
	market2.addChoice("West", "You can smell ocean air", &docks)
	market2.addChoice("East", "Men in robes are coming from there", &temple)
	market2.addChoice("Map", "Map of Tarbean", &map11)
	map11.addChoice("Close", "Close your map", &market2)

	north1.addChoice("Run", "Run away", &market2)
	north1.addChoice("Fight", "He's not THAT big", &badEnding1)
	north1.addChoice("Map", "Map of Tarbean", &mapnorth)
	mapnorth.addChoice("Close", "Close your map", &north1)

	alley1.addChoice("North", "Back to the market", &market)
	alley1.addChoice("South", "Deeper into the maze of a city", &stopped)
	alley1.addChoice("Map", "Map of Tarbean", &map2)
	map2.addChoice("Close", "Close your map", &alley1)

	stopped.addChoice("Inquire", "'Who's there?'", &answer)
	stopped.addChoice("Run", "Try to get away", &badAnswer)

	answer.addChoice("Fight", "Attempt to fight them", &badAnswer)

	badAnswer.addChoice("Stand", "Try to gather your dignity", &hungry)

	hungry.addChoice("North", "Back towards the alley you were beaten in", &alley2)
	hungry.addChoice("East", "Even deeper...", &staircase1)
	hungry.addChoice("Map", "Map of Tarbean", &map3)
	map3.addChoice("Close", "Close your map", &hungry)

	home := storyNode{text:"Your own place of respite with your meager belongings"}

	alley2.addChoice("North", "Back towards the market", &market2)
	alley2.addChoice("East", "Even deeper...", &staircase1)
	alley2.addChoice("Up", "You see a cozy looking spot", &home)
	alley2.addChoice("Map", "Map of Tarbean", &map4)
	map4.addChoice("Close","Close your map", &alley2)

	home.addChoice("Sleep", "These days are long and hard...", &home)
	home.addChoice("Eat", "You have a few scraps", &home)
	home.addChoice("Down", "Back to the world below", &alley2)
	home.addChoice("Map", "Map of Tarbean", &map7)
	map7.addChoice("Close","Close your map", &home)

	staircase1.addChoice("Up", "To the alley above", &alley2)
	staircase1.addChoice("Down","Through the door", &basement)
	staircase1.addChoice("Map","Map of Tarbean", &map5)
	map5.addChoice("Close", "Close your map", &staircase2)

	basement.addChoice("Talk","'What what...'", &basement)
	basement.addChoice("Food","You saw the other children with it", &food)
	basement.addChoice("Map", "Map of Tarbean", &map6)

	staircase2.addChoice("Up", "To the alley above", &alley2)
	staircase2.addChoice("Down","Through the door", &basement)
	staircase2.addChoice("Map","Map of Tarbean", &map5)
	map5.addChoice("Close", "Close your map", &staircase2)

	food.addChoice("Story","You crave a bit of kindness and familiarity", &story)
	food.addChoice("Leave", "Back to the unkind world", &staircase2)

	story.addChoice("Leave", "Back to the unkind world", &staircase2)

	bar := storyNode{text:"Children are lined up to hear the storyteller Skarpi"}

	docks.addChoice("East", "To the market", &market2)
	docks.addChoice("Bar", "You've heard of a storyteller", &bar)
	docks.addChoice("Map","Map of Tarbean", &mapdocks)
	mapdocks.addChoice("Close", "Close your map", &docks)

	skarpi := storyNode{text:"He knows every store ever told..."}

	bar.addChoice("Listen", "His voice is soothing", &skarpi)
	bar.addChoice("Leave", "Back to docks", &docks)
	bar.addChoice("Map","Map of Tarbean", &mapbar)
	mapbar.addChoice("Close", "Close your map", &bar)

	skarpi2 := storyNode{text:"Ah yes, that is certainly an old tale"}

	skarpi.addChoice("Challenge", "He can't know THAT story", &skarpi2)
	skarpi.addChoice("Leave","Back to docks", &docks)

	goodEnding := storyNode{text:`
	You're awakened from the stupor you've been in since 
	you arrived in this city. You remember your family
	and your aspirations and decide it is finally time
	to leave this awful place and travel North.

	Congratulations, you have escaped the city!
	`}

	skarpi2.addChoice("Act","'Now is the time to go to the university!'", &goodEnding)
	skarpi2.addChoice("Leave", "Back to docks", &docks)

	upnclose := storyNode{text:"'Get away you filthy child!'"}

	temple.addChoice("East", "Closer to the large holy building", &upnclose)
	temple.addChoice("West", "Back to the market", &market2)
	temple.addChoice("Map","Map of Tarbean", &maptemple)
	maptemple.addChoice("Close", "Close your map", &temple)

	upnclose.addChoice("Run", "Get away", &temple)

	start.play()
}

func main() {
	build()
}
