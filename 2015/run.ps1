$path = "E:\DISK\Github\advent-of-code\2015\src\main.rs"

function modify {
    param (
        [string]$path,
        [int]$i,
        [int]$j
    )
    $x = $i.ToString().PadLeft(2, '0')
    $y = $j.ToString().PadLeft(2, '0')

    echo $x $y
    
    $content = Get-Content -Path $path
    $content = $content -replace "day$x", "day$y"
    Set-Content -Path $path -Value $content
}

$first = 1
$last = 18

Write-Output "Build 1" > tmp/build.log
cargo build 2>> tmp/build.log
Write-Output "" >> tmp/build.log

($first..$last) | ForEach-Object {
    Write-Output $_
    modify -path $path -i $_ -j ($_ + 1)

    Write-Output "Build $($_ + 1)" >> tmp/build.log
    cargo build 2>> tmp/build.log
    Write-Output "" >> tmp/build.log
}

modify -path $path -i ($last + 1) -j $first
