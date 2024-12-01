open System.IO

module Day1 =
    let private getData path =
        File.ReadLines path
        |> Seq.map(fun x -> x.Split("   "))
        |> Seq.map(fun x -> x |> Array.map int64)
        |> Seq.map(fun x -> (x.[0], x.[1]))
        |> Seq.toArray
        |> Array.unzip

    let RunPart1 path =
        path
        |> getData
        |> fun (x, y) -> (x |> Array.sort, y |> Array.sort)
        |> fun (x, y) -> Array.zip x y
        |> Array.map(fun (x, y) -> abs(x - y))
        |> Array.sum
        |> printfn "%d"

    let RunPart2 path =
        let x, y = path |> getData
        x
        |> Array.map(fun w ->
            y
            |> Array.filter(fun z -> w = z)
            |> Array.length
            |> int64
        )
        |> Array.zip x
        |> Array.map(fun (a, b) -> a * b)
        |> Array.sum
        |> printfn "%d"

[<EntryPoint>]
let main argv =
    if argv.Length = 0 then
        printfn "Specify -s=1 or -s=2 to denote part"
        1
    else
        if argv.[0] = "-s=1" then
            if argv.Length = 2 && argv.[1] = "-t" then
                // Run tests
                Day1.RunPart1 "../day1.test"
                0
            else
                // Run actual data
                Day1.RunPart1  "../day1.input"
                0

        else
            if argv.Length = 2 && argv.[1] = "-t" then
                // Run tests
                let _ = Day1.RunPart2 "../day1.test"
                0
            else
                // Run actual data
                let _ = Day1.RunPart2 "../day1.input"
                0


