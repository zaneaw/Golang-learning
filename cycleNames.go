package main

func cycleNames(name []string, function func(string)) {
	for _, value := range name {
		function(value)
	}
}