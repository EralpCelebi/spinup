/*
		DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
				Version 2, December 2004

Copyright 2022 Eralp Çelebi

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

		DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

0. You just DO WHAT THE FUCK YOU WANT TO.

Eralp Çelebi

*/

package main

import "fmt"

func Try(Result error) {
	if Result != nil {
		Message := fmt.Sprintf("Warning: \n Error: %s\n", Result.Error());
		fmt.Println(Message);
	}
}

func Check(Result error) {
	if Result != nil {
		Message := fmt.Sprintf("Something went wrong!..\n Error: %s", Result.Error());
		panic(Message);
	}
}
