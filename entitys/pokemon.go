package entitys

type Wild_Pokemon struct {
	ID                                              string
	Name                                            string
	HP, Attack, Defense, SpAttack, SpDefense, Speed string
	Available_moves                                 []string
	Moves                                           []Move
}
