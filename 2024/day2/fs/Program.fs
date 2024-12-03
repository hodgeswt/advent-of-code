open System.IO

module Day2 =
    let private getData path =
        File.ReadLines path
        |> Seq.map(fun x -> x.Split(" "))
        |> Seq.map(fun x -> x |> Array.map int)

    let rec private checkNext (start: bool, set: bool, ascending: bool, last: int, level: int[]) =
        if level.Length = 0 then
            true
        else
            if start then
                checkNext(false, false, ascending, level.[0], level[1..])
            else
                let x = level.[0]
                let diff = x - last
                if not set then
                    if x > last then
                        if diff > 0 && diff < 4 then
                            checkNext(false, true, true, x, level[1..])
                        else
                            false
                    else
                        if diff < 0 && diff > -4 then
                            checkNext(false, true, false, x, level[1..])
                        else
                            false
                else
                    if ascending then
                        if diff > 0 && diff < 4 then
                            checkNext(false, set, ascending, x, level[1..])
                        else
                            false
                    else
                        if diff < 0 && diff > -4 then
                            checkNext(false, set, ascending, x, level[1..])
                        else
                            false

    let private isValid level =
        checkNext(true, false, false, -1, level)

    let RunPart1 path =
        path
        |> getData
        |> Seq.map isValid
        |> Seq.map(fun x -> if x then 1 else 0)
        |> Seq.sum
        |> printfn "%A"


[<EntryPoint>]
let main argv =
    if argv.Length = 0 then
        printfn "Specify -s=1 or -s=2 to denote part"
        1
    else
        if argv.[0] = "-s=1" then
            if argv.Length = 2 && argv.[1] = "-t" then
                // Run tests
                Day2.RunPart1 "../day2.test"
                0
            else
                // Run actual data
                Day2.RunPart1  "../day2.input"
                0

        else
            1

