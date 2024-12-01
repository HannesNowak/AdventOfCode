#Requires -PSEdition Core

param([int]$year, [int]$day)

# Set the URI
$uri = "https://adventofcode.com/$year/day/$day"

# Load environment variables
Get-Content .env | foreach {
	$name, $value = $_.split('=')
	if ([string]::IsNullOrWhiteSpace($name) || $name.Contains('#')) {
		continue
	}
	Set-Content env:\$name $value
}

# Create the directory and change to it
$dir = ".\Day$($day.ToString('D2'))"

if (-not (Test-Path $dir)) {
	New-Item $dir -ItemType Directory | Out-Null
}

Set-Location $dir

# Create go files
if (-not (Test-Path ".\part1.go")) {
	cp "..\part.template" ".\part1.go"
	cp "..\part.template" ".\part2.go"
}

# Create Makefile
if (-not (Test-Path ".\Makefile")) {
	cp "..\Makefile.template" ".\Makefile"
}

# Fetch the assignment
if (-not (Test-Path ".\assignment")) {
	curl "$uri" -s | htmlq -p '.day-desc' > .\temp.html
	lynx -dump temp.html -width 80 > .\assignment.txt
	rm -Force temp.html
}

# Fetch the input
if (-not (Test-Path ".\input")) {
	$response = Invoke-WebRequest -Uri "$uri/input" -Headers @{ "Cookie" = "session=$env:SESSION" }
	echo $response.Content > .\input
}
