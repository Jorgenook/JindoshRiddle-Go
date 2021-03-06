# JindoshRiddle

The Jindosh Riddle from Dishonered 2, solved using a **Go** bruteforce search.

__IMPORTANT!__ Those clever game devs have set it up so the puzzle is different each time. The text below is from my latest playthrough. In order to solve another puzzle, you will need to modify the contents of rules.go

The raw text of the puzzle (again, from my latest playthrough) is as follows:

_At the dinner party were Lady Winslow, Doctor Marcolla, Countess Contee, Madam Natsiou, and Baroness Finch_

_The women sat in a row. They all wore different colors and Madam Natsiou wore a jaunty purple hat. Countess Contee was at the far left, next to the guest wearing a white jacket. The lady in red sat left of someone in blue. I remember that red outfit because the woman spilled her beer all over it. The traveler from Dunwall was dressed entirely in green. When one of the dinner guests bragged about her Bird Pendant, the woman next to her said they were finer in Dunwall, where she lived._

_So Doctor Marcolla showed off a prized Ring, at which the lady from Dabokva scoffed, saying it was no match for her War Medal. Someone else carried a valuable Snuff Tin and when she saw it, the visitor from Baleton next to her almost spilled her neighbor's whiskey. Lady Winslow raised her rum in toast. The lady from Fraeport, full of absinthe, jumped up onto the table, falling onto the guest in the center seat, spilling the poor woman's wine. Then Baroness Finch captivated them all with a story about her wild youth in Karanca._

_In the morning, there were four heirlooms under the table: the Bird Pendant, Diamond, the War Medal, and the Snuff Tin._

_But who owned each?_

## Running

This is a console app. When run, it will brute force combinations, checking each against the rules derived from above, and return a final solution. Time taken is completely dependent on the speed of the running machine, and is also effected by the puzzle. 

## Go vs F#

I originally built this in F# [here](https://github.com/ChrisPritchard/JindoshRiddle). As I have learnt Go, I thought this particular problem would be better solved in that language, although ultimately they, impressively, run at about the same speed (once I fixed a lazy execution bug in the F# version). 

Both languages are great: I recommend learning both, and using each for their strengths! If anything, these two projects provide a nice comparison of the way the two languages approach things, as they are structured and implemented almost identically.