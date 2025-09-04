package main

import (
	"belajar-golang-dasar/cmd"
	"belajar-golang-dasar/pkg/env"
	"flag"
)

func main() {
	env.LoadEnv()

	runSeeder := flag.Bool("seed", false, "Menjalanakan seeder")
	runMigration := flag.Bool("migrate", false, "Menjalankan migration")
	flag.Parse()

	if *runMigration {
		cmd.Migration()
	}

	if *runSeeder {
		cmd.Seeder()
	}

	if !*runMigration && !*runSeeder {
		cmd.App()
	}
}
