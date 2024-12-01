#Requires -PSEdition Core

Write-Host ("Day   Part   Duration")

Get-ChildItem -Directory -Filter "Day*" | ForEach-Object {
	cd $_.name

	$dayNumber = $_.Name -replace 'Day', ''

	make -s part1
	make -s part2

    for ($part = 1; $part -le 2; $part++) {
		Write-Host ($dayNumber.PadLeft(3)) -NoNewline
        Write-Host ($part.ToString().PadLeft(7)) -NoNewline

        $START = [System.Diagnostics.Stopwatch]::StartNew()
        $OUT = & make "run$part"
        $START.Stop()
        $END = $START.ElapsedMilliseconds

		Write-Host (($END.ToString("F3") + "ms").PadLeft(11))
    }

    & make -s clean

    cd ..
}
