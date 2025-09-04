package main

import (
	"belajar-golang-dasar/cmd"
	commonutils "belajar-golang-dasar/common/utils"
	"flag"
)

func main() {
	commonutils.LoadEnv()

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
