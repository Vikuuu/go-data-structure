package tests

import ds "github.com/Vikuuu/go-datastructure"

var tree = &ds.BNode{
	Val: 20,
	Right: &ds.BNode{
		Val: 50,
		Right: &ds.BNode{
			Val:   100,
			Right: nil,
			Left:  nil,
		},
		Left: &ds.BNode{
			Val: 30,
			Right: &ds.BNode{
				Val:   45,
				Right: nil,
				Left:  nil,
			},
			Left: &ds.BNode{
				Val:   29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &ds.BNode{
		Val: 10,
		Right: &ds.BNode{
			Val:   15,
			Right: nil,
			Left:  nil,
		},
		Left: &ds.BNode{
			Val: 5,
			Right: &ds.BNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

var tree2 = &ds.BNode{
	Val: 20,
	Right: &ds.BNode{
		Val: 50,
		Right: &ds.BNode{
			Val:   100,
			Right: nil,
			Left:  nil,
		},
		Left: &ds.BNode{
			Val: 30,
			Right: &ds.BNode{
				Val:   45,
				Right: nil,
				Left:  nil,
			},
			Left: &ds.BNode{
				Val:   29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &ds.BNode{
		Val: 10,
		Right: &ds.BNode{
			Val:   15,
			Right: nil,
			Left:  nil,
		},
		Left: &ds.BNode{
			Val: 75,
			Right: &ds.BNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}
