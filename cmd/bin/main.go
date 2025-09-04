package main

import (
	commonutils "belajar-golang-dasar/app/common/utils"
	"belajar-golang-dasar/cmd"
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
