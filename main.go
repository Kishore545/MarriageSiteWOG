package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Family struct {
	Name    string
	Values  string
	Finance string
}

type Girl struct {
	Name     string
	Hobbies  string
	Career   string
	Marriage string
}

func visitHouse() {
	fmt.Println("1️⃣ Arrival at the Girl's House")
	fmt.Println("✅ Respectful Entry - Greeting elders with Namaste")
	fmt.Println("✅ Observing Surroundings - Checking hospitality and atmosphere")
	fmt.Println("✅ Casual Conversation - Talking about family, weather, interests")
}

func meetFamily() Family {
	fmt.Println("2️⃣ Meeting the Girl & Family")
	return Family{Name: "Sharma Family", Values: "Traditional", Finance: "Stable"}
}

func meetGirl() Girl {
	fmt.Println("✅ Speaking with the Girl")
	return Girl{Name: "Priya", Hobbies: "Reading, Music", Career: "Software Engineer", Marriage: "Open-minded"}
}

func discussWithFamily(f Family, g Girl) bool {
	fmt.Println("✅ Family Discussion")
	fmt.Printf("Family Values: %s, Financial Status: %s\n", f.Values, f.Finance)
	fmt.Printf("Girl's Career: %s, Marriage Views: %s\n", g.Career, g.Marriage)

	if f.Values == "Traditional" && g.Marriage == "Open-minded" {
		fmt.Println("⚠ Potential mismatch in values!")
		return false
	}
	return true
}

func finalDecision(match bool) {
	fmt.Println("3️⃣ Decision-Making & Sendoff")
	if match {
		fmt.Println("✔ Outcome: Selected! Both families agree to discuss further.")
	} else {
		fmt.Println("❌ Outcome: Rejected! Families part on a respectful note.")
	}
	fmt.Println("✅ Thanking the hosts and leaving with dignity.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	visitHouse()
	family := meetFamily()
	girl := meetGirl()
	match := discussWithFamily(family, girl)
	finalDecision(match)
}
