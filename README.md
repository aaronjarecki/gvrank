#GvRank

##Usage
`gvrank <inputfile>`

or 

`gvrank <inputfile> verbose`

or 

`gvrank <inputfile> -dot`

The first version of the command will return output that looks like this:
```
       Dragon: 20.68%
         Rock: 8.04%
        Grass: 7.54%
      Psychic: 6.23%
     Fighting: 5.93%
          Ice: 5.60%
          Bug: 4.86%
        Ghost: 4.85%
       Flying: 4.52%
         Fire: 4.41%
         Dark: 4.38%
       Poison: 4.28%
       Ground: 4.13%
        Fairy: 3.76%
        Water: 3.76%
        Steel: 3.57%
       Normal: 1.84%
     Electric: 1.63%
```
The higher the percentage, the more likely it is that any chain of relationships will end with or include that node.

The "verbose" version of the command will show you all the steps leading up to the final result. It is *quite* verbose, so be ready for lots of text.

The "-dot" version of the command won't output percentages to the screen. Instead, it will create a new version of your input file (with ".ranked" appended to the file name), and adds the percentages to the names of the nodes. For example:
```DOT
pokemon.gv
digraph G {
	start = true;
	dimen = 3;
	sep="+20";
	overlap = false;
	center = true;
	splines = true;
	concentrate = true;

	/* Domination Cycle*/
	edge[color = "firebrick4", len = 0.1]
	Fighting -> Rock;
	Fighting -> Dark;
	Flying -> Fighting;
	Poison -> Grass;
	Ground -> Poison;
	Rock -> Flying;
	Rock -> Fire;
	Steel -> Rock;
	Fire -> Steel;
	Fire -> Grass;
	Water -> Fire;
	Grass -> Ground;
	Grass -> Water;
	Psychic -> Fighting;
	Dark -> Psychic;

	/* Dominates */
	edge[color = "black", len = 2]
	Poison -> Fairy
	Flying -> Bug
	Flying -> Grass
	Ground -> Rock
	Ground -> Electric
	Bug -> Grass
	Steel -> Ice
	Steel -> Fairy
	Fire -> Bug
	Fire -> Ice
	Electric -> Flying
	Dark -> Ghost
	Fairy -> Fighting
	Fairy -> Dragon
	Fairy -> Dark

	/* Strong Against */
	edge[style=dashed, len=2]
	Fighting -> Normal;
	Fighting -> Steel;
	Fighting -> Ice;
	Ground -> Steel;
	Ground -> Fire;
	Rock -> Bug;
	Rock -> Ice;
	Bug -> Psychic;
	Bug -> Dark;
	Ghost -> Ghost;
	Ghost -> Psychic;
	Water -> Ground;
	Water -> Rock;
	Grass -> Rock;
	Electric -> Water;
	Psychic -> Poison;
	Ice -> Flying;
	Ice -> Rock;
	Ice -> Grass;
	Ice -> Dragon;
	Dragon -> Dragon;
}
```
becomes:
```DOT
pokemon.gv.ranked
digraph G {
	Fighting [label = "Fighting: 5.93%"]
	Poison [label = "Poison: 4.28%"]
	Water [label = "Water: 3.76%"]
	Electric [label = "Electric: 1.63%"]
	Ghost [label = "Ghost: 4.85%"]
	Flying [label = "Flying: 4.52%"]
	Grass [label = "Grass: 7.54%"]
	Fire [label = "Fire: 4.41%"]
	Steel [label = "Steel: 3.57%"]
	Fairy [label = "Fairy: 3.76%"]
	Ice [label = "Ice: 5.60%"]
	Normal [label = "Normal: 1.84%"]
	Rock [label = "Rock: 8.04%"]
	Dark [label = "Dark: 4.38%"]
	Ground [label = "Ground: 4.13%"]
	Bug [label = "Bug: 4.86%"]
	Psychic [label = "Psychic: 6.23%"]
	Dragon [label = "Dragon: 20.68%"]
	start = true;
	dimen = 3;
	sep="+20";
	overlap = false;
	center = true;
	splines = true;
	concentrate = true;

	/* Domination Cycle*/
	edge[color = "firebrick4", len = 0.1]
	Fighting -> Rock;
	Fighting -> Dark;
	Flying -> Fighting;
	Poison -> Grass;
	Ground -> Poison;
	Rock -> Flying;
	Rock -> Fire;
	Steel -> Rock;
	Fire -> Steel;
	Fire -> Grass;
	Water -> Fire;
	Grass -> Ground;
	Grass -> Water;
	Psychic -> Fighting;
	Dark -> Psychic;

	/* Dominates */
	edge[color = "black", len = 2]
	Poison -> Fairy
	Flying -> Bug
	Flying -> Grass
	Ground -> Rock
	Ground -> Electric
	Bug -> Grass
	Steel -> Ice
	Steel -> Fairy
	Fire -> Bug
	Fire -> Ice
	Electric -> Flying
	Dark -> Ghost
	Fairy -> Fighting
	Fairy -> Dragon
	Fairy -> Dark

	/* Strong Against */
	edge[style=dashed, len=2]
	Fighting -> Normal;
	Fighting -> Steel;
	Fighting -> Ice;
	Ground -> Steel;
	Ground -> Fire;
	Rock -> Bug;
	Rock -> Ice;
	Bug -> Psychic;
	Bug -> Dark;
	Ghost -> Ghost;
	Ghost -> Psychic;
	Water -> Ground;
	Water -> Rock;
	Grass -> Rock;
	Electric -> Water;
	Psychic -> Poison;
	Ice -> Flying;
	Ice -> Rock;
	Ice -> Grass;
	Ice -> Dragon;
	Dragon -> Dragon;
}
```
